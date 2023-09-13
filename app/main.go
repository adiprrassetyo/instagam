package main

import (
	"instagam/infrastructures/config"
	database "instagam/infrastructures/databases"
	routesUsersV1 "instagam/modules/v1/users/routes"
	error "instagam/pkg/http-error"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func main() {
	config := config.New()

	router := gin.Default()
	router.Use(cors.Default())
	db := database.NewDatabases()

	router = routesUsersV1.NewRouter(router, db)

	router.NoRoute(error.NotFound())
	router.NoMethod(error.NoMethod())

	// Listen and Server in
	router.Run(":" + config.App.Port)
}
