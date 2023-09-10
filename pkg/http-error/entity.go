package errorsHandling

import "errors"

var (
	ErrEmailAlreadyExist       = errors.New("ERROR: duplicate key value violates unique constraint \"users_email_key\" (SQLSTATE 23505)")
	ErrUsernameAlreadyExist    = errors.New("ERROR: duplicate key value violates unique constraint \"users_username_key\" (SQLSTATE 23505)")
	ErrDataNotFound            = errors.New("record not found")
	ErrEmailNotFound           = errors.New("email not found")
	ErrUsernameNotFound        = errors.New("username not found")
	ErrUserNotFound            = errors.New("no user found on with that id")
	ErrUserAlreadyExist        = errors.New("user already exist")
	ErrSocialMediaAlreadyExist = errors.New("social media already exist")
	ErrSocialMediaNotFound     = errors.New("social media not found")
	ErrPhotoNotFound           = errors.New("photo not found")
	ErrCommentNotFound         = errors.New("comment not found")
	ErrDataLoginNotFound       = errors.New("crypto/bcrypt: hashedPassword is not the hash of the given password")
)

type Form struct {
	Field   string
	Message string
}
