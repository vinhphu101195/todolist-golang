package router

import (
	"todo/controller"

	"github.com/gin-gonic/gin"
)

// InitializeRoutes ...
func InitializeRoutes(r *gin.Engine) {

	v1 := r.Group("/api/v1/todos")
	{
		v1.POST("/", controller.CreateTodo)
		v1.GET("/", controller.FetchAllTodo)
		v1.GET("/:id", controller.FetchSingleTodo)
		v1.PUT("/:id", controller.UpdateTodo)
		v1.DELETE("/:id", controller.DeleteTodo)
	}

}
