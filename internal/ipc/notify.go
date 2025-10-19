package ipc

import (
	"github.com/godbus/dbus"
)

type Notifier struct {
	conn *dbus.Conn
	name string
}

func NewNotifier(conn *dbus.Conn, appName string) *Notifier {
	return &Notifier{
		conn: conn,
		name: appName,
	}
}

func Notify(n *Notifier, title string, message string) error {

	obj := n.conn.Object("org.freedesktop.Notifications", "/org/freedesktop/Notifications")

	call := obj.Call("org.freedesktop.Notifications.Notify", 0,
		n.name,                    // app_name
		uint32(0),                 //
		"",                        // app_icon
		title,                     // title
		message,                   // body
		[]string{},                // actions
		map[string]dbus.Variant{}, // hints
		int32(2000),               // expire_timeout (in ms)
	)

	return call.Err
}
