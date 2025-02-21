package docs

import (
	"github.com/gin-gonic/gin"
	_ "github.com/septianhari/todo-api/docs"
	ginSwagger "github.com/swaggo/gin-swagger"
	"github.com/swaggo/gin-swagger/swaggerFiles"
)

// @title Todo API
// @version 1.0
// @description API for managing to-do list posts on social media

// @contact.name API Support
// @contact.email support@yourdomain.com

// @BasePath /api/v1
func InitSwagger(router *gin.Engine) {
	swagger := gin.New()
	swagger.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	router.Use(swagger)
}
