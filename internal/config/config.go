package config

import (
	"muse/internal/util"

	"github.com/BurntSushi/toml"
)

type Config struct {
	General  General  `toml:"general"`
	Player   Player   `toml:"player"`
	Fetching Fetching `toml:"fetching"`
	Display  Display  `toml:"display"`
}

type General struct {
	ProgramName string `toml:"program_name"`
	DisplayMode string `toml:"display_mode"`
	Updates     bool   `toml:"check_for_updates"`
}

type Player struct {
	Preferred      []string `toml:"preferred"`
	PositionOffset float32  `toml:"position_offset"` // in seconds
	PollInterval   int      `toml:"poll_interval"`   // in milliseconds
	SilenceTimeout int      `toml:"silence_timeout"` // in milliseconds
}

type Fetching struct {
}

type Display struct {
	FgColor  string `toml:"foreground"`
	BgColor  string `toml:"background"`
	Font     string `toml:"font"`
	FontSize int    `toml:"font_size"`

	WindowX int `toml:"window_x"`
	WindowY int `toml:"window_y"`
	WindowW int `toml:"window_width"`
	WindowH int `toml:"window_height"`
}

func Default() *Config {
	var c Config

	c.General.ProgramName = "muse"
	c.General.DisplayMode = "simple"
	c.General.Updates = true

	c.Player.Preferred = []string{"tauon", "mpv", "spotify", "chromium"}
	c.Player.PositionOffset = -0.52
	c.Player.PollInterval = 250
	c.Player.SilenceTimeout = 3000

	c.Display.FgColor = "#ffffff"
	c.Display.BgColor = "#000000"
	c.Display.Font = "/usr/share/fonts/TTF/Hack-Regular.ttf"
	c.Display.FontSize = 32
	c.Display.WindowX = 410
	c.Display.WindowY = 0
	c.Display.WindowH = 1100
	c.Display.WindowW = 250

	return &c
}

func Load(path string) (*Config, error) {
	cfg := Default()

	if !util.FileExists(path) {
		return cfg, nil
	}

	if _, err := toml.DecodeFile(path, cfg); err != nil {
		return nil, err
	}

	return cfg, nil

}
