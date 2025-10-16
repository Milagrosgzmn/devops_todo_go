package handlers

import (
	"github.com/gin-gonic/gin"
)

type HealthHandler struct {
}

func NewHealthHandler() *HealthHandler {
	return &HealthHandler{}
}


func (h *HealthHandler) HealthCheck(c *gin.Context) {

	// Podemos agregar lógica adicional aquí si es necesario

	// Responder con un estado 200 OK - el servidor está saludable
	c.JSON(200, gin.H{
		"status": "OK",
	})
}