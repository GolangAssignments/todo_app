package controllers

import (
	"net/http"
	"strconv"
	"todo_app/database"
	"todo_app/dtos"
	"todo_app/models"

	"github.com/gin-gonic/gin"
)

func GetTodos(c *gin.Context) {
	todoList := []models.Todo{}

	// Use Unscoped() to fetch all, including soft-deleted records
	// db.Unscoped().Find(&todos)

	database.DB.Find(&todoList)

	todoResponseList := []dtos.SingleTodoResponse{}
	for _, todo := range todoList {
		todoResponse := dtos.SingleTodoResponse{
			ID:        todo.ID,
			Title:     todo.Title,
			Completed: todo.Completed,
			CreatedAt: todo.CreatedAt,
			UpdatedAt: todo.UpdatedAt,
		}

		todoResponseList = append(todoResponseList, todoResponse)
	}

	finalTodoResponseList := dtos.TodoListResponse{
		Todos: todoResponseList,
	}
	c.JSON(http.StatusOK, finalTodoResponseList)
}

func CreateTodo(c *gin.Context) {
	var createTodoRequest dtos.CreateTodoRequest

	if err := c.ShouldBindJSON(&createTodoRequest); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "ShouldBindJSON: " + err.Error(),
		})
		return
	}

	todo := models.Todo{
		Title: createTodoRequest.Title,
	}

	err := database.DB.Create(&todo).Error
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}

	todoResponse := dtos.SingleTodoResponse{
		ID:        todo.ID,
		Title:     todo.Title,
		Completed: todo.Completed,
		CreatedAt: todo.CreatedAt,
		UpdatedAt: todo.UpdatedAt,
	}

	c.JSON(http.StatusOK, todoResponse)
}

func UpdateTodo(c *gin.Context) {
	updateTodoRequest := dtos.UpdateTodoRequest{}

	id, _ := strconv.Atoi(c.Param("id"))

	if err := c.ShouldBindJSON(&updateTodoRequest); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "ShouldBindJSON: " + err.Error(),
		})
		return
	}

	if updateTodoRequest.Title == nil && updateTodoRequest.Completed == nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Title or Completed is required!",
		})
		return
	}

	result := database.DB.Model(&models.Todo{}).Where("id = ?", id).Updates(&updateTodoRequest)

	if result.RowsAffected == 0 {
		c.JSON(http.StatusNotFound, gin.H{
			"error": "Todo not found!",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "Todo updated successfully!",
	})
}

func DeleteTodo(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))

	result := database.DB.Delete(&models.Todo{}, id)

	if result.RowsAffected == 0 {
		c.JSON(http.StatusNotFound, gin.H{
			"error": "Todo not found!",
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"message": "Todo deleted successfully!",
	})
}
