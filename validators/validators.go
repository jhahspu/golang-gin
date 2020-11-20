package validators

import (
	"strings"

	"github.com/go-playground/validator/v10"
)

// ValidateCoolTitle is custom validator for "is-cool"
// and will check if Title field contains "cool"
func ValidateCoolTitle(field validator.FieldLevel) bool {
	return strings.Contains(field.Field().String(), "Cool")
}
