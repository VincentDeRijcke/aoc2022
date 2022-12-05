package utils

// IsNotDigit returns false if ascii digit
func IsNotDigit(r rune) bool {
	return !IsDigit(r)
}

// IsDigit returns true if ascii digit
func IsDigit(r rune) bool {
	return r >= '0' && r <= '9'
}
