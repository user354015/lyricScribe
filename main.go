package main

import (
	"time"

	"muse/internal/core"
	"muse/internal/dbus"
	"muse/internal/display"
	"muse/internal/lyrics"
)

func main() {

	conn, _ := dbus.Connect()
	player, _ := dbus.FindActivePlayer(conn)
	track, _ := dbus.GetTrackInfo(conn, player)

	dbus.WatchTrackChanges(conn, player, func(item *core.Track) {
		track = item
	})

	rawLyrs := `[00:31.71] No more holding back
[00:37.03] They'll wish for demise
[00:42.01] They thought they were invincible
[00:47.07] But you could see through their hollow lies
[00:52.99] You've seen how far they've come
[00:58.19] How many they've slain
[01:03.39] You will gladly tear their limbs apart
[01:08.42] All before they start to kill again
[01:16.96]`
	l, err := lyrics.ParseLrc(rawLyrs)
	lyrics := *l
	if err != nil {
		panic(err)
	}

	for true {
		display.Display(track.Title)
		display.Display(lyrics[2].Lyric)
		time.Sleep(2_000_000_000)
	}
}
