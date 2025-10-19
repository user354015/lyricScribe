package display

import "strings"

type TextFormatter struct {
	MaxWidth int
}

func NewLyricFormatter(maxWidth int) *TextFormatter {
	return &TextFormatter{MaxWidth: maxWidth}
}

func (f *TextFormatter) WrapText(text string) []string {
	size := len(text)
	if size <= f.MaxWidth {
		return []string{text}
	}

	words := strings.Fields(text)
	if len(words) == 0 {
		return []string{}
	}

	linesN := int(size/f.MaxWidth) + 1
	if linesN < 0 {
		linesN = 0
	}

	// var lnSize = linesN
	var lines []string
	selector := len(words) / linesN
	for i := 1; i < linesN+1; i++ {
		start := selector * (i - 1)
		end := selector * i
		if i == linesN {
			end = len(words)
		}
		line := strings.Join(words[start:end], " ")
		lines = append(lines, line)
	}

	return lines
}
