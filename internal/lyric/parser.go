package lyric

import (
	"encoding/json"
	"muse/internal/shared"
	"muse/internal/util"
	"strings"
)

type lrclibTrack struct {
	ID           int     `json:"id"`
	Name         string  `json:"name"`
	TrackName    string  `json:"trackName"`
	ArtistName   string  `json:"artistName"`
	AlbumName    string  `json:"albumName"`
	Duration     float64 `json:"duration"`
	Instrumental bool    `json:"instrumental"`
	PlainLyrics  string  `json:"plainLyrics"`
	SyncedLyrics string  `json:"syncedLyrics"`
}

func ParseLrc(lrcFile string) (*[]shared.Lyric, error) {
	var rawLyrics = strings.Split(lrcFile, "\n")
	var parsedLyrics []shared.Lyric

	if rawLyrics != nil {
	}

	for i := range rawLyrics {
		if rawLyrics[i] != "" {
			var lyric shared.Lyric
			sepLyric := strings.Split(rawLyrics[i], "]")

			var position float64
			timestamp := strings.Split(sepLyric[0], "[")[1]
			position, err := util.TimestampToSeconds(timestamp)
			if err != nil {
				return nil, err
			}

			lyric = shared.Lyric{Lyric: sepLyric[1], Position: int(position)}
			lyric.Lyric = strings.TrimSpace(lyric.Lyric)
			parsedLyrics = append(parsedLyrics, lyric)
		}
	}

	if parsedLyrics[0].Position != 0 {
		parsedLyrics = append([]shared.Lyric{{Lyric: "", Position: 0}}, parsedLyrics...)
	}

	return &parsedLyrics, nil
}

func ParseJson(jsonData []byte) (string, error) {
	var result lrclibTrack
	err := json.Unmarshal(jsonData, &result)
	if err != nil {
		return "", err
	}

	return result.SyncedLyrics, err
}
