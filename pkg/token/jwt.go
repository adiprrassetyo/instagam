package token

import (
	"errors"
	"instagam/infrastructures/config"
	"time"

	"github.com/dgrijalva/jwt-go"
)

func GenerateToken(userID int) (string, error) {
	config := config.New()
	claims := jwt.MapClaims{}
	// Claims is a set of key/value pairs that are stored in a JWT.
	claims["authorized"] = true
	claims["user_id"] = userID
	claims["exp"] = time.Now().Add(time.Minute * 30).Unix()

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	return token.SignedString([]byte(config.App.Secret_key))
}

func ValidateToken(encoded string) (*jwt.Token, error) {
	config := config.New()
	token, err := jwt.Parse(encoded, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, errors.New("invalid token")
		}
		return []byte(config.App.Secret_key), nil
	})

	if err != nil {
		return nil, err
	}

	return token, nil
}
