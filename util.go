package main

import (
	"strconv"
	"strings"
)

func Check(e error) {
	if e != nil {
		panic(e)
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

func ComparePositions(position int, positions []int) int {
	// Find the last lyric whose timestamp has passed (current lyric)
	currentIndex := 0
	for i := range positions {
		if positions[i] <= position {
			currentIndex = i
		} else {
			break
		}
	}
	if currentIndex <= len(positions) {
		return currentIndex
	} else {
		return 0
	}
}
