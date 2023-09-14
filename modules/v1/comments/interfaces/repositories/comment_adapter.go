package repository

import (
	"instagam/modules/v1/comments/domain"

	"gorm.io/gorm"
)

type RepositoryPresenter interface {
	FindAllComments(idPhotos string, idUser int) ([]domain.Comment, error)
	FindCommentById(id string) (domain.Comment, error)
	SaveComment(comment domain.Comment) (domain.Comment, error)
	UpdateComment(comment domain.Comment, id string) (domain.Comment, error)
	DeleteComment(id string) error
	FindAllPhoto() ([]domain.Photo, error)
	FindPhotoById(id int) (domain.Photo, error)
	FindPhoto(idPhotos string, idUser int) (domain.Photo, error)
	SavePhoto(photo domain.CreatedPhoto) (domain.CreatedPhoto, error)
	UpdatePhoto(photo domain.CreatedPhoto, id string) (domain.CreatedPhoto, error)
	DeletePhoto(id string) error
}

type Repository struct {
	db *gorm.DB
}

func NewCommentRepository(db *gorm.DB) *Repository {
	return &Repository{db}
}
