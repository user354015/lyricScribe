package util

import (
	"image/color"
	"strconv"
	"strings"
)

func HexToRGBA(hex string) color.RGBA {
	hex = strings.TrimPrefix(hex, "#")

	r, _ := strconv.ParseUint(hex[0:2], 16, 8)
	g, _ := strconv.ParseUint(hex[2:4], 16, 8)
	b, _ := strconv.ParseUint(hex[4:6], 16, 8)

	return color.RGBA{R: uint8(r), G: uint8(g), B: uint8(b), A: 255}
}
