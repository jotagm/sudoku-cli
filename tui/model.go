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
	fixed     [9][9]bool
}

func NewModel() Model {
	board := core.Generate()

	board.RemoveCells(30)

	var fixed [9][9]bool
	for i := 0; i < 9; i++ {
		for j := 0; j < 9; j++ {
			fixed[i][j] = board.Matriz[i][j] != 0

		}
	}

	return Model{
		board: board,
		fixed: fixed,
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

		case "1", "2", "3", "4", "5", "6", "7", "8", "9":
			num, _ := strconv.Atoi(msg.String())
			if !m.fixed[m.cursorRow][m.cursorCol] {
				m.board.Matriz[m.cursorRow][m.cursorCol] = num
			}
		}

	}
	return m, nil
}
func (m Model) View() string {
	s := ""
	for i := 0; i < m.board.Altura; i++ {
		if i%3 == 0 && i != 0 {
			s += "------+-------+------\n"
		}
		for j := 0; j < m.board.Largura; j++ {
			if j%3 == 0 && j != 0 {
				s += "| "
			}
			if i == m.cursorRow && j == m.cursorCol {
				s += "[" + TrocaZeroPorPonto(m.board.Matriz[i][j]) + "] "
			} else {
				s += TrocaZeroPorPonto(m.board.Matriz[i][j]) + " "
			}
		}
		s += "\n"
	}
	s += "cursor: " + strconv.Itoa(m.cursorRow) + ", " + strconv.Itoa(m.cursorCol) + "\n"
	return s
}
func TrocaZeroPorPonto(val int) string {
	if val == 0 {
		return "."
	}
	return strconv.Itoa(val)
}
