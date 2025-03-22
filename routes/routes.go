package routes

import (
	"todo_app/controllers"
	"github.com/gin-gonic/gin"
)

func SetupRoutes() *gin.Engine {
	r := gin.Default()
	r.GET("/todos", controllers.GetTodos)
	r.POST("/todos", controllers.CreateTodo)
	r.DELETE("/todos/:id", controllers.DeleteTodo)
	r.PATCH("/todos/:id", controllers.UpdateTodo)
	return r
}