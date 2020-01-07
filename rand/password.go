package rand

import "github.com/chouandy/go-sdk/validator"

// PasswordCharacters password characters
var PasswordCharacters = "0123456789abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ~!@#$%^&*()_+`-={}|[]\\:\"<>?,./"

// Password generate randpassword
func Password(n int, number, upper, lower, special bool) string {
	password := String(PasswordCharacters, n)
	for !validator.CheckPassword(password, n, number, upper, lower, special) {
		password = String(PasswordCharacters, n)
	}
	return password
}
