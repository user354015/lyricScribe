package main

import (
	"strings"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
)

type Model struct {
	width       int
	height      int
	currentText string
}

type textUpdateMsg string

func New() Model {
	return Model{
		currentText: "",
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
	case tea.KeyMsg:
		if msg.Type == tea.KeyCtrlC {
			State = 0
			return m, tea.Quit
		}
	case textUpdateMsg:
		m.currentText = strings.TrimSpace(string(msg))
	}
	return m, nil
}

func (m Model) View() string {
	if m.width == 0 || strings.TrimSpace(m.currentText) == "" {
		return ""
	}

	textLines := strings.Count(m.currentText, "\n") + 1

	verticalPadding := (m.height - textLines) / 2
	if verticalPadding < 0 {
		verticalPadding = 0
	}

	style := lipgloss.NewStyle().
		Width(m.width).
		Align(lipgloss.Center).
		Foreground(lipgloss.Color("#ffffff")).
		Background(lipgloss.Color("#000000")).
		Padding(0, 1)

	content := strings.Repeat("\n", verticalPadding) +
		style.Render(m.currentText) +
		strings.Repeat("\n", m.height-verticalPadding-textLines)

	return content
}

type Program struct {
	program *tea.Program
}

func NewProgram() *Program {
	p := &Program{
		program: tea.NewProgram(New(), tea.WithAltScreen()),
	}
	go p.program.Run()
	return p
}

func (p *Program) UpdateDisplay(text string) {
	p.program.Send(textUpdateMsg(text))
}

func (p *Program) Quit() {
	p.program.Quit()
}
