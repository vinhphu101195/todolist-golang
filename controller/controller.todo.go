package controller

import (
	"net/http"
	"strconv"

	"todo/models"

	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
)

// CreateTodo add a new todo
func CreateTodo(c *gin.Context) {
	db := c.MustGet("db").(*gorm.DB)

	completed, _ := strconv.Atoi(c.PostForm("completed"))
	userid, _ := strconv.Atoi(c.PostForm("userid"))
	todo := models.TodoModel{Title: c.PostForm("title"), Completed: completed, Userid: userid}
	db.Save(&todo)
	c.JSON(http.StatusCreated, gin.H{"status": http.StatusCreated, "message": "Todo item created successfully!", "resourceId": todo.ID})
}

// FetchAllTodo fetch all todos
func FetchAllTodo(c *gin.Context) {
	var todos []models.TodoModel
	var _todos []models.TransformedTodo
	db := c.MustGet("db").(*gorm.DB)

	db.Find(&todos)
	if len(todos) <= 0 {
		c.JSON(http.StatusNotFound, gin.H{"status": http.StatusNotFound, "message": "No todo found!"})
		return
	}
	//transforms the todos for building a good response
	for _, item := range todos {
		completed := false
		if item.Completed == 1 {
			completed = true
		} else {
			completed = false
		}
		_todos = append(_todos, models.TransformedTodo{ID: item.ID, Title: item.Title, Completed: completed, Userid: item.Userid})
	}
	c.JSON(http.StatusOK, gin.H{"status": http.StatusOK, "data": _todos})
}

// FetchSingleTodo fetch a single todo
func FetchSingleTodo(c *gin.Context) {
	db := c.MustGet("db").(*gorm.DB)

	var todo models.TodoModel
	todoID := c.Param("id")
	db.First(&todo, todoID)
	if todo.ID == 0 {
		c.JSON(http.StatusNotFound, gin.H{"status": http.StatusNotFound, "message": "No todo found!"})
		return
	}
	completed := false
	if todo.Completed == 1 {
		completed = true
	} else {
		completed = false
	}
	_todo := models.TransformedTodo{ID: todo.ID, Title: todo.Title, Completed: completed, Userid: todo.Userid}
	c.JSON(http.StatusOK, gin.H{"status": http.StatusOK, "data": _todo})
}

// UpdateTodo update a todo
func UpdateTodo(c *gin.Context) {
	db := c.MustGet("db").(*gorm.DB)
	var todo models.TodoModel
	todoID := c.Param("id")
	db.First(&todo, todoID)
	if todo.ID == 0 {
		c.JSON(http.StatusNotFound, gin.H{"status": http.StatusNotFound, "message": "No todo found!"})
		return
	}
	db.Model(&todo).Update("title", c.PostForm("title"))
	completed, _ := strconv.Atoi(c.PostForm("completed"))
	db.Model(&todo).Update("completed", completed)
	c.JSON(http.StatusOK, gin.H{"status": http.StatusOK, "message": "Todo updated successfully!"})
}

// DeleteTodo remove a todo
func DeleteTodo(c *gin.Context) {
	db := c.MustGet("db").(*gorm.DB)
	var todo models.TodoModel
	todoID := c.Param("id")
	db.First(&todo, todoID)
	if todo.ID == 0 {
		c.JSON(http.StatusNotFound, gin.H{"status": http.StatusNotFound, "message": "No todo found!"})
		return
	}
	db.Delete(&todo)
	c.JSON(http.StatusOK, gin.H{"status": http.StatusOK, "message": "Todo deleted successfully!"})
}
