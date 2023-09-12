package routes

import (
	userControllerV1 "instagam/modules/v1/users/interfaces/controllers"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func NewRouter(router *gin.Engine, db *gorm.DB) *gin.Engine {
	userControllerV1 := userControllerV1.NewUserController(db)

	//User
	api := router.Group("/api/v1")
	{
		users := api.Group("/users")
		{
			users.POST("/register", userControllerV1.Register)
		}
	}
	return router
}
