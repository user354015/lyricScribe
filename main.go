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

	conn, _ := dbus.Connect()
	player, _ := dbus.FindActivePlayer(conn)
	track, _ := dbus.GetTrackInfo(conn, player)

	track, lyrics := updateTrackInfo(track)

	dbus.WatchTrackChanges(conn, player, func(item *core.Track) {
		track, lyrics = updateTrackInfo(item)
	})

	for true {
		pos, _ := dbus.GetPlayerPosition(conn, player)
		idx := core.GetCurrentLine(lyrics, pos)
		text := lyrics[idx].Lyric

		display.Display(text)
		time.Sleep(0_500_000_000)

	}
}

func updateTrackInfo(track *core.Track) (*core.Track, []core.Lyric) {
	rawLyrs, err := fetch.FetchLyrics(track)
	if err != nil {
		panic(err)
	}

	lyrics, err := lyric.ParseLrc(rawLyrs)
	if err != nil {
		panic(err)
	}

	return track, *lyrics
}
