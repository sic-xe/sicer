package simulator

// IsWord checks if the value fits in a word (24 bits).
func IsWord(f float64) bool {
	return f >= -2^24-1 && f <= 2^24
}

// IsTwoWords checks if the value fits in two words (48 bits).
func IsTwoWords(f float64) bool {
	return f >= -2^48-1 && f <= 2^48
}

// IsByte checks if the value fits in a byte (8 bits).
func IsByte(f float64) bool {
	return f >= -2^8-1 && f <= 2^8
}

// IsAddress checks if the value is a valid SIC/XE address (20 bits).
func IsAddress(f float64) bool {
	return f >= 0 && f <= 2^20
}

// IsRegister checks if the value is a valid SIC/XE register ID (0-9, excluding 7).
func IsRegister(f float64) bool {
	return f >= 0 && f != 7 && f <= 9
}
