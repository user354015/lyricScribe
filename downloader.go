package main

import (
	"encoding/json"
	"io"
	"net/http"
	"net/url"
	"strconv"
	"strings"
)

func DownloadLyrics(track Song) []Lyric {
	url := FetchApiUrl +
		"artist_name=" + url.QueryEscape(track.Artist) +
		"&track_name=" + url.QueryEscape(track.Name) +
		"&album_name=" + url.QueryEscape(track.Album) +
		"&track_duration=" + strconv.Itoa(track.Length)

	response, e := http.Get(url)
	Check(e)
	defer response.Body.Close()

	if response.StatusCode == 404 {
		// fmt.Println(url)
		NotifyUser(MsgSongNotFound, MsgSongNotFound)
		return ReturnSongNotFound
	}
	if response.StatusCode != 200 {
		NotifyUser(MsgSongNotFound, strings.Join([]string{MsgSongNotFound, ": "}, strconv.Itoa(response.StatusCode)))
		return ReturnSongNotFound
	}

	body, e := io.ReadAll(response.Body)
	Check(e)

	lyrics := parseLyrics(string(body))

	return lyrics
}

func parseLyrics(request string) []Lyric {
	var parsedLyrics []Lyric

	var decodedLyrics Request
	e := json.Unmarshal([]byte(request), &decodedLyrics)
	Check(e)

	if decodedLyrics.SyncedLyrics == "" {
		NotifyUser(MsgNoLiveLyrics, MsgNoLiveLyrics)
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
