package database

import "time"

type GormModel struct {
	ID        int        `json:"id" gorm:"column:id;type:int;primaryKey;autoIncrement;not null"`
	CreatedAt *time.Time `json:"created_at" gorm:"column:created_at;type:timestamp;default:CURRENT_TIMESTAMP"`
	UpdatedAt *time.Time `json:"updated_at" gorm:"column:updated_at;type:timestamp;default:CURRENT_TIMESTAMP"`
}

type User struct {
	GormModel
	UserName    string       `json:"username" gorm:"column:username;type:varchar(100);unique;not null"`
	Email       string       `json:"email" gorm:"column:email;type:varchar(100);unique;not null"`
	Password    string       `json:"password" gorm:"column:password;type:varchar(200);not null"`
	Age         int          `json:"age" gorm:"column:age;type:int;not null"`
	Photos      []Photo      `json:"photos" gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL;foreignKey:user_id;references:ID"`
	Comments    []Comment    `json:"comments" gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL;foreignKey:user_id;references:ID"`
	SocialMedia *SocialMedia `json:"social_media" gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL;foreignKey:user_id;references:ID"`
}

type Photo struct {
	GormModel
	Title    string    `json:"title" gorm:"column:title;type:varchar(100);not null"`
	Caption  string    `json:"caption" gorm:"column:caption;type:varchar(500);not null"`
	PhotoUrl string    `json:"photo_url" gorm:"column:photo_url;type:varchar(200);not null"`
	UserID   int       `json:"user_id" gorm:"column:user_id;type:int;not null"`
	User     *User     `json:"user" gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL"`
	Comments []Comment `json:"comments" gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL;foreignKey:photo_id;references:ID"`
}

type Comment struct {
	GormModel
	UserID  int    `json:"user_id" gorm:"column:user_id;type:int;not null"`
	PhotoID int    `json:"photo_id" gorm:"column:photo_id;type:int;not null"`
	Message string `json:"message" gorm:"column:message;type:varchar(500);not null"`
	Photo   *Photo `json:"photo" gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL"`
	User    *User  `json:"user" gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL"`
}

type SocialMedia struct {
	GormModel
	Name             string `json:"name" gorm:"column:name;type:varchar(100);not null"`
	Social_media_url string `json:"social_media_url" gorm:"column:social_media_url;type:varchar(200);not null"`
	UserID           int    `json:"user_id" gorm:"column:user_id;type:int;not null"`
	User             *User  `json:"user" gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL"`
}
