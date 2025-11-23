package global

import (
	"app.local/app/routes"
	"github.com/gin-gonic/gin"
)

var Router *gin.Engine

func SetApiRoutes() {
	// API路由组
	api := Router.Group("/api")
	{
		api.GET("/data", routes.GetData)
		api.POST("/data", routes.GetData)
	}
}
