package piscine

func isValid(board [][]rune, row, col int, num rune) bool { // return true if the num is not in the 3x3 square
	// or in the same row and column
	for c := 0; c < len(board[0]); c++ { // check if num is in the col till length of column
		if c != col && board[row][c] == num {
			return false
		}
	}
	for r := 0; r < len(board); r++ { // check if num is in the row till length of row
		if r != row && board[r][col] == num {
			return false
		}
	}
	startRow := (row / 3) * 3 // // get where it starts from the top left corner of the  specific 3x3 square
	startCol := (col / 3) * 3
	for i := startRow; i < startRow+3; i++ { // nested loop to check if the num is in 3x3 square, this loop for line
		for j := startCol; j < startCol+3; j++ {
			if i != row && j != col && board[i][j] == num {
				// check if num is in board starting startrow till startrow+3 then colmuns
				return false
			}
		}
	}
	return true // otherwise return true is valid
}
