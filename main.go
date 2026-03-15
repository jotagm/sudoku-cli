package main

import (
	"fmt"
	"math/rand/v2"
)

type Sudoku struct {
	Largura int
	Altura  int
	Matriz  [][]int
}

func (s Sudoku) IsValid(row, col, num int) bool {
	// linha
	for j := 0; j < s.Largura; j++ {
		if s.Matriz[row][j] == num {
			return false
		}
	}
	// coluna
	for i := 0; i < s.Altura; i++ {
		if s.Matriz[i][col] == num {
			return false
		}
	}
	// bloco 3x3
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
func (s Sudoku) gerarNumero() int {
	num := rand.IntN(9) + 1

	return num
}

func (s *Sudoku) fill(row, col int) bool {
	/*
					fill(row, col):
		    caso base: se row == 9, acabou, retorna true

		    calcula próxima posição:
		        se col == 8 → próxima é (row+1, 0)
		        senão       → próxima é (row, col+1)

		    se célula já tem número:
		        só avança → fill(nextRow, nextCol)

		    senão:
		        para cada num em 1-9 embaralhado:
		            se IsValid(row, col, num):
		                coloca num
		                se fill(nextRow, nextCol) retornar true:
		                    retorna true
		                remove num  ← backtrack
		        retorna false
	*/
	for i := 0; i < s.Altura; i++ {
		for j := 0; j < s.Largura; j++ {
			if s.Matriz[i][j] == 0 {
				for num := 1; num <= 9; num++ {
					if s.IsValid(i, j, num) {
						s.Matriz[i][j] = num
						if s.fill(i, j) {
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

func puzzle() {
	fmt.Println("iniciando matriz")

	sudoku := Sudoku{
		Largura: 9,
		Altura:  9,
		Matriz: [][]int{
			{5, 3, 0, 0, 7, 0, 0, 0, 0},
			{6, 0, 0, 1, 9, 5, 0, 0, 0},
			{0, 9, 8, 0, 0, 0, 0, 6, 0},
			{8, 0, 0, 0, 6, 0, 0, 0, 3},
			{4, 0, 0, 8, 0, 3, 0, 0, 1},
			{7, 0, 0, 0, 2, 0, 0, 0, 6},
			{0, 6, 0, 0, 0, 0, 2, 8, 0},
			{0, 0, 0, 4, 1, 9, 0, 0, 5},
			{0, 0, 0, 0, 8, 0, 0, 7, 9},
		},
	}

	sudoku.Print()
}

func main() {
	puzzle()
}
