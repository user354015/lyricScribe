package main

import (
	"errors"
	"fmt"
	"time"
)

// var testSong Song = Song{
// 	"Stained, Brutal Calamity",
// 	// "Threats of the Ocean Floor",
// 	"DM DOKURO",
// 	"The Tale of a Cruel World (Calamity Original Soundtrack)",
// 	733363645}

var player string
var lyrics []Lyric
var currentSong Song
var positions []int

var prevText string

func setup() {
	_ = ReadConfig()

	player = GetPlayer(c.Player.Player)
	if player == "none" {
		Check(errors.New(ReturnNoActivePlayer))
	}

	currentSong = GetPlayerInfo(player)
	lyrics = DownloadLyrics(currentSong)

	positions = []int{}
	for i := range lyrics {
		positions = append(positions, lyrics[i].Position)
	}
}

func displayText(text string) {

	switch c.General.ProgramMode {
	case "debug":
		if text != prevText {
			fmt.Printf("%v\n", text)
			prevText = text
		}
	case "display":
		program.UpdateDisplay(text)
	}
}

var mode string = "display"
var program *Program

func main() {
	ReadConfig()
	mode = c.General.ProgramMode

	switch mode {
	case "debug":
	case "display":
		program = NewProgram()
		defer program.Quit()
	}

	InitDBusConnection()
	defer CloseDBusConnection()

	SetLogMessages()

	for true {
		setup()

		playerInfo := GetPlayerInfo(player)
		for currentSong == playerInfo {
			playerInfo = GetPlayerInfo(player)
			position := GetPlayerPosition(player)
			id := ComparePositions(position-int(c.Player.PositionOffset*1_000_000), positions)

			var text string
			if positions[id]-position < c.Player.SilenceTimeout*1_000_000 {
				text = lyrics[id].Lyric
			} else {
				text = ""
			}

			displayText(text)

			time.Sleep(time.Duration(c.Player.Step * 1_000_000))
		}

		// Small cooldown to prevent spamming the api
		// time.Sleep(time.Duration(1_500 * 1_000_000))
	}
}
