package validator

import (
	"github.com/go-playground/validator/v10"

	dataex "github.com/chouandy/go-sdk/data"
)

// CountryCodeValidation country code validation
func CountryCodeValidation(fl validator.FieldLevel) bool {
	// Get countryCode string
	countryCode := fl.Field().String()

	if len(countryCode) != 2 {
		return false
	}

	_, exists := dataex.Countries[countryCode]
	return exists
}

func init() {
	Validator.RegisterValidation("country_code", CountryCodeValidation)
}
