package usecases

import (
	"instagam/modules/v1/comments/domain"
	errorHandling "instagam/pkg/http-error"
)

func (cu *CommentUseCase) GetAllComments(idPhotos string, idUser int) ([]domain.Comment, error) {
	comment, err := cu.repoComment.FindAllComments(idPhotos, idUser)
	if err != nil {
		return nil, err
	}

	if len(comment) == 0 {
		return nil, errorHandling.ErrDataNotFound
	}

	return comment, nil
}

func (cu *CommentUseCase) GetCommentById(id string) (domain.Comment, error) {
	comment, err := cu.repoComment.FindCommentById(id)
	if err != nil {
		return domain.Comment{}, err
	}

	if comment.ID == 0 {
		return domain.Comment{}, errorHandling.ErrDataNotFound
	}

	return comment, nil
}

func (cu *CommentUseCase) CreateComment(input domain.InsertComment) (domain.Comment, error) {
	//Validate Photo Exist
	photo, err := cu.repoComment.FindPhotoById(input.Photo_id)
	if err != nil {
		return domain.Comment{}, errorHandling.ErrPhotoNotFound
	}

	if photo.ID == 0 {
		return domain.Comment{}, errorHandling.ErrPhotoNotFound
	}

	comment := domain.Comment{
		PhotoID: input.Photo_id,
		UserID:  input.UserID,
		Message: input.Message,
	}

	return cu.repoComment.SaveComment(comment)
}

func (cu *CommentUseCase) UpdateComment(idComments string, input domain.UpdateComment, idUser int) (domain.Comment, error) {
	if input.PhotoID != 0 {
		//Validate Photo Exist
		photo, err := cu.repoComment.FindPhotoById(input.PhotoID)
		if err != nil {
			return domain.Comment{}, errorHandling.ErrPhotoNotFound
		}

		if photo.ID == 0 {
			return domain.Comment{}, errorHandling.ErrPhotoNotFound
		}
	}
	//Check Comment Exist
	comment, err := cu.GetCommentById(idComments)
	if err != nil {
		return domain.Comment{}, err
	}

	if comment.ID == 0 || comment.UserID != idUser {
		return domain.Comment{}, errorHandling.ErrCommentNotFound
	}

	updateComment := domain.Comment{
		PhotoID: input.PhotoID,
		Message: input.Message,
	}

	return cu.repoComment.UpdateComment(updateComment, idComments)
}

func (cu *CommentUseCase) DeleteComment(idComment string, idUser int) error {
	//Check Comment Exist
	comment, err := cu.repoComment.FindCommentById(idComment)
	if err != nil {
		if errorHandling.IsSame(err, errorHandling.ErrDataNotFound) {
			return errorHandling.ErrCommentNotFound
		}
		return err
	}

	if comment.ID == 0 || comment.UserID != idUser {
		return errorHandling.ErrCommentNotFound
	}

	return cu.repoComment.DeleteComment(idComment)
}

func (cu *CommentUseCase) GetAllPhotos() ([]domain.Photo, error) {
	photo, err := cu.repoComment.FindAllPhoto()
	if err != nil {
		if errorHandling.IsSame(err, errorHandling.ErrDataNotFound) {
			return nil, errorHandling.ErrDataNotFound
		}
		return nil, err
	}

	if len(photo) == 0 {
		return nil, errorHandling.ErrDataNotFound
	}

	return photo, nil
}

func (cu *CommentUseCase) GetPhotoById(idPhotos string, idUser int) (domain.Photo, error) {
	photo, err := cu.repoComment.FindPhoto(idPhotos, idUser)
	if err != nil {
		if errorHandling.IsSame(err, errorHandling.ErrDataNotFound) {
			return domain.Photo{}, errorHandling.ErrPhotoNotFound
		}
		return domain.Photo{}, err
	}

	if photo.ID == 0 {
		return domain.Photo{}, errorHandling.ErrPhotoNotFound
	}

	return photo, nil
}

func (cu *CommentUseCase) CreatePhoto(input domain.InsertPhoto) (domain.CreatedPhoto, error) {
	photo := domain.CreatedPhoto{
		Title:    input.Title,
		Caption:  input.Caption,
		PhotoUrl: input.Photo_url,
		UserID:   input.UserID,
	}

	return cu.repoComment.SavePhoto(photo)
}

func (cu *CommentUseCase) UpdatePhoto(id string, input domain.UpdatePhoto) (domain.CreatedPhoto, error) {
	//Check Photo Exist
	photo, err := cu.repoComment.FindPhoto(id, input.UserID)
	if err != nil {
		if errorHandling.IsSame(err, errorHandling.ErrDataNotFound) {
			return domain.CreatedPhoto{}, errorHandling.ErrPhotoNotFound
		}
		return domain.CreatedPhoto{}, err
	}

	if photo.ID == 0 {
		return domain.CreatedPhoto{}, errorHandling.ErrPhotoNotFound
	}

	updatePhoto := domain.CreatedPhoto{
		Title:    input.Title,
		Caption:  input.Caption,
		PhotoUrl: input.Photo_url,
	}

	return cu.repoComment.UpdatePhoto(updatePhoto, id)
}

func (cu *CommentUseCase) DeletePhoto(idPhoto string, idUser int) error {
	photo, err := cu.repoComment.FindPhoto(idPhoto, idUser)
	if err != nil {
		if errorHandling.IsSame(err, errorHandling.ErrDataNotFound) {
			return errorHandling.ErrPhotoNotFound
		}
		return err
	}

	if photo.ID == 0 {
		return errorHandling.ErrPhotoNotFound
	}

	return cu.repoComment.DeletePhoto(idPhoto)
}
