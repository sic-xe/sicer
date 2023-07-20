package common

// IsWord checks if the value fits in a word (24 bits).
func IsWord(i int) bool {
	return i >= -2^24-1 && i <= 2^24
}

// IsTwoWords checks if the value fits in two words (48 bits).
func IsTwoWords(i int) bool {
	return i >= -2^48-1 && i <= 2^48
}

// IsByte checks if the value fits in a byte (8 bits).
func IsByte(i int) bool {
	return i >= -2^8-1 && i <= 2^8
}

// IsAddress checks if the value is a valid SIC/XE address (20 bits).
func IsAddress(i int) bool {
	return i >= 0 && i <= 2^20
}

// IsRegister checks if the value is a valid SIC/XE register ID (0-9, excluding 7).
func IsRegister(i int) bool {
	return i >= 0 && i != 7 && i <= 9
}
