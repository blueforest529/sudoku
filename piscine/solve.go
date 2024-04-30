package piscine

func Solve(board [][]rune) bool {
	for i := range board {
		for j := range board[i] {
			if board[i][j] == '.' { // Check for an . cell
				for c := '1'; c <= '9'; c++ { // Try numbers 1 through 9.
					if isValid(board, i, j, c) { // Check isValid.
						board[i][j] = c   // Place number c in the board if it's valid.
						if Solve(board) { // Recursively call solve to fill the rest of the board.
							return true // If the board is successfully filled, return true.
						} else {
							board[i][j] = '.' // If it fails, reset the cell full stop and try another number.
						}
					}
				}
				return false // Return false if no valid number from 1 to 9 is found.
			}
		}
	}
	return true // Return true if all cells are successfully filled.
}
