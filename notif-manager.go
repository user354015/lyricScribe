package main

import "github.com/godbus/dbus/v5"

func NotifyUser(title string, message string) error {
	obj := DbusConn.Object("org.freedesktop.Notifications", "/org/freedesktop/Notifications")

	call := obj.Call("org.freedesktop.Notifications.Notify", 0,
		ProgramName,               // app_name
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
