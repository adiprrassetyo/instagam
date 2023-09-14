package controllers

import (
	commentRepository "instagam/modules/v1/comments/interfaces/repositories"
	commentUseCase "instagam/modules/v1/comments/usecases"

	"gorm.io/gorm"
)

type CommentController struct {
	CommentUseCase *commentUseCase.CommentUseCase
}

func NewCommentController(db *gorm.DB) *CommentController {
	repo := commentRepository.NewCommentRepository(db)
	cu := commentUseCase.NewCommentUseCase(repo)
	return &CommentController{
		CommentUseCase: cu,
	}
}
