package main

const Version string = "v0.3"
const TargetRepo string = "https://api.github.com/repos/user354015/lyricScribe/releases/latest"

type Song struct {
	Name   string
	Artist string
	Album  string
	Length int
	Path   string
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
	ReturnSongNotFound []Lyric
	ReturnNoLiveLyrics []Lyric

	ReturnNoActivePlayer string
)

const (
	MsgSongNotFound   string = "Song not found"
	MsgNoLiveLyrics   string = "No live lyrics found"
	MsgNoActivePlayer string = "No MPRIS-compatible players are active"
)

func SetLogMessages() {
	switch c.General.Logging {
	case "silent":
		ReturnSongNotFound = []Lyric{{"", 0}}
		ReturnNoLiveLyrics = []Lyric{{"", 0}}
		ReturnNoActivePlayer = "No MPRIS-compatible players active"
	case "display":
		ReturnSongNotFound = []Lyric{{"Song not found", 0}}
		ReturnNoLiveLyrics = []Lyric{{"No live lyrics found", 0}}
		ReturnNoActivePlayer = "No MPRIS-compatible players active"
	}

}
