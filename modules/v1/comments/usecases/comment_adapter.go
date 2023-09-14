package usecases

import (
	"instagam/modules/v1/comments/domain"
	commentRepository "instagam/modules/v1/comments/interfaces/repositories"
)

type CommentAdapter interface {
	GetAllComments(idPhotos string, idUser int) ([]domain.Comment, error)
	GetCommentById(id string) (domain.Comment, error)
	CreateComment(input domain.InsertComment) (domain.Comment, error)
	UpdateComment(idComments string, input domain.UpdateComment, idUser int) (domain.Comment, error)
	DeleteComment(idComment string, idUser int) error
	GetAllPhotos() ([]domain.Photo, error)
	GetPhotoById(idPhotos string, idUser int) (domain.Photo, error)
	CreatePhoto(input domain.InsertPhoto) (domain.CreatedPhoto, error)
	UpdatePhoto(id string, input domain.UpdatePhoto) (domain.CreatedPhoto, error)
	DeletePhoto(idPhoto string, idUser int) error
}

type CommentUseCase struct {
	repoComment *commentRepository.Repository
}

func NewCommentUseCase(repoComment *commentRepository.Repository) *CommentUseCase {
	return &CommentUseCase{repoComment}
}
