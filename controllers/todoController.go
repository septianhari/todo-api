package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/septianhari/todo-api/models"
	"gorm.io/gorm"
)

var DB *gorm.DB

func CreateTodo(c *gin.Context) {
	var todo models.Todo
	if err := c.ShouldBindJSON(&todo); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := DB.Create(&todo).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Todo created", "data": todo})
}

func GetTodos(c *gin.Context) {
	var todos []models.Todo
	if err := DB.Find(&todos).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": todos})
}

func GetTodoByID(c *gin.Context) {
	id := c.Param("id")
	var todo models.Todo
	if err := DB.First(&todo, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Todo not found"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": todo})
}

func UpdateTodo(c *gin.Context) {
	id := c.Param("id")
	var todo models.Todo
	if err := DB.First(&todo, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Todo not found"})
		return
	}

	if err := c.ShouldBindJSON(&todo); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := DB.Save(&todo).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Todo updated", "data": todo})
}

func DeleteTodo(c *gin.Context) {
	id := c.Param("id")
	if err := DB.Delete(&models.Todo{}, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Todo not found"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Todo deleted"})
}
