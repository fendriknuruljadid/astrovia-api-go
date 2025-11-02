package utils

import (
    "github.com/go-playground/validator/v10"
    "strings"
)

var Validate = validator.New()

func ValidationErrors(err error) map[string]string {
    errors := make(map[string]string)

    for _, e := range err.(validator.ValidationErrors) {
        field := strings.ToLower(e.Field())

        switch e.Tag() {
        case "required":
            errors[field] = "this field is required"
        case "email":
            errors[field] = "invalid email format"
        case "min":
            errors[field] = "must be at least " + e.Param() + " characters"
        case "max":
            errors[field] = "cannot exceed " + e.Param() + " characters"
        default:
            errors[field] = "invalid value"
        }
    }

    return errors
}
