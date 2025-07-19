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

func setup() {
	player = GetPlayer(PreferedPlayer)
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

var mode string = ProgramMode

func main() {
	var prevText string
	var program *Program

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
			id := ComparePositions(position-int(PositionOffset*1_000_000), positions)

			var text string
			if positions[id]-position < SilenceTimout*1_000_000 {
				text = lyrics[id].Lyric
			} else {
				text = ""
			}

			switch mode {
			case "debug":
				if text != prevText {
					fmt.Printf("%v\n", text)
					prevText = text
				}
			case "display":
				program.UpdateDisplay(text)
			}

			time.Sleep(time.Duration(Step * 1_000_000))
		}

		// Small cooldown to prevent spamming the api
		time.Sleep(time.Duration(1_500 * 1_000_000))
	}
}
