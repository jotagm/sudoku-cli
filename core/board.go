package core

import "fmt"

type Sudoku struct {
	Largura int
	Altura  int
	Matriz  [][]int
}

func (s Sudoku) IsValid(row, col, num int) bool {
	for j := 0; j < s.Largura; j++ {
		if s.Matriz[row][j] == num {
			return false
		}
	}
	for i := 0; i < s.Altura; i++ {
		if s.Matriz[i][col] == num {
			return false
		}
	}
	startRow := (row / 3) * 3
	startCol := (col / 3) * 3
	for i := startRow; i < startRow+3; i++ {
		for j := startCol; j < startCol+3; j++ {
			if s.Matriz[i][j] == num {
				return false
			}
		}
	}
	return true
}

func (s Sudoku) Print() {
	for i := 0; i < s.Altura; i++ {
		if i%3 == 0 && i != 0 {
			fmt.Println("---------------------")
		}
		for j := 0; j < s.Largura; j++ {
			if j%3 == 0 && j != 0 {
				fmt.Print("| ")
			}
			fmt.Printf("%d ", s.Matriz[i][j])
		}
		fmt.Println()
	}
}
