package main

import (
	"time"

	"muse/internal/core"
	"muse/internal/dbus"
	"muse/internal/display"
	"muse/internal/fetch"
	"muse/internal/lyric"
)

func main() {

	conn, err := dbus.Connect()
	if err != nil {
		panic("could not establish dbus connection")
	}
	player, err := dbus.FindActivePlayer(conn)
	if err != nil {
		panic(err)
	}
	track, err := dbus.GetTrackInfo(conn, player)
	if err != nil {
		panic(err)
	}

	var lastLine string
	track, lyrics := updateTrackInfo(track)

	dbus.WatchTrackChanges(conn, player, func(item *core.Track) {
		track, lyrics = updateTrackInfo(item)
		lastLine = ""
	})

	for {
		pos, err := dbus.GetPlayerPosition(conn, player)
		if err != nil {
			panic(core.ErrNoPlayerPos)
		}

		idx := core.GetCurrentLine(lyrics, pos)
		text := "…"
		if idx < len(lyrics) && len(lyrics) > 0 {
			text = lyrics[idx].Lyric
		}

		if text != lastLine {
			display.Display(text)
			lastLine = text
		}

		time.Sleep(0_500_000_000)

	}
}

func updateTrackInfo(track *core.Track) (*core.Track, []core.Lyric) {
	rawLyrs, err := fetch.FetchLyrics(track)
	if err != nil {
		if err == core.ErrNoLyricsFound {
			return track, nil
		}
		panic(err)
	}

	lyrics, err := lyric.ParseLrc(rawLyrs)
	if err != nil {
		panic(err)
	}

	return track, *lyrics
}
