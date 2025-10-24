package routes

import (
	"github.com/Milagrosgzmn/devops_todo_go.git/internal/handlers"
	"github.com/Milagrosgzmn/devops_todo_go.git/internal/repository"
	"github.com/gin-gonic/gin"
)

func SetupRouter(repo repository.IRepository) *gin.Engine {
	router := gin.Default()

	itemHandler := handlers.NewItemHandler(repo)

	router.GET("/health", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"status": "OK",
		})
	});

	router.GET("/items", itemHandler.GetItems)
	router.GET("/items/:id", itemHandler.GetItemByID)
	router.POST("/items", itemHandler.CreateItem)
	router.PUT("/items/:id", itemHandler.UpdateItem)
	router.DELETE("/items/:id", itemHandler.DeleteItem)

	return router
}