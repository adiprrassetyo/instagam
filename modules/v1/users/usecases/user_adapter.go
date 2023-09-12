package usecases

import (
	"instagam/modules/v1/users/domain"
	userRepository "instagam/modules/v1/users/interfaces/repositories"
)

type UserAdapter interface {
	RegisterUser(input domain.RegisterUserInput) (domain.User, error)
	LoginUser(input domain.LoginUserInput) (domain.User, error)
	GetUserByID(id int) (domain.User, error)
	AllSocialMedia() ([]domain.SocialMedia, error)
	OneSocialMedia(id string) (domain.SocialMedia, error)
	CreateSocialMedia(input domain.InsertSocialMedia, id int) (domain.CreatedSocialMedia, error)
	CheckSocialMedia(id int) error
	UpdateSocialMedia(input domain.UpdateSocialMedia, id_sosmed string, id_user int) (domain.CreatedSocialMedia, error)
	DeleteSocialMedia(id_sosmed string, id_user int) error
}

type UserUseCase struct {
	repoUser *userRepository.Repository
}

func NewUserUseCase(repoUser *userRepository.Repository) *UserUseCase {
	return &UserUseCase{repoUser}
}
