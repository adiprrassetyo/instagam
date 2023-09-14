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
		// if username already exist
		if error.IsSame(err, error.ErrUsernameAlreadyExist) {
			errorMessage := api.SetError("Username already exist")
			resp := api.APIResponse("Register Account Failed", http.StatusBadRequest, "error", errorMessage)
			c.JSON(http.StatusBadRequest, resp)
			return
			// if email already exist
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

// login
func (ctrl *UserController) Login(c *gin.Context) {
	var input domain.LoginUserInput
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
			resp := api.APIResponse("Login Failed", http.StatusBadRequest, "error", result)
			c.JSON(http.StatusBadRequest, resp)
			return
		}
		errorMessage := api.SetError(err.Error())
		resp := api.APIResponse("Login Failed", http.StatusBadRequest, "error", errorMessage)
		c.JSON(http.StatusBadRequest, resp)
		return
	}

	// check user in database
	user, err := ctrl.UserUseCase.LoginUser(input)
	if err != nil {
		log.Println(err)
		// if email not found
		if error.IsSame(err, error.ErrEmailNotFound) || error.IsSame(err, error.ErrDataLoginNotFound) {
			errorMessage := api.SetError("Email or Password is wrong")
			resp := api.APIResponse("Login Failed", http.StatusBadRequest, "error", errorMessage)
			c.JSON(http.StatusBadRequest, resp)
			return
		}
		// if username not found
		if error.IsSame(err, error.ErrUsernameNotFound) {
			errorMessage := api.SetError("Username or Password is wrong")
			resp := api.APIResponse("Login Failed", http.StatusBadRequest, "error", errorMessage)
			c.JSON(http.StatusBadRequest, resp)
			return
		}
		resp := api.APIResponse("Login Failed", http.StatusInternalServerError, "error", nil)
		c.JSON(http.StatusInternalServerError, resp)
		return
	}
	// Generate JWT Token
	token, err := jwt.GenerateToken(user.ID)
	if err != nil {
		log.Println(err)
		resp := api.APIResponse("Login Failed", http.StatusInternalServerError, "error", nil)
		c.JSON(http.StatusInternalServerError, resp)
		return
	}
	userResponse := api.SetUserResponse(user, token)
	resp := api.APIResponse("Login Success", http.StatusOK, "success", userResponse)
	c.JSON(http.StatusOK, resp)
}

// get all social media
func (ctrl *UserController) GetAllSocialMedia(c *gin.Context) {
	media, err := ctrl.UserUseCase.AllSocialMedia()
	if err != nil {
		log.Println(err)
		resp := api.APIResponse("Get Social Media Failed", http.StatusInternalServerError, "error", nil)
		c.JSON(http.StatusInternalServerError, resp)
		return
	}
	// if media not found
	if len(media) == 0 {
		resp := api.APIResponse("Social Media Not Found", http.StatusOK, "success", nil)
		c.JSON(http.StatusOK, resp)
		return
	}
	resp := api.APIResponse("Get Social Media Success", http.StatusOK, "success", media)
	c.JSON(http.StatusOK, resp)
}

// get by id social media
func (ctrl *UserController) GetOneSocialMedia(c *gin.Context) {
	// get id from url
	id := c.Param("id")
	media, err := ctrl.UserUseCase.OneSocialMedia(id)
	// if media not found
	if err != nil {
		log.Println(err)
		if error.IsSame(err, error.ErrDataNotFound) {
			errMessage := api.SetError("Social Media Not Found!")
			resp := api.APIResponse("Get Social Media Failed", http.StatusNotFound, "error", errMessage)
			c.JSON(http.StatusNotFound, resp)
			return
		}
		resp := api.APIResponse("Get Social Media Failed", http.StatusInternalServerError, "error", nil)
		c.JSON(http.StatusInternalServerError, resp)
		return
	}
	resp := api.APIResponse("Get Social Media Success", http.StatusOK, "success", media)
	c.JSON(http.StatusOK, resp)
}

