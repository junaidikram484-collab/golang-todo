package main

import (
	"os"

	handlers "github.com/junaid/todo-service/internal/handler"
	"github.com/junaid/todo-service/internal/repository"

	_ "github.com/junaid/todo-service/docs"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"

	"github.com/gin-gonic/gin"
)

func main() {
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	repo := repository.NewMemoryRepository()
	handler := handlers.NewTodoHandler(repo)

	router := gin.Default()
	// Swagger endpoint
	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	v1 := router.Group("/api/v1")
	{
		v1.GET("/todos", handler.GetAll)
		v1.POST("/todos", handler.Create)
		v1.PUT("/todos/:id", handler.Update)
		v1.DELETE("/todos/:id", handler.Delete)
		v1.GET("/health", handler.Health)
	}

	handler.RegisterRoutes(router)

	//router.Run(":" + port)
	router.Run("0.0.0.0:" + port)
}
