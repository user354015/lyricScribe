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

	params := url.Values{}
	params.Add("artist_name", track.Artist)
	params.Add("track_name", track.Title)
	params.Add("album_name", track.Album)
	params.Add("track_duration", strconv.Itoa(int(track.Duration/1_000_000)))

	matchUrl := baseURL + "?" + params.Encode()

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
