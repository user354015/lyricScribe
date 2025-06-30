package main

import (
	"errors"
	"fmt"
)

var testSong Song = Song{
	"Stained, Brutal Calamity",
	// "Threats of the Ocean Floor",
	"DM DOKURO",
	"The Tale of a Cruel World (Calamity Original Soundtrack)",
	733363645}

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

	SetLogMessages()

	for true {
		setup()

		for currentSong == GetPlayerInfo(player) {
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
		}
	}
}
