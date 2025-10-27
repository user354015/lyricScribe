package display

import (
	"muse/internal/config"
	"muse/internal/shared"
	"strings"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
)

type Model struct {
	width       int
	height      int
	formatter   *TextFormatter
	currentText string

	fgColor string
	bgColor string
}

type TextUpdateMsg shared.Lyric

func NewTUI(cfg *config.Config) Model {
	return Model{
		formatter:   NewTextLyricFormatter(100),
		currentText: "",

		fgColor: cfg.Display.FgColor,
		bgColor: cfg.Display.BgColor,
	}
}

func (m Model) Init() tea.Cmd {
	return nil
}

func (m Model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.WindowSizeMsg:
		m.width = msg.Width
		m.height = msg.Height
		m.formatter.MaxWidth = msg.Width - 8
	case tea.KeyMsg:
		if msg.Type == tea.KeyCtrlC || msg.String() == "q" {
			return m, tea.Quit
		}
	case TextUpdateMsg:
		m.currentText = strings.TrimSpace(string(msg.Lyric))
	}
	return m, nil
}

func (m Model) View() string {
	if m.width == 0 || m.height == 0 {
		return ""
	}

	wrappedLines := m.formatter.WrapTextChar(m.currentText)
	if len(wrappedLines) == 1 {
		if wrappedLines[0] == "" {
			return ""
		}
	}

	contentHeight := len(wrappedLines)
	verticalPadding := (m.height - contentHeight) / 2

	if verticalPadding < 0 {
		verticalPadding = 0
	}

	style := lipgloss.NewStyle().
		Width(m.width).
		Align(lipgloss.Center).
		Foreground(lipgloss.Color(m.fgColor)).
		Background(lipgloss.Color(m.bgColor))

	var styledLines []string
	for _, line := range wrappedLines {
		styledLines = append(styledLines, style.Render(line))
	}

	content := strings.Join(styledLines, "\n")

	paddedContent := strings.Repeat("\n", verticalPadding) +
		content +
		strings.Repeat("\n", verticalPadding)

	return paddedContent
}
