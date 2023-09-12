package controllers

import (
	userRepository "instagam/modules/v1/users/interfaces/repositories"
	userUseCase "instagam/modules/v1/users/usecases"

	"gorm.io/gorm"
)

type UserController struct {
	UserUseCase *userUseCase.UserUseCase
}

func NewUserController(db *gorm.DB) *UserController {
	repo := userRepository.NewUserRepository(db)
	cu := userUseCase.NewUserUseCase(repo)
	return &UserController{
		UserUseCase: cu,
	}
}
