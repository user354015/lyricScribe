package main

import (
	"os"
	"path/filepath"

	"github.com/BurntSushi/toml"
)

const (
	PlayerPrefix string = "org.mpris.MediaPlayer2."
	PlayerPath   string = "/org/mpris/MediaPlayer2"

	FetchApiUrl    string = "https://lrclib.net/api/get?"
	FetchSearchUrl string = "https://lrclib.net/api/search?q="
)

type ConfigOptions struct {
	General struct {
		ProgramName string `toml:"program_name"`
		ProgramMode string `toml:"program_mode"`
		Logging     string `toml:"logging"`
	} `toml:"general"`

	Internal struct {
		ApiUrl    string `toml:"api_url"`
		SearchUrl string `toml:"search_url"`
	} `toml:"internal"`

	Search struct {
		Depth string `toml:"depth"`
	} `toml:"search"`

	Player struct {
		Player         string  `toml:"player"`
		PositionOffset float64 `toml:"position_offset"`
		Step           float64 `toml:"step"`
		SilenceTimeout int     `toml:"silence_timeout"`
	} `toml:"player"`
}

var DefaultConfig ConfigOptions
var c ConfigOptions

func SetupDefaultConfig() {
	DefaultConfig.General.ProgramName = "LyricScribe"
	DefaultConfig.General.ProgramMode = "display"
	DefaultConfig.General.Logging = "silent"

	DefaultConfig.Internal.ApiUrl = FetchApiUrl
	DefaultConfig.Internal.SearchUrl = FetchSearchUrl

	DefaultConfig.Search.Depth = "both"

	DefaultConfig.Player.Player = "mpv"
	DefaultConfig.Player.PositionOffset = -0.52
	DefaultConfig.Player.Step = 0.3
	DefaultConfig.Player.SilenceTimeout = 3
}

func ReadConfig() {
	configPath, e := os.UserConfigDir()
	Check(e)
	configPath = filepath.Join(configPath, "lyrics", "config.toml")

	config := &ConfigOptions{}
	*config = DefaultConfig

	_, e = toml.DecodeFile(configPath, config)
	Check(e)

	c = *config
}
