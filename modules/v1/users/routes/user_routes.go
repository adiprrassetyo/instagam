package routes

import (
	userControllerV1 "instagam/modules/v1/users/interfaces/controllers"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func NewRouter(router *gin.Engine, db *gorm.DB) *gin.Engine {
	userControllerV1 := userControllerV1.NewUserController(db)

	api := router.Group("/api/v1")
	{
		// User
		users := api.Group("/users")
		{
			users.POST("/register", userControllerV1.Register)
			users.POST("/login", userControllerV1.Login)
		}
		// Social media
		social := api.Group("/media")
		{
			social.GET("", userControllerV1.GetAllSocialMedia)
			social.GET("/:id", userControllerV1.GetOneSocialMedia)
			social.POST("", userControllerV1.CreateSocialMedia)
			social.PUT("/:id", userControllerV1.UpdateSocialMedia)
			social.DELETE("/:id", userControllerV1.DeleteSocialMedia)
		}
	}

	return router
}
