package routes

import (
	"github.com/Milagrosgzmn/devops_todo_go.git/internal/handlers"
	"github.com/Milagrosgzmn/devops_todo_go.git/internal/repository"
	"github.com/gin-gonic/gin"
	"github.com/newrelic/go-agent/v3/integrations/nrgin"
	"github.com/newrelic/go-agent/v3/newrelic"
)

func SetupRouter(repo repository.IRepository, nrApp *newrelic.Application) *gin.Engine {
	router := gin.Default()

	// si esta configurado New Relic, usamos el middleware
	if nrApp != nil {
		router.Use(nrgin.Middleware(nrApp))
	}

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