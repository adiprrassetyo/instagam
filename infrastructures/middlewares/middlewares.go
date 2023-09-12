package middlewares

import (
	userRepository "instagam/modules/v1/users/interfaces/repositories"

	userUseCase "instagam/modules/v1/users/usecases"
	api "instagam/pkg/api_response"
	token_lib "instagam/pkg/token"
	"net/http"
	"strings"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type Middleware struct {
	userUseCase *userUseCase.UserUseCase
}

func NewMiddleware(db *gorm.DB) *Middleware {
	repo := userRepository.NewUserRepository(db)
	cu := userUseCase.NewUserUseCase(repo)
	return &Middleware{
		userUseCase: cu,
	}
}

func (m *Middleware) Auth() gin.HandlerFunc {
	return func(c *gin.Context) {
		header := c.Request.Header.Get("Authorization")
		if !strings.Contains(header, "Bearer") {
			response := api.APIResponse("Unauthorized", http.StatusUnauthorized, "error", nil)
			c.AbortWithStatusJSON(http.StatusUnauthorized, response)
			return
		}

		tokenString := ""
		arrayToken := strings.Split(header, " ")
		if len(arrayToken) == 2 {
			tokenString = arrayToken[1]
		}

		token, err := token_lib.ValidateToken(tokenString)
		if err != nil {
			response := api.APIResponse("Unauthorized", http.StatusUnauthorized, "error", nil)
			c.AbortWithStatusJSON(http.StatusUnauthorized, response)
			return
		}

		claim, ok := token.Claims.(jwt.MapClaims)
		if !ok || !token.Valid {
			response := api.APIResponse("Unauthorized", http.StatusUnauthorized, "error", nil)
			c.AbortWithStatusJSON(http.StatusUnauthorized, response)
			return
		}

		userID := int(claim["user_id"].(float64))
		user, err := m.userUseCase.GetUserByID(userID)
		if err != nil {
			response := api.APIResponse("Unauthorized", http.StatusUnauthorized, "error", nil)
			c.AbortWithStatusJSON(http.StatusUnauthorized, response)
			return
		}
		c.Set("currentUser", user)
	}
}
