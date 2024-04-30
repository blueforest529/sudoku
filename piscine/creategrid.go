package piscine

import "os"

//If the input was invalid, returns the empty [][]rune, and false. If the input was valid, returns the grid and true
func CreateGrid() ([][]rune, bool) {
	var grid [][]rune
	var empty []rune
	new := os.Args[1:]
	if len(new) != 9 {
		return [][]rune(nil), false
	}
	//The for loops are essentially grid[i][j] = os.Args[i][j], if os.Args was a 2D array
	for i := 0; i < 9; i++ {
		grid = append(grid, empty)
		newnew := []rune(new[i])
		if len(newnew) != 9 {
			return [][]rune(nil), false
		} else {
			for j := 0; j < 9; j++ {
				grid[i] = append(grid[i], newnew[j])
			}
		}
	}
	return grid, true
}
