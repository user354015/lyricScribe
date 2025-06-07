package main

import (
	"strings"
	"time"

	"github.com/godbus/dbus/v5"
)

func getActivePlayers() []string {
	conn, e := dbus.ConnectSessionBus()
	Check(e)
	defer conn.Close()

	var dbusObjects []string
	e = conn.BusObject().Call("org.freedesktop.DBus.ListNames", 0).Store(&dbusObjects)
	Check(e)

	var players []string
	for i := range dbusObjects {
		if strings.HasPrefix(dbusObjects[i], PlayerPrefix) {
			players = append(players, dbusObjects[i])
		}
	}

	return players
}

func GetPlayer(pref string) string {
	players := getActivePlayers()

	for i := range players {
		if strings.Contains(players[i], pref) {
			return players[i]
		}
	}
	return "none"
}

func GetPlayerInfo(name string) Song {
	var playerInfo Song

	conn, e := dbus.ConnectSessionBus()
	Check(e)
	defer conn.Close()

	player := conn.Object(name, dbus.ObjectPath(PlayerPath))
	var metadata map[string]dbus.Variant
	e = player.Call("org.freedesktop.DBus.Properties.Get", 0, "org.mpris.MediaPlayer2.Player", "Metadata").Store(&metadata)
	Check(e)

	playerInfo.Album = metadata["xesam:album"].Value().(string)
	playerInfo.Artist = metadata["xesam:artist"].Value().([]string)[0]
	playerInfo.Name = metadata["xesam:title"].Value().(string)
	playerInfo.Length = int(metadata["mpris:length"].Value().(int64) * int64(time.Millisecond))

	return playerInfo
}
