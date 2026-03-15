package tui

import (
	"strconv"
	"sudoku/core"

	tea "github.com/charmbracelet/bubbletea"
)

func (m Model) Init() tea.Cmd {
	return nil
}

type Model struct {
	board     *core.Sudoku
	cursorRow int
	cursorCol int
}

func NewModel() Model {
	return Model{
		board: core.Generate(),
	}
}

func (m Model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.String() {
		case "up", "k":
			if m.cursorRow > 0 {
				m.cursorRow--
			}
		case "down", "j":
			if m.cursorRow < 8 {
				m.cursorRow++
			}
		case "left", "h":
			if m.cursorCol > 0 {
				m.cursorCol--
			}
		case "right", "l":
			if m.cursorCol < 8 {
				m.cursorCol++
			}
		case "q", "ctrl+c":
			return m, tea.Quit
		}
	}
	return m, nil
}
func (m Model) View() string {
	return "cursor está em: " + strconv.Itoa(m.cursorRow) + ", " + strconv.Itoa(m.cursorCol) + "\n"
}
