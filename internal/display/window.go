package display

import (
	"bytes"
	"image/color"
	"muse/internal/config"
	"muse/internal/util"
	"strings"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/examples/resources/fonts"
	"github.com/hajimehoshi/ebiten/v2/text/v2"
)

type Display struct {
	text string

	posx   int
	posy   int
	width  int
	height int

	fgCol color.RGBA
	bgCol color.RGBA

	face      text.Face
	config    *config.Config
	formatter *TextFormatter

	shouldQuit bool
}

func (d *Display) Layout(outsideHeight, outsideWidth int) (screenW, screenH int) {
	return d.width, d.height
}
func (d *Display) Update() error {
	return nil
}

func (d *Display) Draw(screen *ebiten.Image) {
	screen.Fill(color.Transparent)
	// screen.Fill(d.bgCol)

	textOps := &text.DrawOptions{}
	textOps.GeoM.Translate(float64(d.width)/2, float64(d.height)/2)
	textOps.ColorScale.ScaleWithColor(d.fgCol)
	textOps.LineSpacing = float64(d.config.Display.FontSize) * 1.2
	textOps.PrimaryAlign = text.AlignCenter
	textOps.SecondaryAlign = text.AlignCenter
	text.Draw(screen, d.text, d.face, textOps)

}

// ---------
// --
// ---------

func SetUpGui(c *config.Config) *Display {
	var d Display

	d.config = c
	d.posx = d.config.Display.WindowX
	d.posy = d.config.Display.WindowY
	d.width = d.config.Display.WindowW
	d.height = d.config.Display.WindowH
	d.bgCol = util.HexToRGBA(d.config.Display.BgColor)
	d.fgCol = util.HexToRGBA(d.config.Display.FgColor)

	d.formatter = NewTextLyricFormatter(int(float64(d.width) * 0.8))

	ff, _ := text.NewGoTextFaceSource(bytes.NewReader(fonts.MPlus1pRegular_ttf))

	d.face = &text.GoTextFace{
		Source: ff,
		Size:   float64(d.config.Display.FontSize),
	}

	return &d
}

func RunGui(d *Display) error {
	ebiten.SetWindowSize(d.width, d.height)
	ebiten.SetWindowPosition(d.posx, d.posy)
	// ebiten.SetWindowTitle(d.config.General.ProgramName)
	ebiten.SetWindowTitle("app")

	opts := &ebiten.RunGameOptions{ScreenTransparent: true}
	if err := ebiten.RunGameWithOptions(d, opts); err != nil {
		return err
	}
	return nil
}

func (d *Display) UpdateText(lyric string) {
	s, _ := text.Measure(lyric, d.face, 0)
	size := int(s)

	if size <= d.formatter.MaxWidth {
		d.text = lyric
		return
	}

	words := strings.Fields(lyric)
	if len(words) == 0 {
		d.text = ""
		return
	}

	linesN := int(size/d.formatter.MaxWidth) + 1
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

	d.text = strings.Join(lines, "\n")
}
