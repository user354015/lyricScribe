package lyrics

import (
	"muse/internal/core"
	"muse/internal/util"
	"strings"
)

func ParseLrc(lrcFile string) (*[]core.Lyric, error) {
	var rawLyrics = strings.Split(lrcFile, "\n")
	var parsedLyrics []core.Lyric = []core.Lyric{}

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
