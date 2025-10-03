package util

import (
	"errors"
	"os"
	"path/filepath"
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

	return (minutes*60 + seconds) * 1_000_000, nil
}

func FileExists(p string) bool {
	if _, err := os.Stat(p); err == nil {
		return true
	}

	return false
}

func ReadFile(p string) (string, error) {
	file, err := os.ReadFile(p)
	if err != nil {
		return "", err
	}
	return string(file), nil

}

func ReplaceExtension(p string, newExtension string) (string, error) {
	if !FileExists(p) {
		return "", errors.New("file does not exist")
	}

	ext := filepath.Ext(p)
	base := p[:len(p)-len(ext)]
	if newExtension[0] != '.' {
		newExtension = "." + newExtension
	}
	return base + newExtension, nil
}
