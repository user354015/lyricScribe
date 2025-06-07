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

func main() {
	// lyrics := DownloadLyrics(testSong)
	player := GetPlayer(PreferedPlayer)
	if player == "none" {
		Check(errors.New(MsgNoActivePlayer))
	}
	playerInfo := GetPlayerInfo(player)
	fmt.Println(playerInfo)

}
