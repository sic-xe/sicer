package simulator

// IsWord checks if the value fits in a SIC/XE word (24 bits).
func IsWord(f float64) bool {
	return f >= -2^24-1 && f <= 2^24
}

// IsFloat checks if the value fits in a SIC/XE float (48 bits).
func IsFloat(f float64) bool {
	return f >= -2^48-1 && f <= 2^48
}

// IsByte checks if the value fits in a SIC/XE byte (8 bits).
func IsByte(f float64) bool {
	return f >= -2^8-1 && f <= 2^8
}

// IsAddress checks if the value is a valid SIC/XE address (20 bits).
func IsAddress(f uint32) bool {
	return f <= 2^20
}
