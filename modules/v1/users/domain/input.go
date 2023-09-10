package domain

type RegisterUserInput struct {
	Email            string `json:"email" binding:"required,email"`
	Username         string `json:"username" binding:"required"`
	Password         string `json:"password" binding:"required,min=6"`
	Confirm_password string `json:"confirm_password" binding:"required,min=6,eqfield=Password"`
	Age              int    `json:"age" binding:"required,min=8"`
}

type LoginUserInput struct {
	Email    string `json:"email" binding:"required_without=Username"`
	Username string `json:"username" binding:"required_without=Email"`
	Password string `json:"password" binding:"required"`
}

type InsertSocialMedia struct {
	Name             string `json:"name" binding:"required"`
	Social_media_url string `json:"social_media_url" binding:"required,url"`
}

type UpdateSocialMedia struct {
	Name             string `json:"name"`
	Social_media_url string `json:"social_media_url" binding:"omitempty,url"`
}
