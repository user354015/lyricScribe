package main

import "fmt"

var testSong Song = Song{
	"Stained, Brutal Calamity",
	// "Threats of the Ocean Floor",
	"DM DOKURO",
	"The Tale of a Cruel World (Calamity Original Soundtrack)",
	733363645}

func main() {
	lyrics := DownloadLyrics(testSong)
	fmt.Println(lyrics[0])
}
