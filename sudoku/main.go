package main

import "fmt"

// Vérifie si un nombre peut être placé à une position donnée sur le plateau
func isValid(board [9][9]int, row, col, num int) bool {
	// Vérifie la ligne
	for x := 0; x < 9; x++ {
		if board[row][x] == num {
			return false
		}
	}

	// Vérifie la colonne
	for x := 0; x < 9; x++ {
		if board[x][col] == num {
			return false
		}
	}

	// Vérifie la boîte 3x3
	startRow := row - row%3
	startCol := col - col%3
	for i := 0; i < 3; i++ {
		for j := 0; j < 3; j++ {
			if board[i+startRow][j+startCol] == num {
				return false
			}
		}
	}
	return true
}

// Résout le Sudoku en utilisant une méthode de backtracking
func solveSudoku(board *[9][9]int) bool {
	for i := 0; i < 9; i++ {
		for j := 0; j < 9; j++ {
			if (*board)[i][j] == 0 {
				for num := 1; num <= 9; num++ {
					if isValid(*board, i, j, num) {
						(*board)[i][j] = num
						if solveSudoku(board) {
							return true
						}
						(*board)[i][j] = 0
					}
				}
				return false
			}
		}
	}
	return true
}

// Affiche le plateau de Sudoku résolu
func printBoard(board [9][9]int) {
	for i := 0; i < 9; i++ {
		if i%3 == 0 && i != 0 {
			fmt.Println("- - - - - - - - - - - -")
		}
		for j := 0; j < 9; j++ {
			if j%3 == 0 && j != 0 {
				fmt.Print(" | ")
			}
			if j == 8 {
				fmt.Println(board[i][j])
			} else {
				fmt.Print(board[i][j], " ")
			}
		}
	}
}

func main() {
	board := [9][9]int{
		{5, 0, 4, 7, 0, 0, 0, 0, 0},
		{0, 0, 7, 0, 0, 3, 9, 0, 1},
		{0, 0, 0, 8, 0, 9, 6, 0, 0},
		{7, 0, 0, 0, 4, 0, 0, 6, 0},
		{1, 3, 5, 0, 0, 0, 0, 0, 4},
		{4, 0, 0, 0, 7, 0, 0, 1, 0},
		{0, 0, 0, 2, 0, 7, 4, 0, 0},
		{0, 0, 2, 0, 0, 4, 3, 0, 5},
		{3, 0, 8, 6, 0, 0, 0, 0, 0},
	}

	if solveSudoku(&board) {
		printBoard(board)
	} else {
		fmt.Println("Aucune solution n'existe")
	}
}
