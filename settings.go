package main

const (
	ProgramName string = "LyricScribe"

	ProgramMode string = "display" // Available modes: debug, display
	LoggingMode string = "silent"  // Available modes: silent, display

	FetchMode string = "both" // Available modes: match, search, both

	FetchApiUrl    string = "https://lrclib.net/api/get?"
	FetchSearchUrl string = "https://lrclib.net/api/search?q="

	PreferedPlayer string = "tauon"
	// PreferedPlayer string = "chromim.instance"

	PlayerPrefix string = "org.mpris.MediaPlayer2."
	PlayerPath   string = "/org/mpris/MediaPlayer2"

	PositionOffset float64 = -0.52
	Step           float64 = 0.33
	SilenceTimout  int     = 3
)
