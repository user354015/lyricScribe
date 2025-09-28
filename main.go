package main

import (
	"fmt"
	"os"
	"os/signal"

	"muse/internal/core"
	"muse/internal/dbus"
)

func main() {
	conn, _ := dbus.Connect()
	player, _ := dbus.FindActivePlayer(conn)

	// Just pass the callback, signal handling is hidden
	dbus.WatchTrackChanges(conn, player, func(track *core.Track) {
		fmt.Println(track)
	})

	// Keep alive
	sigChan := make(chan os.Signal, 1)
	signal.Notify(sigChan, os.Interrupt)
	<-sigChan
}
