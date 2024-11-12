package validation

import (
	"github.com/go-playground/validator/v10"
)

func ValidatePAN(fl validator.FieldLevel) bool {
	pan := fl.Field().String()
	if len(pan) != 10 {
		return false
	}
	for i, r := range pan {
		if i < 5 || i == 9 {
			if r < 'A' || r > 'Z' {
				return false
			}
		} else {
			if r < '0' || r > '9' {
				return false
			}
		}
	}
	return true
}
