package display

import (
	"muse/internal/config"
	"muse/internal/util"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/layout"
	"fyne.io/fyne/v2/widget"
)

type Display struct {
	app        fyne.App
	window     fyne.Window
	theme      fyneTheme
	cfg        *config.Config
	richText   *widget.RichText
	background *canvas.Rectangle

	updateChan chan string
}

type fyneTheme struct {
	fontResource fyne.Resource
	fontSize     float32
}

func NewDisplay(c *config.Config) *Display {
	var d Display

	d.app = app.New()
	d.window = d.app.NewWindow(c.General.ProgramName)
	d.cfg = c

	// Load and apply custom theme
	d.theme = createThemeResource(c)
	fontResource := loadFontResource(c.Display.Font)
	if fontResource != nil {
		d.app.Settings().SetTheme(&fyneTheme{
			fontResource: fontResource,
			fontSize:     float32(c.Display.FontSize),
		})
	}

	return &d
}

func (d *Display) Start() {

	// Set window size
	d.window.SetFixedSize(true)
	d.window.Resize(fyne.NewSize(
		float32(d.cfg.Display.WindowW),
		float32(d.cfg.Display.WindowH),
	))

	d.richText = widget.NewRichTextFromMarkdown("")
	d.richText.Wrapping = fyne.TextWrapWord
	for _, segment := range d.richText.Segments {
		if text, ok := segment.(*widget.TextSegment); ok {
			text.Style.Alignment = fyne.TextAlignCenter
		}
	}

	col := util.HexToRGBA(d.cfg.Display.BgColor)
	d.background = canvas.NewRectangle(col)
	d.background.Resize(fyne.NewSize(0, 0))

	textWithBg := container.NewStack(
		d.background,
		d.richText,
	)

	content := container.NewStack(
		container.NewVBox(
			layout.NewSpacer(),
			textWithBg,
			layout.NewSpacer(),
		))

	d.window.SetContent(content)
	d.window.ShowAndRun()
}

func (d *Display) UpdateText(text string) {
	if d.richText != nil {
		fyne.DoAndWait(func() {
			d.richText.ParseMarkdown(text)
			for _, segment := range d.richText.Segments {
				if textSeg, ok := segment.(*widget.TextSegment); ok {
					textSeg.Style.Alignment = fyne.TextAlignCenter
				}
			}

			d.richText.Refresh()
		})
	}
}
