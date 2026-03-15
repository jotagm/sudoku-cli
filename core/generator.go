package core

import "math/rand/v2"

func (s *Sudoku) fill() bool {
	for i := 0; i < s.Altura; i++ {
		for j := 0; j < s.Largura; j++ {
			if s.Matriz[i][j] == 0 {
				nums := rand.Perm(9)
				for _, n := range nums {
					num := n + 1
					if s.IsValid(i, j, num) {
						s.Matriz[i][j] = num
						if s.fill() {
							return true
						}
						s.Matriz[i][j] = 0
					}
				}
				return false
			}
		}
	}
	return true
}

func Generate() *Sudoku {
	s := &Sudoku{
		Largura: 9,
		Altura:  9,
		Matriz:  make([][]int, 9),
	}
	for i := range s.Matriz {
		s.Matriz[i] = make([]int, 9)
	}
	s.fill()
	return s
}

func (s *Sudoku) RemoveCells(clues int) {

	posicao := rand.Perm(s.Largura * s.Altura)
	removido := 0
	target := s.Largura*s.Altura - clues
	for _, p := range posicao {
		if removido == target {
			break
		}
		row := p / s.Largura
		col := p % s.Largura
		if s.Matriz[row][col] != 0 {
			s.Matriz[row][col] = 0
			removido++
		}
	}
}
