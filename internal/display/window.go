package display

import (
	"muse/internal/config"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/widget"
)

type Display struct {
	app        fyne.App
	window     fyne.Window
	cfg        *config.Config
	label      *widget.Label
	updateChan chan string
}

func NewDisplay(c *config.Config) *Display {
	var d Display
	d.app = app.New()
	d.window = d.app.NewWindow(c.General.ProgramName)
	d.cfg = c

	return &d
}

func (d *Display) Start() {

	// Set window size
	d.window.SetFixedSize(true)
	d.window.Resize(fyne.NewSize(
		float32(d.cfg.Display.WindowW),
		float32(d.cfg.Display.WindowH),
	))

	// Set up text display
	d.label = widget.NewLabel("")
	d.label.Alignment = fyne.TextAlignCenter
	d.label.Wrapping = fyne.TextWrapWord
	d.label.TextStyle = fyne.TextStyle{Bold: true}

	d.window.SetContent(d.label)

	d.window.ShowAndRun()
}

func (w *Display) UpdateText(text string) {
	if w.label != nil {
		fyne.DoAndWait(func() {
			w.label.SetText(text)
			w.label.Refresh()
		})
	}
}
