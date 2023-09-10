package errorsHandling

import (
	api "instagam/pkg/api_response"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

func IsSame(err error, target error) bool {
	return err.Error() == target.Error()
}

func NotFound() gin.HandlerFunc {
	return func(c *gin.Context) {
		resp := api.APIResponse("error", http.StatusNotFound, "Not Found", nil)
		c.JSON(http.StatusNotFound, resp)
	}
}

func NoMethod() gin.HandlerFunc {
	return func(c *gin.Context) {
		resp := api.APIResponse("error", http.StatusMethodNotAllowed, "Method Not Allowed", nil)
		c.JSON(http.StatusMethodNotAllowed, resp)
	}
}

func FormValidationError(fe validator.FieldError) string {
	switch fe.Tag() {
	case "required":
		return fe.Field() + " is required!"
	case "email":
		return fe.Field() + " must be a valid email address!"
	case "min":
		return fe.Field() + " minimum " + fe.Param()
	case "max":
		return fe.Field() + " maximum " + fe.Param()
	case "alphanum":
		return fe.Field() + " must be alphanumeric!"
	case "number":
		return fe.Field() + " must be a number!"
	case "numeric":
		return fe.Field() + " must be numeric!"
	case "eqfield":
		return fe.Field() + " must be equal to " + fe.Param() + "!"
	case "alphanumunicode":
		return fe.Field() + " must be alphanumeric unicode!"
	case "required_without":
		return fe.Field() + " is required because " + fe.Param() + " is empty!"
	case "url":
		return fe.Field() + " must be a valid URL with https or http!"
	default:
		return fe.Field() + " is invalid!"
	}
}