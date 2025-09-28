package main

import (
	"fmt"

	"muse/internal/core"
	"muse/internal/dbus"
)

func main() {

	conn, _ := dbus.Connect()
	player, _ := dbus.FindActivePlayer(conn)
	track, _ := dbus.GetTrackInfo(conn, player)

	dbus.WatchTrackChanges(conn, player, func(item *core.Track) {
		track = item
	})

	for true {
		fmt.Println(track.Title)
	}
}
