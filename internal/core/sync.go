package core

import (
	"sort"
)

func GetCurrentLine(lyric []Lyric, pos int) int {
	var positions []int
	for i := range lyric {
		positions = append(positions, lyric[i].Position)
	}

	if len(positions) == 0 {
		return 0
	}

	// find the first index where positions[i] > position
	i := sort.Search(len(positions), func(i int) bool {
		return positions[i] > pos
	})

	// get the first lyric behind current position
	if i != 0 {
		return i - 1
	} else {
		return 0
	}
}