// create social media
func (ctrl *UserController) CreateSocialMedia(c *gin.Context) {
	// check input users and validation
	var input domain.InsertSocialMedia
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
			resp := api.APIResponse("Create Social Media Failed", http.StatusBadRequest, "error", result)
			c.JSON(http.StatusBadRequest, resp)
			return
		}
		errorMessage := api.SetError(err.Error())
		resp := api.APIResponse("Create Social Media Failed", http.StatusBadRequest, "error", errorMessage)
		c.JSON(http.StatusBadRequest, resp)
		return
	}
	// check if user already have social media
	currentUser := c.MustGet("currentUser").(domain.User)
	err := ctrl.UserUseCase.CheckSocialMedia(currentUser.ID)
	if err != nil {
		if error.IsSame(err, error.ErrSocialMediaAlreadyExist) {
			errorMessage := api.SetError("User Already Have Social Media")
			resp := api.APIResponse("Register Account Failed", http.StatusBadRequest, "error", errorMessage)
			c.JSON(http.StatusBadRequest, resp)
			return
		}
	}
	// insert data to database
	socialmedia, err := ctrl.UserUseCase.CreateSocialMedia(input, currentUser.ID)
	if err != nil {
		log.Println(err)
		resp := api.APIResponse("Create Social Media Failed", http.StatusInternalServerError, "error", nil)
		c.JSON(http.StatusInternalServerError, resp)
		return
	}
	resp := api.APIResponse("Create Social Media Success", http.StatusOK, "success", socialmedia)
	c.JSON(http.StatusOK, resp)
}

// update social media
func (ctrl *UserController) UpdateSocialMedia(c *gin.Context) {
	// get id from url
	id := c.Param("id")
	var input domain.UpdateSocialMedia
	// verification input
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
			resp := api.APIResponse("Update Social Media Failed", http.StatusBadRequest, "error", result)
			c.JSON(http.StatusBadRequest, resp)
			return
		}
		errorMessage := api.SetError(err.Error())
		resp := api.APIResponse("Update Social Media Failed", http.StatusBadRequest, "error", errorMessage)
		c.JSON(http.StatusBadRequest, resp)
		return
	}
	// check if input name and social media url empty
	if input.Name == "" && input.Social_media_url == "" {
		errorMessage := api.SetError("Name and Social Media Url Cannot Be Empty")
		resp := api.APIResponse("Update Social Media Failed", http.StatusBadRequest, "error", errorMessage)
		c.JSON(http.StatusBadRequest, resp)
		return
	}
	// update data to database
	currentUser := c.MustGet("currentUser").(domain.User)
	socialmedia, err := ctrl.UserUseCase.UpdateSocialMedia(input, id, currentUser.ID)
	// if media not found
	if err != nil {
		log.Println(err)
		if error.IsSame(err, error.ErrSocialMediaNotFound) {
			errorMessage := api.SetError("Social Media Not Found!")
			resp := api.APIResponse("Update Social Media Failed", http.StatusNotFound, "error", errorMessage)
			c.JSON(http.StatusNotFound, resp)
			return
		}
		resp := api.APIResponse("Update Social Media Failed", http.StatusInternalServerError, "error", nil)
		c.JSON(http.StatusInternalServerError, resp)
		return
	}
	// response dataS
	resp := api.APIResponse("Update Social Media Success", http.StatusOK, "success", socialmedia)
	c.JSON(http.StatusOK, resp)
}

// delete social media
func (ctrl *UserController) DeleteSocialMedia(c *gin.Context) {
	// get id from url
	id := c.Param("id")
	currentUser := c.MustGet("currentUser").(domain.User)
	err := ctrl.UserUseCase.DeleteSocialMedia(id, currentUser.ID)
	if err != nil {
		log.Println(err)
		if error.IsSame(err, error.ErrSocialMediaNotFound) {
			errorMessage := api.SetError("Social Media Not Found!")
			resp := api.APIResponse("Delete Social Media Failed", http.StatusNotFound, "error", errorMessage)
			c.JSON(http.StatusNotFound, resp)
			return
		}
		resp := api.APIResponse("Delete Social Media Failed", http.StatusInternalServerError, "error", nil)
		c.JSON(http.StatusInternalServerError, resp)
		return
	}
	resp := api.APIResponse("Delete Social Media Success", http.StatusOK, "success", nil)
	c.JSON(http.StatusOK, resp)
}
