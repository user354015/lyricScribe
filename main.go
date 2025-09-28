package main

import (
	"fmt"
	"muse/internal/dbus"
)

func main() {
	conn, err := dbus.Connect()

	player, err := dbus.FindActivePlayer(conn)
	if err != nil {
		panic(err)
	}
	track, err := dbus.GetTrackInfo(conn, player)

	fmt.Println(track.Title)

}
