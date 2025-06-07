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

	return minutes + seconds*60
}
