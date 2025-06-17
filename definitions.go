package main

type Song struct {
	Name   string
	Artist string
	Album  string
	Length int
}

type Lyric struct {
	Lyric    string
	Position int
}

type Request struct {
	Id           int     `json:"id"`
	Name         string  `json:"name"`
	TrackName    string  `json:"trackName"`
	ArtistName   string  `json:"artistName"`
	AlbumName    string  `json:"albumName"`
	Duration     float32 `json:"duration"`
	Instrumental bool    `json:"instrumental"`
	PlainLyrics  string  `json:"plainLyrics"`
	SyncedLyrics string  `json:"syncedLyrics"`
}

var (
	MsgSongNotFound []Lyric
	MsgNoLiveLyrics []Lyric

	MsgNoActivePlayer string
)

func SetLogMessages() {
	switch LoggingMode {
	case "silent":
		MsgSongNotFound = []Lyric{{"", 0}}
		MsgNoLiveLyrics = []Lyric{{"", 0}}
		MsgNoActivePlayer = "No MPRIS-compatible players active"
	case "display":
		MsgSongNotFound = []Lyric{{"Song not found", 0}}
		MsgNoLiveLyrics = []Lyric{{"No live lyrics found", 0}}
		MsgNoActivePlayer = "No MPRIS-compatible players active"
	}
}
