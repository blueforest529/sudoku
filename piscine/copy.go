package piscine

func Copy(source [][]rune) [][]rune {
	var destination [][]rune
	var empty []rune
	for i := 0; i < len(source); i++ {
		destination = append(destination, empty)
		for j := 0; j < len(source[i]); j++ {
			destination[i] = append(destination[i], source[i][j])
		}
	}
	return destination
}
