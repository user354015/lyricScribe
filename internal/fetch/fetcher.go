package fetch

import (
	"muse/internal/core"
	"muse/internal/util"
)

func FetchLyrics(track *core.Track) (lyrics string, err error) {
	lyrics, err = FetchLocalLyrics(track)
	if lyrics != "" && err == nil {
		return lyrics, nil
	}

	lyrics, err = FetchFromLRCLIB(track)
	if lyrics != "" && err == nil {
		return lyrics, nil
	}

	return lyrics, core.ErrNoLyricsFound
}

func FetchLocalLyrics(track *core.Track) (string, error) {
	var lyrics string
	location, err := util.ReplaceExtension(track.Location, "lrc")
	if err != nil {
		return "", nil
	}

	if util.FileExists(location) {
		lyrics, err = util.ReadFile(location)
		if err != nil {
			return "", err
		}
	}

	return lyrics, nil
}
