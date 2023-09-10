package main

import (
	"instagam/infrastructures/config"
	database "instagam/infrastructures/databases"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func main() {
	config := config.New()

	router := gin.Default()
	router.Use(cors.Default())
	db := database.NewDatabases()

	router.Run(":" + config.App.Port)
}