package display

import (
	"image/color"
	"muse/internal/config"
	"muse/internal/shared"
	"os"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/theme"
)

/*
- - CUSTOM FYNE THEME - - start
*/
func (t *fyneTheme) Font(style fyne.TextStyle) fyne.Resource {
	if t.fontResource != nil {
		return t.fontResource
	}
	return theme.DefaultTheme().Font(style) // Fallback
}

func (t *fyneTheme) Size(name fyne.ThemeSizeName) float32 {
	if name == theme.SizeNameText {
		return t.fontSize
	}
	return theme.DefaultTheme().Size(name)
}

func (t *fyneTheme) Color(name fyne.ThemeColorName, variant fyne.ThemeVariant) color.Color {
	if name == theme.ColorNameBackground {
		return color.Transparent
	}
	return theme.DefaultTheme().Color(name, variant)
}

func (t *fyneTheme) Icon(name fyne.ThemeIconName) fyne.Resource {
	return theme.DefaultTheme().Icon(name)
}

func loadFontResource(fontPath string) fyne.Resource {
	fontData, err := os.ReadFile(fontPath)
	if err != nil {
		shared.Warn("Custom font could not be loaded: %v", err)
		return nil
	}
	return fyne.NewStaticResource("AppFont", fontData)
}

func createThemeResource(c *config.Config) fyneTheme {
	var theme fyneTheme
	fontData, err := os.ReadFile(c.Display.Font)
	if err != nil {
		shared.Warn("Custom font could not be loaded, falling back. Reason: ", err)
		return theme
	}

	theme.fontResource = fyne.NewStaticResource("AppFont", fontData)

	return theme
}

/*
- - CUSTOM FYNE THEME - - end
*/
