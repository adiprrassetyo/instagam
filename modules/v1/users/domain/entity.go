package domain

import "time"

type GormModel struct {
	ID        int        `json:"id" gorm:"column:id"`
	CreatedAt *time.Time `json:"created_at" gorm:"column:created_at"`
	UpdatedAt *time.Time `jsod:"updated_at" gorm:"column:updated_at"`
}

type User struct {
	GormModel
	UserName string `json:"username" gorm:"column:username"`
	Email    string `json:"email" gorm:"column:email"`
	Password string `json:"-" gorm:"column:password"`
	Age      int    `json:"age" gorm:"column:age"`
}

type SocialMedia struct {
	GormModel
	Name             string           `json:"name" gorm:"column:name"`
	Social_media_url string           `json:"social_media_url" gorm:"column:social_media_url"`
	UserID           int              `json:"-" gorm:"column:user_id"`
	User             *UserSocialMedia `json:"user"`
}

type UserSocialMedia struct {
	ID       int    `json:"id" gorm:"column:id"`
	UserName string `json:"username" gorm:"column:username"`
	Email    string `json:"email" gorm:"column:email"`
	Age      int    `json:"age" gorm:"column:age"`
}

func (UserSocialMedia) TableName() string {
	return "users"
}

type CreatedSocialMedia struct {
	GormModel
	Name             string `json:"name" gorm:"column:name"`
	Social_media_url string `json:"social_media_url" gorm:"column:social_media_url"`
	UserID           int    `json:"-" gorm:"column:user_id"`
}
