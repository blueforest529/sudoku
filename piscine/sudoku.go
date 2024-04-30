package piscine

func Sudoku(x int, y int, grid [][]rune, solutions *[][][]rune) {
	for i := 1; i <= 9; i++ {
		if isValid(grid, x, y, rune(i+48)) {
			grid[x][y] = rune(i + 48)
			if Solved(grid) { // checks if all squares are filled, this is equivalent to it being solved
				*solutions = append(*solutions, Copy(grid))
			} else { //this else block means it never tries to NextSquare when the square is already 64
				a, b := NextSquare(x, y, grid)
				Sudoku(a, b, grid, solutions)
			}
		}
	}
	grid[x][y] = '.' //triggers if the path has already been completely explored
}
