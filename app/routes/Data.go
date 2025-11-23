package routes

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func Data(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"message": "Hello from API",
		"data": gin.H{
			"items":     []string{"item1", "item2", "item3"},
			"status":    "active",
			"timestamp": "2024-01-01T00:00:00Z",
		},
	})
}
