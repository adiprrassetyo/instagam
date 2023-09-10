package repository

import (
	"instagam/modules/v1/users/domain"

	"gorm.io/gorm/clause"
)

func (r *Repository) Save(user domain.User) (domain.User, error) {
	err := r.db.Create(&user).Error
	return user, err
}

func (r *Repository) FindUser(field string, value string) (domain.User, error) {
	var user domain.User
	err := r.db.Where(field+" = ?", value).First(&user).Error
	return user, err
}

func (r *Repository) FindUserByID(id int) (domain.User, error) {
	var user domain.User
	err := r.db.Where("id = ?", id).First(&user).Error
	return user, err
}

func (r *Repository) AllSocialMedia() ([]domain.SocialMedia, error) {
	var socialMedia []domain.SocialMedia
	err := r.db.Preload("User").Find(&socialMedia).Error
	return socialMedia, err
}

func (r *Repository) FindSocialMediaByID(id string) (domain.SocialMedia, error) {
	var socialMedia domain.SocialMedia
	err := r.db.Preload("User").Where("id = ?", id).First(&socialMedia).Error
	return socialMedia, err
}

func (r *Repository) SaveSocialMedia(socialMedia domain.SocialMedia) (domain.SocialMedia, error) {
	err := r.db.Create(&socialMedia).Error
	return socialMedia, err
}

func (r *Repository) FindSocialMediaByUserID(id int) (domain.SocialMedia, error) {
	var socialMedia domain.SocialMedia
	err := r.db.Preload("User").Where("user_id = ?", id).First(&socialMedia).Error
	return socialMedia, err
}

func (r *Repository) UpdateSocialMedia(socialMedia domain.SocialMedia, id int) (domain.SocialMedia, error) {
	err := r.db.Model(&socialMedia).Clauses(clause.Returning{}).Where("id = ?", id).Updates(socialMedia).Error
	return socialMedia, err
}

func (r *Repository) DeleteSocialMedia(id int) error {
	return r.db.Delete(&domain.SocialMedia{}, id).Error
}
