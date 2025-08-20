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

func DownloadLyrics(track Song) []Lyric {
	matchUrl := FetchApiUrl +
		"artist_name=" + url.QueryEscape(track.Artist) +
		"&track_name=" + url.QueryEscape(track.Name) +
		"&album_name=" + url.QueryEscape(track.Album) +
		"&track_duration=" + strconv.Itoa(track.Length)

	searchUrl := FetchSearchUrl +
		url.QueryEscape(track.Artist) + "+" +
		url.QueryEscape(track.Name)
	//  url.QueryEscape(track.Album) +
	//  strconv.Itoa(track.Length)

	var lyrics []Lyric

	match := slices.Contains([]string{"both", "match"}, c.Search.Depth)
	search := slices.Contains([]string{"both", "search"}, c.Search.Depth)

	var status int

	if match {
		lyrics, status = findLyrics(matchUrl)
	}

	if search && len(lyrics) == 1 {
		lyrics, _ = searchLyrics(searchUrl)
	}

	if status == 404 {
		NotifyUser(MsgSongNotFound, MsgSongNotFound)
		return ReturnSongNotFound
	}
	if status != 200 {
		NotifyUser(MsgSongNotFound, strings.Join([]string{MsgSongNotFound, " - "}, strconv.Itoa(status)))
		return ReturnSongNotFound
	}

	if len(lyrics) == 1 {
		if lyrics[0].Lyric == "" {
			lyrics[0].Lyric = MsgNoLiveLyrics
		}
	}

	if lyrics != nil {
		if lyrics[0].Lyric == MsgNoLiveLyrics {
			NotifyUser(MsgNoLiveLyrics, MsgNoLiveLyrics)
		}
	}

	return lyrics
}

func findLyrics(url string) ([]Lyric, int) {
	response, e := http.Get(url)
	Check(e)
	defer response.Body.Close()

	body, e := io.ReadAll(response.Body)
	Check(e)

	lyrics := parseLyrics(string(body), false)

	return lyrics, response.StatusCode
}

func searchLyrics(searchUrl string) ([]Lyric, int) {
	response, e := http.Get(searchUrl)
	Check(e)
	defer response.Body.Close()

	body, e := io.ReadAll(response.Body)
	Check(e)

	lyrics := parseLyrics(string(body), true)

	return lyrics, response.StatusCode
}

func parseLyrics(request string, search bool) []Lyric {
	var parsedLyrics []Lyric = []Lyric{{"", 0}}

	var decodedLyrics Request
	if !search {
		e := json.Unmarshal([]byte(request), &decodedLyrics)
		Check(e)
	} else {
		var decodedResults []Request
		e := json.Unmarshal([]byte(request), &decodedResults)
		Check(e)
		for i := range decodedResults {
			if decodedResults[i].SyncedLyrics != "" {
				decodedLyrics = decodedResults[i]
			}
		}
	}

	if decodedLyrics.SyncedLyrics == "" {
		return ReturnNoLiveLyrics
	}

	rawLyrics := strings.Split(decodedLyrics.SyncedLyrics, "\n")
	for i := range rawLyrics {
		var lyric Lyric
		sepLyric := strings.Split(rawLyrics[i], "]")

		var position float64
		timestamp := strings.Split(sepLyric[0], "[")[1]
		position = ConvertTimestampToSeconds(timestamp) * 1_000_000

		lyric = Lyric{sepLyric[1], int(position)}
		parsedLyrics = append(parsedLyrics, lyric)
	}

	return parsedLyrics
}
