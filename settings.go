package main

const (
	ProgramName string = "LyricScribe"

	ProgramMode string = "display" // Available modes: debug, display
	LoggingMode string = "silent"  // Available modes: silent, display

	FetchApiUrl    string = "https://lrclib.net/api/get?"
	FetchSearchUrl string = "https://lrclib.net/api/search?q="

	PreferedPlayer string = "tauon"
	// PreferedPlayer string = "chromium.instance934"
	PlayerPrefix string = "org.mpris.MediaPlayer2."
	PlayerPath   string = "/org/mpris/MediaPlayer2"

	PositionOffset float64 = -0.52
	Step           float64 = 0.05
	SilenceTimout  int     = 3
)
