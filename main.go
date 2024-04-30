package main

import (
	"fmt"

	"sudoku/piscine"
)

func main() {
	// make a grid and check valid
	grid, valid := piscine.CreateGrid()
	if !valid {
		fmt.Println("Error")
		return
	}
	var solutions [][][]rune
	piscine.Sudoku(0, 0, grid, &solutions)
	if len(solutions) != 1 {
		fmt.Println("Error")
	} else {
		piscine.PrintBoard(solutions[0])
	}
}
