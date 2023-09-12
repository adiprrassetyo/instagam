package controllers

import (
	"errors"
	"instagam/modules/v1/users/domain"
	api "instagam/pkg/api_response"
	error "instagam/pkg/http-error"
	jwt "instagam/pkg/token"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

// register
func (ctrl *UserController) Register(c *gin.Context) {
	// check input users and validation
	var input domain.RegisterUserInput
	if err := c.ShouldBindJSON(&input); err != nil {
		log.Println(err)
		var verification validator.ValidationErrors
		if errors.As(err, &verification) {
			result := make([]error.Form, len(verification))
			for i, val := range verification {
				result[i] = error.Form{
					Field:   val.Field(),
					Message: error.FormValidationError(val),
				}
			}
			resp := api.APIResponse("Register Account Failed", http.StatusBadRequest, "error", result)
			c.JSON(http.StatusBadRequest, resp)
			return
		}
		errorMessage := api.SetError(err.Error())
		resp := api.APIResponse("Register Account Failed", http.StatusBadRequest, "error", errorMessage)
		c.JSON(http.StatusBadRequest, resp)
		return
	}
	// insert data to database
	newUser, err := ctrl.UserUseCase.RegisterUser(input)
	if err != nil {
		log.Println(err)
		if error.IsSame(err, error.ErrUsernameAlreadyExist) {
			errorMessage := api.SetError("Username already exist")
			resp := api.APIResponse("Register Account Failed", http.StatusBadRequest, "error", errorMessage)
			c.JSON(http.StatusBadRequest, resp)
			return
		} else if error.IsSame(err, error.ErrEmailAlreadyExist) {
			errorMessage := api.SetError("Email already exist")
			resp := api.APIResponse("Register Account Failed", http.StatusBadRequest, "error", errorMessage)
			c.JSON(http.StatusBadRequest, resp)
			return
		}
		resp := api.APIResponse("Register Account Failed", http.StatusBadRequest, "error", nil)
		c.JSON(http.StatusBadRequest, resp)
		return
	}
	// Generate JWT Token
	token, err := jwt.GenerateToken(newUser.ID)
	if err != nil {
		log.Println(err)
		resp := api.APIResponse("Register Account Failed", http.StatusInternalServerError, "error", nil)
		c.JSON(http.StatusInternalServerError, resp)
		return
	}
	user := api.SetUserResponse(newUser, token)
	resp := api.APIResponse("Register Account Success", http.StatusOK, "success", user)
	c.JSON(http.StatusOK, resp)
}
