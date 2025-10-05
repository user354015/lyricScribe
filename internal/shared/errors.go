package shared

import "errors"

var (
	ErrNoLyricsFound = errors.New("no lyrics found")
	ErrNoPlayerPos   = errors.New("failed to get player's position")

	ErrInvalidTimestamp = errors.New("invalid timestamp format")

	ErrNoActivePlayers = errors.New("no active mpris-compatible media players found")
)
