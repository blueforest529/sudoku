package piscine

import "fmt"

// Print answer board
func PrintBoard(board [][]rune) {
	for _, row := range board {
		for _, val := range row {
			if val != ' ' {
				//runt to char
				fmt.Printf("%c ", val)
			} else {
				// If it isn't solve print pull stop.
				fmt.Print(". ")
			}
		}
		fmt.Println()
	}
	fmt.Println()
}
