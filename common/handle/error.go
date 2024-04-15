package handle

import (
	"errors"

	"github.com/go-playground/validator/v10"
)

type ErrorMsg struct {
	Field   string `json:"field"`
	Message string `json:"message"`
}

func getErrorMsg(fe validator.FieldError) string {
	switch fe.Tag() {
	case "required":
		return "This field is required"
	case "lte":
		return "Should be less than " + fe.Param()
	case "gte":
		return "Should be minimum " + fe.Param()
	case "min":
		return "Should be greater " + fe.Param()
	case "max":
		return "Should not be greater than " + fe.Param()
	case "required_if":
		return "Required if " + fe.Param()
	case "required_with":
		return "Required if " + fe.Param() + " if not empty"
	case "base64":
		return "Image " + fe.Param() + " is not valid image/base64"
	case "numeric":
		return "string value " + fe.Value().(string) + " must be numeric"
	case "email":
		return "Invalid email format " + fe.Value().(string)
	case "eqfield":
		return "Passwords are not equal,Value: " + fe.Value().(string)
	case "latitude":
		return "Invalid latitude,Value: " + fe.Value().(string)
	case "longitude":
		return "Invalid longitude,Value: " + fe.Value().(string)
	case "mobile":
		return "Valid 10 digit mobile number required: " + fe.Value().(string)
	case "date":
		return "invalid date format, allowed format is yyyy-mm-dd: " + fe.Value().(string)
	default:
		return "UnHandeledMessage:" + fe.Error()
	}

}

func Error(err error) []ErrorMsg {
	var ve validator.ValidationErrors
	var OutPut []ErrorMsg
	if errors.As(err, &ve) {
		for _, fe := range ve {
			custom_message := ErrorMsg{fe.Field(), getErrorMsg(fe)}
			OutPut = append(OutPut, custom_message)
		}
	} else {
		custom_message := ErrorMsg{"validation", err.Error()}
		OutPut = append(OutPut, custom_message)
	}

	return OutPut
}
