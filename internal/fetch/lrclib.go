package fetch

import (
	"io"
	"muse/internal/lyric"
	"muse/internal/shared"
	"net/http"
	"net/url"
	"strconv"
)

func FetchFromLRCLIB(track *shared.Track) (string, error) {
	var lyrics string = ""
	var status int = 0
	var err error

	lyrics, status, err = MatchLRCLIB(track)
	if status == 200 {
		return lyrics, err
	}

	lyrics, status, err = SearchLRCLIB(track)
	if status == 200 {
		return lyrics, err
	}

	return lyrics, err
}

func MatchLRCLIB(track *shared.Track) (string, int, error) {
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
		return "", response.StatusCode, err
	}
	defer response.Body.Close()

	body, err := io.ReadAll(response.Body)
	if err != nil {
		return "", response.StatusCode, err
	}

	lyrics, err = lyric.ParseJson(body)
	if err != nil {
		return "", response.StatusCode, err
	}

	return lyrics, 200, nil
}

func SearchLRCLIB(track *shared.Track) (string, int, error) {
	var lyrics string

	baseURL := "https://lrclib.net/api/search"

	params := url.Values{}
	params.Add("q", track.Artist+" "+track.Title)

	searchUrl := baseURL + "?" + params.Encode()
	shared.Debug(searchUrl)

	response, err := http.Get(searchUrl)
	if err != nil {
		return "", response.StatusCode, err
	}
	defer response.Body.Close()

	body, err := io.ReadAll(response.Body)
	if err != nil {
		return "", response.StatusCode, err
	}

	lyrics, err = lyric.ParseJsonArr(body)
	if err != nil {
		return "", response.StatusCode, err
	}

	return lyrics, 200, nil
}
