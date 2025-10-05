package lyric

import (
	"encoding/json"
	"muse/internal/core"
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

func ParseLrc(lrcFile string) (*[]core.Lyric, error) {
	var rawLyrics = strings.Split(lrcFile, "\n")
	var parsedLyrics []core.Lyric = []core.Lyric{{}}

	if rawLyrics != nil {
	}

	for i := range rawLyrics {
		if rawLyrics[i] != "" {
			var lyric core.Lyric
			sepLyric := strings.Split(rawLyrics[i], "]")

			var position float64
			timestamp := strings.Split(sepLyric[0], "[")[1]
			position, err := util.TimestampToSeconds(timestamp)
			if err != nil {
				return nil, err
			}

			lyric = core.Lyric{Lyric: sepLyric[1], Position: int(position)}
			lyric.Lyric = strings.TrimSpace(lyric.Lyric)
			parsedLyrics = append(parsedLyrics, lyric)
		}
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
