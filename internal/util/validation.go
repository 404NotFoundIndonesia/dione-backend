package util

import (
	"fmt"
	"github.com/go-playground/validator/v10"
)

func Validate[T any](data T) map[string]string {
	err := validator.New().Struct(data)
	if err == nil {
		return nil
	}

	res := make(map[string]string)
	for _, msg := range err.(validator.ValidationErrors) {
		res[msg.Field()] = TranslateValidationTag(msg)
	}

	return res
}

func TranslateValidationTag(fd validator.FieldError) string {
	switch fd.ActualTag() {
	case "required":
		return fmt.Sprintf("%s is required", fd.Field())
	case "email":
		return fmt.Sprintf("%s must be a valid email address", fd.Field())
	case "min":
		return fmt.Sprintf("%s must be at least %s characters", fd.Field(), fd.Param())
	case "max":
		return fmt.Sprintf("%s must be at most %s characters", fd.Field(), fd.Param())
	case "eq":
		return fmt.Sprintf("%s must be equal to %s", fd.Field(), fd.Param())
	case "ne":
		return fmt.Sprintf("%s must not be equal to %s", fd.Field(), fd.Param())
	case "lt":
		return fmt.Sprintf("%s must be less than %s", fd.Field(), fd.Param())
	case "lte":
		return fmt.Sprintf("%s must be less than or equal to %s", fd.Field(), fd.Param())
	case "gt":
		return fmt.Sprintf("%s must be greater than %s", fd.Field(), fd.Param())
	case "gte":
		return fmt.Sprintf("%s must be greater than or equal to %s", fd.Field(), fd.Param())
	case "len":
		return fmt.Sprintf("%s must be exactly %s characters", fd.Field(), fd.Param())
	case "oneof":
		return fmt.Sprintf("%s must be one of [%s]", fd.Field(), fd.Param())
	case "url":
		return fmt.Sprintf("%s must be a valid URL", fd.Field())
	case "uuid":
		return fmt.Sprintf("%s must be a valid UUID", fd.Field())
	case "alphanum":
		return fmt.Sprintf("%s must contain only alphanumeric characters", fd.Field())
	case "alpha":
		return fmt.Sprintf("%s must contain only alphabetic characters", fd.Field())
	case "numeric":
		return fmt.Sprintf("%s must be a valid number", fd.Field())
	case "boolean":
		return fmt.Sprintf("%s must be true or false", fd.Field())
	case "datetime":
		return fmt.Sprintf("%s must be a valid datetime format", fd.Field())
	case "contains":
		return fmt.Sprintf("%s must contain the text '%s'", fd.Field(), fd.Param())
	case "excludes":
		return fmt.Sprintf("%s must not contain the text '%s'", fd.Field(), fd.Param())
	case "startswith":
		return fmt.Sprintf("%s must start with '%s'", fd.Field(), fd.Param())
	case "endswith":
		return fmt.Sprintf("%s must end with '%s'", fd.Field(), fd.Param())
	default:
		return fmt.Sprintf("%s is invalid (%s)", fd.Field(), fd.Tag())
	}
}
