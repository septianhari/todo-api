package main

import (
	"github.com/gin-gonic/gin"
	"github.com/septianhari/todo-api/controllers"
	_ "github.com/septianhari/todo-api/docs"
	"github.com/septianhari/todo-api/models"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

var DB *gorm.DB
var err error

func main() {
	// Initialize database
	DB, err = gorm.Open(sqlite.Open("todo.db"), &gorm.Config{})
	if err != nil {
		panic("failed to connect to the database")
	}

	DB.AutoMigrate(&models.Todo{})

	// Initialize Gin
	r := gin.Default()

	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	// Define routes
	api := r.Group("/api/v1")
	{
		api.POST("/todos", controllers.CreateTodo)
		api.GET("/todos", controllers.GetTodos)
		api.GET("/todos/:id", controllers.GetTodoByID)
		api.PUT("/todos/:id", controllers.UpdateTodo)
		api.DELETE("/todos/:id", controllers.DeleteTodo)
	}

	// Run the server
	r.Run(":8080")
}
