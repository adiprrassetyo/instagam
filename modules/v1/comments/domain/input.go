package domain

type InsertComment struct {
	UserID   int    `json:"-" gorm:"column:user_id"`
	Photo_id int    `json:"photo_id" gorm:"column:photo_id" binding:"required,number"`
	Message  string `json:"message" gorm:"column:message" binding:"required"`
}

type UpdateComment struct {
	UserID  int    `json:"-" gorm:"column:user_id"`
	PhotoID int    `json:"photo_id" gorm:"column:photo_id" binding:"omitempty,number"`
	Message string `json:"message" gorm:"column:message" binding:"omitempty"`
}

type InsertPhoto struct {
	Title     string `json:"title" gorm:"column:title" binding:"required"`
	Caption   string `json:"caption" gorm:"column:caption" binding:"required"`
	Photo_url string `json:"photo_url" gorm:"column:photo_url" binding:"required,url"`
	UserID    int    `json:"-" gorm:"column:user_id"`
}

type UpdatePhoto struct {
	Title     string `json:"title" gorm:"column:title" binding:"omitempty"`
	Caption   string `json:"caption" gorm:"column:caption" binding:"omitempty"`
	Photo_url string `json:"photo_url" gorm:"column:photo_url" binding:"omitempty,url"`
	UserID    int    `json:"-" gorm:"column:-"`
}
