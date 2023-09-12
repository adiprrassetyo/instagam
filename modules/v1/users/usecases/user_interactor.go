package usecases

import (
	"fmt"
	"instagam/modules/v1/users/domain"
	"strconv"

	errorHandling "instagam/pkg/http-error"
	"instagam/pkg/times"

	"golang.org/x/crypto/bcrypt"
)

func (u *UserUseCase) RegisterUser(input domain.RegisterUserInput) (domain.User, error) {
	user := domain.User{}
	user.UserName = input.Username
	user.Email = input.Email
	user.Age = input.Age

	passwordHash, err := bcrypt.GenerateFromPassword([]byte(input.Password), bcrypt.MinCost)
	if err != nil {
		return user, err
	}
	user.Password = string(passwordHash)
	newUser, err := u.repoUser.Save(user)
	if err != nil {
		return user, err
	}

	return newUser, nil
}

func (u *UserUseCase) LoginUser(input domain.LoginUserInput) (domain.User, error) {
	var (
		user domain.User
		err  error
	)

	if input.Email != "" {
		user, err = u.repoUser.FindUser("email", input.Email)
		if err != nil {
			if errorHandling.IsSame(err, errorHandling.ErrDataNotFound) {
				return user, errorHandling.ErrEmailNotFound
			}
			return user, err
		}
	} else {
		user, err = u.repoUser.FindUser("username", input.Username)
		if err != nil {
			fmt.Println("Ga ada data username")
			if errorHandling.IsSame(err, errorHandling.ErrDataNotFound) {
				fmt.Println("custome error")
				return user, errorHandling.ErrUsernameNotFound
			}
			return user, err
		}
	}

	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(input.Password))
	if err != nil {
		return user, err
	}

	return user, nil
}

func (u *UserUseCase) GetUserByID(id int) (domain.User, error) {
	user, err := u.repoUser.FindUserByID(id)
	if err != nil {
		return user, err
	}

	if user.ID == 0 {
		return user, errorHandling.ErrUserNotFound
	}
	return user, nil
}

func (u *UserUseCase) AllSocialMedia() ([]domain.SocialMedia, error) {
	return u.repoUser.AllSocialMedia()
}

func (u *UserUseCase) OneSocialMedia(id string) (domain.SocialMedia, error) {
	return u.repoUser.FindSocialMediaByID(id)
}

func (u *UserUseCase) CreateSocialMedia(input domain.InsertSocialMedia, id int) (domain.CreatedSocialMedia, error) {
	socialMedia := domain.SocialMedia{
		Name:             input.Name,
		Social_media_url: input.Social_media_url,
		UserID:           id,
	}
	newSocialMedia := domain.CreatedSocialMedia{}
	socialMedia, err := u.repoUser.SaveSocialMedia(socialMedia)
	if err != nil {
		return newSocialMedia, err
	}

	newSocialMedia.ID = socialMedia.ID
	newSocialMedia.Name = socialMedia.Name
	newSocialMedia.Social_media_url = socialMedia.Social_media_url
	newSocialMedia.UserID = socialMedia.UserID
	newSocialMedia.UpdatedAt = socialMedia.UpdatedAt
	newSocialMedia.CreatedAt = socialMedia.CreatedAt

	return newSocialMedia, nil
}

func (u *UserUseCase) CheckSocialMedia(id int) error {
	user, err := u.repoUser.FindSocialMediaByUserID(id)
	if err != nil {
		return err
	}
	if user.ID != 0 {
		return errorHandling.ErrSocialMediaAlreadyExist
	}
	return nil
}

func (u *UserUseCase) UpdateSocialMedia(input domain.UpdateSocialMedia, id_sosmed string, id_user int) (domain.CreatedSocialMedia, error) {
	var (
		now            = times.Now("Asia/Jakarta")
		newSocialMedia = domain.CreatedSocialMedia{}
		socialMedia    = domain.SocialMedia{}
	)

	id_sos, err := strconv.Atoi(id_sosmed)
	if err != nil {
		return domain.CreatedSocialMedia{}, err
	}

	socialMedia_user, err := u.repoUser.FindSocialMediaByUserID(id_user)
	if err != nil {
		return newSocialMedia, err
	}
	//Not Permitted Because Social Media doesn't belong to user
	if socialMedia_user.ID != id_sos {
		return newSocialMedia, errorHandling.ErrSocialMediaNotFound
	}

	socialMedia.Name = input.Name
	socialMedia.Social_media_url = input.Social_media_url
	socialMedia.UpdatedAt = &now
	socialMedia, err = u.repoUser.UpdateSocialMedia(socialMedia, id_sos)
	if err != nil {
		return newSocialMedia, err
	}

	newSocialMedia.ID = socialMedia.ID
	newSocialMedia.Name = socialMedia.Name
	newSocialMedia.Social_media_url = socialMedia.Social_media_url
	newSocialMedia.UserID = socialMedia.UserID
	newSocialMedia.UpdatedAt = socialMedia.UpdatedAt
	newSocialMedia.CreatedAt = socialMedia.CreatedAt

	return newSocialMedia, nil
}

func (u *UserUseCase) DeleteSocialMedia(id_sosmed string, id_user int) error {
	id_sos, err := strconv.Atoi(id_sosmed)
	if err != nil {
		return err
	}

	socialMedia_user, err := u.repoUser.FindSocialMediaByUserID(id_user)
	if err != nil {
		return err
	}
	//Not Permitted Because Social Media doesn't belong to user
	if socialMedia_user.ID != id_sos {
		return errorHandling.ErrSocialMediaNotFound
	}

	return u.repoUser.DeleteSocialMedia(id_sos)
}
