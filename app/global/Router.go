package global

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

var Router *gin.Engine

func SetApiRoutes() {
	// API路由组
	api := Router.Group("/api")
	{
		// 示例：返回一些JSON数据
		api.GET("/data", func(c *gin.Context) {
			c.JSON(http.StatusOK, gin.H{
				"message": "Hello from API",
				"data": gin.H{
					"items":     []string{"item1", "item2", "item3"},
					"status":    "active",
					"timestamp": "2024-01-01T00:00:00Z",
				},
			})
		})

		// 健康检查接口
		api.GET("/health", func(c *gin.Context) {
			c.JSON(http.StatusOK, gin.H{
				"status":  "healthy",
				"service": "hookexec",
			})
		})
	}
}
