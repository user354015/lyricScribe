package main

import (
	"muse/internal/config"
	"muse/internal/core"
	"muse/internal/shared"
)

func main() {

	shared.InitLogger(false)

	cfg, err := config.Load("/home/unicorn/.config/muse/config.toml")
	if err != nil {
		panic(err)
	}
	app := core.NewApp(cfg)
	if err := app.Start(); err != nil {
		panic(err)
	}
}
