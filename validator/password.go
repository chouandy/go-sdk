package validator

import "unicode"

// CheckPassword check password
func CheckPassword(password string, n int, number, lower, upper, special bool) bool {
	// Check length
	if len(password) < n {
		return false
	}

	// Check Number, Lowercase character, Uppercase character, Special character
	var isNumber, isLower, isUpper, isSpecial bool

	// Ignore number check
	if !number {
		isNumber = true
	}
	// Ignore lower check
	if !lower {
		isLower = true
	}
	// Ignore upper check
	if !upper {
		isUpper = true
	}
	// Ignore special check
	if !special {
		isSpecial = true
	}

	// Check password
	for _, c := range password {
		switch {
		case unicode.IsNumber(c):
			isNumber = true
		case unicode.IsLower(c):
			isLower = true
		case unicode.IsUpper(c):
			isUpper = true
		case unicode.IsPunct(c) || unicode.IsSymbol(c):
			isSpecial = true
		}
	}

	return isNumber && isLower && isUpper && isSpecial
}
