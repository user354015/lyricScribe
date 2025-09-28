package util

import (
	"errors"
	"strconv"
	"strings"
)

func TimestampToSeconds(timestamp string) (float64, error) {
	parts := strings.Split(timestamp, ":")
	if len(parts) != 2 {
		return 0, errors.New("Invalid timestamp format")
	}

	minutes, err := strconv.ParseFloat(parts[0], 64)
	if err != nil {
		return 0, errors.New("Invalid timestamp format")
	}

	seconds, err := strconv.ParseFloat(parts[1], 64)
	if err != nil {
		return 0, errors.New("Invalid timestamp format")
	}

	return minutes*60 + seconds*1_000_000, nil
}
