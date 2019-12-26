package db

import "regexp"

// IsDuplicateError returns true if error contains a Duplicate entry '' for key '' error
func IsDuplicateError(err error) bool {
	var isDuplicate = regexp.MustCompile(`Duplicate entry '\w*' for key`)
	return isDuplicate.MatchString(err.Error())
}
