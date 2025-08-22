package main

import (
	"encoding/json"
	"io"
	"net/http"
	"net/url"
	"slices"
	"strconv"
	"strings"
)

var (
	local  bool
	match  bool
	search bool
)

func FetchLyrics(track Song) []Lyric {
	matchUrl := c.Internal.ApiUrl +
		"artist_name=" + url.QueryEscape(track.Artist) +
		"&track_name=" + url.QueryEscape(track.Name) +
		"&album_name=" + url.QueryEscape(track.Album) +
		"&track_duration=" + strconv.Itoa(track.Length)

	searchUrl := c.Internal.SearchUrl +
		url.QueryEscape(track.Artist) + "+" +
		url.QueryEscape(track.Name)
	//  url.QueryEscape(track.Album) +
	//  strconv.Itoa(track.Length)

	// local = slices.Contains([]string{"local"}, c.Search.Depth)
	local = true
	match = slices.Contains([]string{"both", "match"}, c.Search.Depth)
	search = slices.Contains([]string{"both", "search"}, c.Search.Depth)

	var lyrics []Lyric
	var status int

	if local {
		lyrics, status = findLocalLyrics(track.Path)
	}

	if match && status != 200 {
		lyrics, status = matchLyrics(matchUrl)
	}

	if search && (len(lyrics) == 1 || len(lyrics) == 0) {
		lyrics, status = searchLyrics(searchUrl)
	}

	if status == 404 {
		NotifyUser(MsgSongNotFound, MsgSongNotFound)
		return ReturnSongNotFound
	}
	if status != 200 {
		NotifyUser(MsgSongNotFound, strings.Join([]string{MsgSongNotFound, " - "}, strconv.Itoa(status)))
		return ReturnSongNotFound
	}

	if lyrics != nil {
		if len(lyrics) == 1 && lyrics[0].Lyric == "" {
			NotifyUser(MsgNoLiveLyrics, MsgNoLiveLyrics)
			return ReturnNoLiveLyrics
		}
	}

	return lyrics
}

func findLocalLyrics(path string) ([]Lyric, int) {
	pathArr := strings.Split(path, ".")
	pathArr = pathArr[:len(pathArr)-1]

	lrcPath := strings.Join(pathArr, ".")
	lrcPath = lrcPath + ".lrc"

	if !FileExists(lrcPath) {
		return nil, 404
	}

	lrcFile := ReadFile(lrcPath)
	lyrics := parseLyrics(lrcFile, "local")

	return lyrics, 200
}

func matchLyrics(url string) ([]Lyric, int) {
	response, e := http.Get(url)
	Check(e)
	defer response.Body.Close()

	body, e := io.ReadAll(response.Body)
	Check(e)

	lyrics := parseLyrics(string(body), "match")

	return lyrics, response.StatusCode
}

func searchLyrics(searchUrl string) ([]Lyric, int) {
	response, e := http.Get(searchUrl)
	Check(e)
	defer response.Body.Close()

	body, e := io.ReadAll(response.Body)
	Check(e)

	lyrics := parseLyrics(string(body), "search")

	return lyrics, response.StatusCode
}

func parseLyrics(request string, mode string) []Lyric {
	var parsedLyrics []Lyric = []Lyric{{"", 0}}

	var decodedLyrics Request
	var rawLyrics []string = []string{""}

	switch mode {
	case "match":
		var decodedLyrics Request
		e := json.Unmarshal([]byte(request), &decodedLyrics)
		Check(e)
		rawLyrics[0] = decodedLyrics.SyncedLyrics

	case "search":
		var decodedResults []Request
		e := json.Unmarshal([]byte(request), &decodedResults)
		Check(e)
		for i := range decodedResults {
			if decodedResults[i].SyncedLyrics != "" {
				decodedLyrics = decodedResults[i]
			}
		}
		rawLyrics[0] = decodedLyrics.SyncedLyrics

	case "local":
		rawLyrics[0] = request
	}

	if rawLyrics[0] == "" {
		return ReturnNoLiveLyrics
	}

	rawLyrics = strings.Split(rawLyrics[0], "\n")
	// Remove trailing newlines
	if rawLyrics[len(rawLyrics)-1] == "" {
		rawLyrics = rawLyrics[:len(rawLyrics)-2]
	}

	for i := range rawLyrics {
		if rawLyrics[i] != "" {
			var lyric Lyric
			sepLyric := strings.Split(rawLyrics[i], "]")

			var position float64
			timestamp := strings.Split(sepLyric[0], "[")[1]
			position = ConvertTimestampToSeconds(timestamp) * 1_000_000

			lyric = Lyric{sepLyric[1], int(position)}
			parsedLyrics = append(parsedLyrics, lyric)
		}
	}

	return parsedLyrics
}
