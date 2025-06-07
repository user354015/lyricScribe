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
		Check(errors.New(MsgNoActivePlayer))
	}

	currentSong = GetPlayerInfo(player)
	lyrics = DownloadLyrics(currentSong)

	for i := range lyrics {
		positions = append(positions, lyrics[i].Position)
	}
}

func main() {
	for true {
		setup()

		for currentSong == GetPlayerInfo(player) {
			position := GetPlayerPosition(player)
			id := ComparePositions(position-int(PositionOffset*100_000), positions)
			lyric := lyrics[id].Lyric

			fmt.Printf("lyric: %v\n", lyric)
		}
	}
}
