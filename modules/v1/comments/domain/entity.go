package domain

import "time"

type GormModel struct {
	ID        int        `json:"id" gorm:"column:id"`
	CreatedAt *time.Time `json:"created_at" gorm:"column:created_at"`
	UpdatedAt *time.Time `json:"updated_at" gorm:"column:updated_at"`
}

type Comment struct {
	GormModel
	UserID  int    `json:"user_id" gorm:"column:user_id"`
	PhotoID int    `json:"photo_id" gorm:"column:photo_id"`
	Message string `json:"message" gorm:"column:message"`
}

type Photo struct {
	GormModel
	Title    string         `json:"title" gorm:"column:title"`
	Caption  string         `json:"caption" gorm:"column:caption"`
	PhotoUrl string         `json:"photo_url" gorm:"column:photo_url"`
	UserID   int            `json:"user_id" gorm:"column:user_id"`
	User     *UserPhoto     `json:"user"`
	Comments []CommentPhoto `json:"comments"`
}

type UserPhoto struct {
	GormModel
	UserName string `json:"username" gorm:"column:username"`
	Email    string `json:"email" gorm:"column:email"`
	Age      int    `json:"age" gorm:"column:age"`
}

func (UserPhoto) TableName() string {
	return "users"
}

type CommentPhoto struct {
	GormModel
	UserID  int    `json:"user_id" gorm:"column:user_id"`
	PhotoID int    `json:"photo_id" gorm:"column:photo_id"`
	Message string `json:"message" gorm:"column:message"`
}

func (CommentPhoto) TableName() string {
	return "comments"
}

type CreatedPhoto struct {
	GormModel
	Title    string `json:"title" gorm:"column:title"`
	Caption  string `json:"caption" gorm:"column:caption"`
	PhotoUrl string `json:"photo_url" gorm:"column:photo_url"`
	UserID   int    `json:"user_id" gorm:"column:user_id"`
}

func (CreatedPhoto) TableName() string {
	return "photos"
}
