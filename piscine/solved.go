package piscine

func Solved(grid [][]rune) bool {
	for i := 0; i < 9; i++ {
		for j := 0; j < 9; j++ {
			if grid[i][j] == '.' {
				return false
			}
		}
	}
	return true
}
