package fetch

import (
	"io"
	"muse/internal/core"
	"muse/internal/lyric"
	"net/http"
	"net/url"
	"strconv"
)

func FetchFromLRCLIB(track *core.Track) (string, error) {
	var lyrics string

	baseURL := "https://lrclib.net/api/get"

	// Build query parameters
	matchUrl := baseURL + "?" +
		"artist_name=" + url.QueryEscape(track.Artist) +
		"&track_name=" + url.QueryEscape(track.Title) +
		"&album_name=" + url.QueryEscape(track.Album) +
		"&track_duration=" + strconv.Itoa(track.Duration/1_000_000)

	response, err := http.Get(matchUrl)
	if err != nil {
		return "", err
	}
	defer response.Body.Close()

	body, err := io.ReadAll(response.Body)
	if err != nil {
		return "", err
	}

	lyrics, err = lyric.ParseJson(body)
	if err != nil {
		return "", err
	}

	return lyrics, nil
}
