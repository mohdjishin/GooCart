package utils

import (
	"fmt"

	"github.com/go-playground/validator/v10"
)

func ValidateStruct[T any](s T) error {
	validate := validator.New()
	if err := validate.Struct(s); err != nil {
		// Validation failed
		if _, ok := err.(*validator.InvalidValidationError); ok {
			return err
		}

		// Extract the individual validation errors
		var errMsg string
		for _, err := range err.(validator.ValidationErrors) {
			field := err.Field()
			errMsg += fmt.Sprintf("Validation failed for field '%s'\n", field)
		}
		return fmt.Errorf("validation errors:\n%s", errMsg)
	}
	return nil
}
