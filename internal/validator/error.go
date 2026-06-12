package validator

import "github.com/go-playground/validator/v10"

func FormatValidationErrors(err error) map[string]string {

	errors := make(map[string]string)

	for _, e := range err.(validator.ValidationErrors) {

		switch e.Tag() {

		case "required":
			errors[e.Field()] = "field is required"

		case "datetime":
			errors[e.Field()] = "invalid date"

		case "min":
			errors[e.Field()] = "value too short"

		default:
			errors[e.Field()] = "invalid value"
		}
	}

	return errors
}
