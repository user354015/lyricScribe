package dbus

import (
	"net/url"
	"strings"

	"muse/internal/config"
	"muse/internal/core"
	"muse/internal/util"

	"github.com/godbus/dbus"
)

func Connect() (dbusConn *dbus.Conn, err error) {
	dbusConn, err = dbus.SessionBus()
	return
}

func FindActivePlayer(conn *dbus.Conn, cfg *config.Config) (string, error) {
	var dbusObjects []string
	err := conn.BusObject().Call("org.freedesktop.DBus.ListNames", 0).Store(&dbusObjects)
	if err != nil {
		return "", err
	}

	var players []string
	for i := range dbusObjects {
		if strings.HasPrefix(dbusObjects[i], "org.mpris.MediaPlayer2.") {
			players = append(players, dbusObjects[i])
		}
	}

	if len(players) == 0 {
		return "", core.ErrNoActivePlayers
	}

	for pref := range cfg.Player.Preferred {
		for i := range players {
			if strings.Contains(players[i], cfg.Player.Preferred[pref]) {
				return players[i], nil
			}
		}
	}

	return players[0], err
}

func GetTrackInfo(conn *dbus.Conn, playerService string) (*core.Track, error) {
	obj := conn.Object(playerService, "/org/mpris/MediaPlayer2")
	call := obj.Call("org.freedesktop.DBus.Properties.Get", 0,
		"org.mpris.MediaPlayer2.Player", "Metadata")

	var metadata map[string]dbus.Variant
	err := call.Store(&metadata)
	if err != nil {
		return nil, err
	}

	var trackInfo *core.Track

	trackInfo = &core.Track{
		Album:    metadata["xesam:album"].Value().(string),
		Artist:   metadata["xesam:artist"].Value().([]string)[0],
		Title:    metadata["xesam:title"].Value().(string),
		Duration: int64(metadata["mpris:length"].Value().(int64)),
	}

	path := metadata["xesam:url"].Value().(string)
	location := ""

	if path != "" {
		path, err := url.PathUnescape(path)
		path, ok := strings.CutPrefix(path, "file://")

		if err != nil || !ok {
			return nil, err
		}

		if util.FileExists(path) {
			location = path
		}
	}

	trackInfo.Location = location

	return trackInfo, err
}

func WatchTrackChanges(conn *dbus.Conn, playerService string, callback func(*core.Track)) {

	signalChan := make(chan *dbus.Signal, 10)
	conn.Signal(signalChan)

	matchRule := `type='signal', interface='org.freedesktop.DBus.Properties',
		member='PropertiesChanged', path='/org/mpris/MediaPlayer2', sender='` +
		playerService + `'`

	conn.BusObject().Call("org.freedesktop.DBus.AddMatch", 0, matchRule)

	go func() {
		for signal := range signalChan {
			if signal.Name == "org.freedesktop.DBus.Properties.PropertiesChanged" && len(signal.Body) >= 2 {
				// Check if Metadata changed
				if changedProps, ok := signal.Body[1].(map[string]dbus.Variant); ok {
					if _, hasMetadata := changedProps["Metadata"]; hasMetadata {
						track, err := GetTrackInfo(conn, playerService)
						if err == nil {
							callback(track)
						}
					}
				}
			}
		}
	}()
}

func GetPlayerPosition(conn *dbus.Conn, playerPath string) (int, error) {
	player := conn.Object(playerPath, "/org/mpris/MediaPlayer2")
	var position int
	err := player.Call("org.freedesktop.DBus.Properties.Get", 0,
		"org.mpris.MediaPlayer2.Player", "Position").Store(&position)

	return position, err
}
