package main

import (
	"todo/models"
	"todo/router"

	"github.com/gin-gonic/gin"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

func main() {
	//router.InitializeRoutes()
	r := gin.Default()

	db := models.SetupModels() // new

	// Middleware Provide db variable to controllers
	r.Use(func(c *gin.Context) {
		c.Set("db", db)
		c.Next()
	})

	router.InitializeRoutes(r)

	r.Run()
}
