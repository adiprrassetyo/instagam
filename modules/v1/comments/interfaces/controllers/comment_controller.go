package controllers

import (
	"errors"
	"instagam/modules/v1/comments/domain"
	domainUser "instagam/modules/v1/users/domain"
	api "instagam/pkg/api_response"
	error "instagam/pkg/http-error"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

// get all comments
func (cc *CommentController) GetAllComments(c *gin.Context) {
	idPhotos := c.Param("id_photos")
	user := c.MustGet("currentUser").(domainUser.User)
	comments, err := cc.CommentUseCase.GetAllComments(idPhotos, user.ID)
	if err != nil {
		log.Println(err)
		if error.IsSame(err, error.ErrDataNotFound) {
			errMessage := api.SetError("Comment not found!")
			resp := api.APIResponse("Get All Comments Failed", http.StatusNotFound, "error", errMessage)
			c.JSON(http.StatusNotFound, resp)
			return
		}
		resp := api.APIResponse("Get All Comments Failed", http.StatusInternalServerError, "error", nil)
		c.JSON(http.StatusInternalServerError, resp)
		return
	}
	resp := api.APIResponse("Get All Comments Success", http.StatusOK, "success", comments)
	c.JSON(http.StatusOK, resp)
}

// get comment by id
func (cc *CommentController) GetCommentById(c *gin.Context) {
	id := c.Param("id")
	comment, err := cc.CommentUseCase.GetCommentById(id)
	if err != nil {
		log.Println(err)
		if error.IsSame(err, error.ErrDataNotFound) {
			errMessage := api.SetError("Comment Not Found!")
			resp := api.APIResponse("Get Comment By Id Failed", http.StatusNotFound, "error", errMessage)
			c.JSON(http.StatusNotFound, resp)
			return
		}
		resp := api.APIResponse("Get Comment By Id Failed", http.StatusInternalServerError, "error", nil)
		c.JSON(http.StatusInternalServerError, resp)
		return
	}
	resp := api.APIResponse("Get Comment By Id Success", http.StatusOK, "success", comment)
	c.JSON(http.StatusOK, resp)
}

// create comment
func (cc *CommentController) CreateComment(c *gin.Context) {
	var input domain.InsertComment
	if err := c.ShouldBindJSON(&input); err != nil {
		log.Println(err)
		var validation validator.ValidationErrors
		if errors.As(err, &validation) {
			result := make([]error.Form, len(validation))
			for i, v := range validation {
				result[i] = error.Form{
					Field:   v.Field(),
					Message: error.FormValidationError(v),
				}
			}
			resp := api.APIResponse("Create Comment Failed", http.StatusBadRequest, "error", result)
			c.JSON(http.StatusBadRequest, resp)
			return
		}
		errorMessage := api.SetError(err.Error())
		resp := api.APIResponse("Create Comment Failed", http.StatusBadRequest, "error", errorMessage)
		c.JSON(http.StatusBadRequest, resp)
		return
	}

	user := c.MustGet("currentUser").(domainUser.User)
	input.UserID = user.ID
	comment, err := cc.CommentUseCase.CreateComment(input)
	if err != nil {
		log.Println(err)
		if error.IsSame(err, error.ErrPhotoNotFound) {
			errMessage := api.SetError("Cannot Create Comment, Photo Not Found!")
			resp := api.APIResponse("Create Comment Failed", http.StatusNotFound, "error", errMessage)
			c.JSON(http.StatusNotFound, resp)
			return
		}
		resp := api.APIResponse("Create Comment Failed", http.StatusInternalServerError, "error", nil)
		c.JSON(http.StatusInternalServerError, resp)
		return
	}
	resp := api.APIResponse("Create Comment Success", http.StatusOK, "success", comment)
	c.JSON(http.StatusOK, resp)
}

// update comment
func (cc *CommentController) UpdateComment(c *gin.Context) {
	// get id from param
	id := c.Param("id")
	var input domain.UpdateComment
	if err := c.ShouldBindJSON(&input); err != nil {
		log.Println(err)
		var validation validator.ValidationErrors
		if errors.As(err, &validation) {
			result := make([]error.Form, len(validation))
			for i, v := range validation {
				result[i] = error.Form{
					Field:   v.Field(),
					Message: error.FormValidationError(v),
				}
			}
			resp := api.APIResponse("Update Comment Failed", http.StatusBadRequest, "error", result)
			c.JSON(http.StatusBadRequest, resp)
			return
		}
		errorMessage := api.SetError(err.Error())
		resp := api.APIResponse("Update Comment Failed", http.StatusBadRequest, "error", errorMessage)
		c.JSON(http.StatusBadRequest, resp)
		return
	}

	if input.PhotoID == 0 && input.Message == "" {
		errMessage := api.SetError("Please Fill Photo ID or Comment!")
		resp := api.APIResponse("Update Comment Failed", http.StatusBadRequest, "error", errMessage)
		c.JSON(http.StatusBadRequest, resp)
		return
	}

	user := c.MustGet("currentUser").(domainUser.User)
	comment, err := cc.CommentUseCase.UpdateComment(id, input, user.ID)
	if err != nil {
		log.Println(err)
		// if error is photo not found
		if error.IsSame(err, error.ErrPhotoNotFound) {
			errMessage := api.SetError("Cannot Update Comment, Photo Not Found!")
			resp := api.APIResponse("Update Comment Failed", http.StatusNotFound, "error", errMessage)
			c.JSON(http.StatusNotFound, resp)
			return
		}
		// if error is comment not found
		if error.IsSame(err, error.ErrCommentNotFound) || error.IsSame(err, error.ErrDataNotFound) {
			errMessage := api.SetError("Comment Not Found!")
			resp := api.APIResponse("Update Comment Failed", http.StatusNotFound, "error", errMessage)
			c.JSON(http.StatusNotFound, resp)
			return
		}
		resp := api.APIResponse("Update Comment Failed", http.StatusInternalServerError, "error", nil)
		c.JSON(http.StatusInternalServerError, resp)
		return
	}
	resp := api.APIResponse("Update Comment Success", http.StatusOK, "success", comment)
	c.JSON(http.StatusOK, resp)
}

// delete comment
func (cc *CommentController) DeleteComment(c *gin.Context) {
	// get id from param
	id := c.Param("id")
	user := c.MustGet("currentUser").(domainUser.User)
	err := cc.CommentUseCase.DeleteComment(id, user.ID)
	if err != nil {
		log.Println(err)
		// if error is comment not found
		if error.IsSame(err, error.ErrCommentNotFound) {
			errMessage := api.SetError("Comment Not Found!")
			resp := api.APIResponse("Delete Comment Failed", http.StatusNotFound, "error", errMessage)
			c.JSON(http.StatusNotFound, resp)
			return
		}
		resp := api.APIResponse("Delete Comment Failed", http.StatusInternalServerError, "error", nil)
		c.JSON(http.StatusInternalServerError, resp)
		return
	}
	resp := api.APIResponse("Delete Comment Success", http.StatusOK, "success", nil)
	c.JSON(http.StatusOK, resp)
}

// get all photos
func (cc *CommentController) GetAllPhotos(c *gin.Context) {
	photos, err := cc.CommentUseCase.GetAllPhotos()
	if err != nil {
		log.Println(err)
		if error.IsSame(err, error.ErrPhotoNotFound) {
			errMessage := api.SetError("Photo Not Found!")
			resp := api.APIResponse("Get All Photos Failed", http.StatusNotFound, "error", errMessage)
			c.JSON(http.StatusNotFound, resp)
			return
		}
		resp := api.APIResponse("Get All Photos Failed", http.StatusInternalServerError, "error", nil)
		c.JSON(http.StatusInternalServerError, resp)
		return
	}
	resp := api.APIResponse("Get All Photos Success", http.StatusOK, "success", photos)
	c.JSON(http.StatusOK, resp)
}

// get photo by id
func (cc *CommentController) GetPhotoById(c *gin.Context) {
	// get id from param
	id := c.Param("id")
	user := c.MustGet("currentUser").(domainUser.User)
	photo, err := cc.CommentUseCase.GetPhotoById(id, user.ID)
	if err != nil {
		log.Println(err)
		if error.IsSame(err, error.ErrPhotoNotFound) {
			errMessage := api.SetError("Photo Not Found!")
			resp := api.APIResponse("Get Photo By ID Failed", http.StatusNotFound, "error", errMessage)
			c.JSON(http.StatusNotFound, resp)
			return
		}
		resp := api.APIResponse("Get Photo By ID Failed", http.StatusInternalServerError, "error", nil)
		c.JSON(http.StatusInternalServerError, resp)
		return
	}
	resp := api.APIResponse("Get Photo By ID Success", http.StatusOK, "success", photo)
	c.JSON(http.StatusOK, resp)
}

// create photo
func (cc *CommentController) CreatePhoto(c *gin.Context) {
	// get data from body input
	var input domain.InsertPhoto
	if err := c.ShouldBindJSON(&input); err != nil {
		log.Println(err)
		var validation validator.ValidationErrors
		if errors.As(err, &validation) {
			result := make([]error.Form, len(validation))
			for i, v := range validation {
				result[i] = error.Form{
					Field:   v.Field(),
					Message: error.FormValidationError(v),
				}
			}
			resp := api.APIResponse("Create Photo Failed", http.StatusBadRequest, "error", result)
			c.JSON(http.StatusBadRequest, resp)
			return
		}
		errorMessage := api.SetError(err.Error())
		resp := api.APIResponse("Create Photo Failed", http.StatusBadRequest, "error", errorMessage)
		c.JSON(http.StatusBadRequest, resp)
		return
	}

	user := c.MustGet("currentUser").(domainUser.User)
	input.UserID = user.ID
	photo, err := cc.CommentUseCase.CreatePhoto(input)
	if err != nil {
		log.Println(err)
		resp := api.APIResponse("Create Photo Failed", http.StatusInternalServerError, "error", nil)
		c.JSON(http.StatusInternalServerError, resp)
		return
	}
	resp := api.APIResponse("Create Photo Success", http.StatusOK, "success", photo)
	c.JSON(http.StatusOK, resp)
}

// update photo
func (cc *CommentController) UpdatePhoto(c *gin.Context) {
	// get id from param
	id := c.Param("id")
	var input domain.UpdatePhoto
	if err := c.ShouldBindJSON(&input); err != nil {
		log.Println(err)
		var validation validator.ValidationErrors
		if errors.As(err, &validation) {
			result := make([]error.Form, len(validation))
			for i, v := range validation {
				result[i] = error.Form{
					Field:   v.Field(),
					Message: error.FormValidationError(v),
				}
			}
			resp := api.APIResponse("Update Photo Failed", http.StatusBadRequest, "error", result)
			c.JSON(http.StatusBadRequest, resp)
			return
		}
		errorMessage := api.SetError(err.Error())
		resp := api.APIResponse("Update Photo Failed", http.StatusBadRequest, "error", errorMessage)
		c.JSON(http.StatusBadRequest, resp)
		return
	}

	if input.Caption == "" && input.Title == "" && input.Photo_url == "" {
		errMessage := api.SetError("Caption, Title, Photo_url cannot be empty!")
		resp := api.APIResponse("Update Photo Failed", http.StatusBadRequest, "error", errMessage)
		c.JSON(http.StatusBadRequest, resp)
		return
	}

	user := c.MustGet("currentUser").(domainUser.User)
	input.UserID = user.ID
	photo, err := cc.CommentUseCase.UpdatePhoto(id, input)
	if err != nil {
		log.Println(err)
		if error.IsSame(err, error.ErrPhotoNotFound) {
			errMessage := api.SetError("Photo Not Found!")
			resp := api.APIResponse("Update Photo Failed", http.StatusNotFound, "error", errMessage)
			c.JSON(http.StatusNotFound, resp)
			return
		}
		resp := api.APIResponse("Update Photo Failed", http.StatusInternalServerError, "error", nil)
		c.JSON(http.StatusInternalServerError, resp)
		return
	}
	resp := api.APIResponse("Update Photo Success", http.StatusOK, "success", photo)
	c.JSON(http.StatusOK, resp)
}

// delete photo
func (cc *CommentController) DeletePhoto(c *gin.Context) {
	// get id from param
	id := c.Param("id")
	user := c.MustGet("currentUser").(domainUser.User)
	err := cc.CommentUseCase.DeletePhoto(id, user.ID)
	if err != nil {
		log.Println(err)
		if error.IsSame(err, error.ErrPhotoNotFound) {
			errMessage := api.SetError("Photo Not Found!")
			resp := api.APIResponse("Delete Photo Failed", http.StatusNotFound, "error", errMessage)
			c.JSON(http.StatusNotFound, resp)
			return
		}
		resp := api.APIResponse("Delete Photo Failed", http.StatusInternalServerError, "error", nil)
		c.JSON(http.StatusInternalServerError, resp)
		return
	}
	resp := api.APIResponse("Delete Photo Success", http.StatusOK, "success", nil)
	c.JSON(http.StatusOK, resp)
}
