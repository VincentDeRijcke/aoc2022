package utils

// IsNotDigit returns false if ascii digit
func IsNotDigit(r rune) bool {
	return !IsDigit(r)
}

// IsDigit returns true if ascii digit
func IsDigit(r rune) bool {
	return r >= '0' && r <= '9'
}

func GridRunes(r rune, cols int, rows int) [][]rune {
	grid := make([][]rune, rows)
	for i := 0; i < rows; i++ {
		grid[i] = make([]rune, cols)
		for j := 0; j < cols; j++ {
			grid[i][j] = r
		}
	}

	return grid
}
