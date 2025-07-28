package main

import (
	"strconv"
	"strings"

	"github.com/godbus/dbus/v5"
)

func Check(e error) {
	if e != nil {
		panic(e)
	}
}

var DbusConn *dbus.Conn

func InitDBusConnection() {
	var err error
	DbusConn, err = dbus.SessionBus()
	Check(err)
}

func CloseDBusConnection() {
	if DbusConn != nil {
		DbusConn.Close()
	}
}

func ConvertTimestampToSeconds(timestamp string) float64 {
	time := strings.Split(timestamp, ":")
	minutes, e := strconv.ParseFloat(time[0], 64)
	Check(e)
	seconds, e := strconv.ParseFloat(time[1], 64)
	Check(e)

	return minutes*60 + seconds
}

func ComparePositions(position int, positions []int, expectedPos int) int {
	if len(positions) == 0 {
		return 0
	}

	for i := expectedPos; i < len(positions); i++ {
		if positions[i] > position && i > 1 {
			return i - 1
		}
	}

	return 0
}
