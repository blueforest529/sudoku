package piscine

func NextSquare(x int, y int, grid [][]rune) (int, int) {
	for i := x; i < 9; i++ {
		for j := 0; j < 9; j++ {
			if grid[i][j] == '.' {
				return i, j
			}
		}
	}
	return -2, -2
}
