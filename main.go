package main

import (
	"muse/internal/config"
	"muse/internal/core"
)

func main() {

	cfg, err := config.Load("~/.config/muse/config.toml")
	if err != nil {
		panic(err)
	}
	app := core.NewApp(cfg)
	if err := app.Start(); err != nil {
		panic(err)
	}
}
