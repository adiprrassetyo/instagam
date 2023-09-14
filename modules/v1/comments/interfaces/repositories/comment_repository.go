package repository

import (
	"instagam/modules/v1/comments/domain"

	"gorm.io/gorm/clause"
)

func (r *Repository) FindAllComments(idPhotos string, idUser int) ([]domain.Comment, error) {
	var comments []domain.Comment
	err := r.db.Where("photo_id = ? AND user_id = ?", idPhotos, idUser).Find(&comments).Error
	return comments, err
}

func (r *Repository) FindCommentById(id string) (domain.Comment, error) {
	var comment domain.Comment
	err := r.db.Where("id = ?", id).First(&comment).Error
	return comment, err
}

func (r *Repository) SaveComment(comment domain.Comment) (domain.Comment, error) {
	err := r.db.Create(&comment).Error
	return comment, err
}

func (r *Repository) UpdateComment(comment domain.Comment, id string) (domain.Comment, error) {
	err := r.db.Model(&comment).Clauses(clause.Returning{}).Where("id = ?", id).Updates(&comment).Error
	return comment, err
}

func (r *Repository) DeleteComment(id string) error {
	return r.db.Where("id = ?", id).Delete(&domain.Comment{}).Error
}

func (r *Repository) FindAllPhoto() ([]domain.Photo, error) {
	var photos []domain.Photo
	err := r.db.Preload("User").Preload("Comments").Find(&photos).Error
	return photos, err
}

func (r *Repository) FindPhotoById(id int) (domain.Photo, error) {
	var photo domain.Photo
	err := r.db.Preload("User").Preload("Comments").Where("id = ?", id).First(&photo).Error
	return photo, err
}

func (r *Repository) FindPhoto(idPhotos string, idUser int) (domain.Photo, error) {
	var photo domain.Photo
	err := r.db.Preload("User").Preload("Comments").Where("id = ? AND user_id = ?", idPhotos, idUser).First(&photo).Error
	return photo, err
}

func (r *Repository) SavePhoto(photo domain.CreatedPhoto) (domain.CreatedPhoto, error) {
	err := r.db.Create(&photo).Error
	return photo, err
}

func (r *Repository) UpdatePhoto(photo domain.CreatedPhoto, id string) (domain.CreatedPhoto, error) {
	err := r.db.Model(&photo).Clauses(clause.Returning{}).Where("id = ?", id).Updates(&photo).Error
	return photo, err
}

func (r *Repository) DeletePhoto(id string) error {
	return r.db.Where("id = ?", id).Delete(&domain.Photo{}).Error
}
