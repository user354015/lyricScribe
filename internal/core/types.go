package core

type Track struct {
	Title    string
	Artist   string
	Album    string
	Duration int64
	Location string
}

type Lyric struct {
	Lyric    string
	Position int
}
