package main

import (
	"encoding/json"
	"io"
	"net/http"
	"sort"
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

// func ComparePositions(position int, positions []int) int {
// 	for i := range positions {
// 		if positions[i] > position {
// 			return i
// 		}
// 	}
// 	return 0
// }

// func ComparePositions(position int, positions []int) int {
// 	// Find the last lyric whose timestamp has passed (current lyric)
// 	currentIndex := 0
// 	for i := range positions {
// 		if positions[i] <= position {
// 			currentIndex = i
// 		} else {
// 			break
// 		}
// 	}
// 	if currentIndex <= len(positions) {
// 		return currentIndex
// 	} else {
// 		return 0
// 	}
// }

func ComparePositions(position int, positions []int) int {
	if len(positions) == 0 {
		return 0
	}

	// Find the first index where positions[i] > position
	i := sort.Search(len(positions), func(i int) bool {
		return positions[i] > position
	})

	// Get the first lyric behind current position
	if i != 0 {
		return i - 1
	} else {
		return 0
	}
}

// https://api.github.com/repos/user354015/lyricScribe/releases/latest
func IsLatestVers(version string, repo string) bool {
	response, e := http.Get(repo)
	Check(e)
	defer response.Body.Close()

	body, e := io.ReadAll(response.Body)
	Check(e)

	var repoData struct {
		Version string `json:"tag_name"`
	}

	e = json.Unmarshal(body, &repoData)
	Check(e)

	if version == repoData.Version {
		return true
	} else {
		return false
	}
}
