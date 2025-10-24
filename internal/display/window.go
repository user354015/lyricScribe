package display

import (
	"fmt"
	"image/color"
	"muse/internal/config"
	"time"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/theme"
)

type WindowDisplay struct {
	app       fyne.App
	window    fyne.Window
	config    *config.Config
	content   *fyne.Container
	textLabel *canvas.Text
	updates   chan string
}

type customTheme struct{}

func (c customTheme) Color(name fyne.ThemeColorName, variant fyne.ThemeVariant) color.Color {
	switch name {
	case theme.ColorNameBackground:
		return color.RGBA{R: 30, G: 30, B: 30, A: 255}
	case theme.ColorNameForeground:
		return color.RGBA{R: 200, G: 200, B: 200, A: 255}
	default:
		return theme.DefaultTheme().Color(name, variant)
	}
}

func (c customTheme) Font(style fyne.TextStyle) fyne.Resource {
	return theme.DefaultTheme().Font(style)
}

func (c customTheme) Icon(name fyne.ThemeIconName) fyne.Resource {
	return theme.DefaultTheme().Icon(name)
}

func (c customTheme) Size(name fyne.ThemeSizeName) float32 {
	return theme.DefaultTheme().Size(name)
}

func NewWindow(cfg *config.Config) *WindowDisplay {
	a := app.New()
	a.Settings().SetTheme(&customTheme{})

	return &WindowDisplay{
		app:     a,
		config:  cfg,
		updates: make(chan string, 100),
	}
}

func (w *WindowDisplay) Start() {
	// <CHANGE> Wrap window creation in fyne.Do() to ensure it runs on the main thread
	fyne.Do(func() {
		w.window = w.app.NewWindow(w.config.General.ProgramName)

		w.textLabel = canvas.NewText("Waiting for updates...", color.White)
		w.textLabel.TextSize = float32(w.config.Display.FontSize)
		w.textLabel.Alignment = fyne.TextAlignCenter

		w.content = container.NewVBox(w.textLabel)
		w.window.SetContent(w.content)

		w.window.Resize(fyne.NewSize(
			float32(w.config.Display.WindowW),
			float32(w.config.Display.WindowH),
		))

	})

	// <CHANGE> Start updateLoop in a goroutine so it doesn't block ShowAndRun
	go w.updateLoop()

	w.window.ShowAndRun()
}

func (w *WindowDisplay) updateLoop() {
	for update := range w.updates {
		// <CHANGE> Wrap UI updates in fyne.Do() to ensure they run on the main thread
		updateText := update
		fyne.Do(func() {
			w.textLabel.Text = updateText
			w.textLabel.Refresh()
		})
	}
}

func (w *WindowDisplay) Send(message string) {
	select {
	case w.updates <- fmt.Sprintf("[%s] %s", time.Now().Format("15:04:05"), message):
	default:
		fmt.Println("Update channel full, dropping message:", message)
	}
}

func (w *WindowDisplay) Close() {
	close(w.updates)
	if w.window != nil {
		w.window.Close()
	}
}
