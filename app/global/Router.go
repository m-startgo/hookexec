package global

import (
	"embed"
	"net/http"
	"path/filepath"

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

// 初始化静态文件服务
func SetStaticServer(distDir embed.FS) {
	// 首先处理根路径的index.html
	Router.GET("/", func(c *gin.Context) {
		data, err := distDir.ReadFile("frontend/dist/index.html")
		if err != nil {
			c.String(http.StatusInternalServerError, "无法读取index.html")
			return
		}
		c.Data(http.StatusOK, "text/html", data)
	})

	// 处理所有其他非API路径，用于SPA路由支持
	Router.NoRoute(func(c *gin.Context) {
		// 获取请求路径
		path := c.Request.URL.Path
		// 排除API路径
		if len(path) > 4 && path[:4] == "/api" {
			c.String(http.StatusNotFound, "API路径不存在")
			return
		}
		// 尝试直接提供文件
		if path != "/" {
			data, err := distDir.ReadFile("frontend/dist" + path)
			if err == nil {
				// 设置适当的Content-Type
				ext := filepath.Ext(path)
				switch ext {
				case ".js":
					c.Data(http.StatusOK, "application/javascript", data)
				case ".css":
					c.Data(http.StatusOK, "text/css", data)
				case ".html":
					c.Data(http.StatusOK, "text/html", data)
				default:
					c.Data(http.StatusOK, "application/octet-stream", data)
				}
				return
			}
		}
		// 检查是否有404.html
		if path != "/404.html" {
			data, err404 := distDir.ReadFile("frontend/dist/404.html")
			if err404 == nil {
				c.Data(http.StatusNotFound, "text/html", data)
				return
			}
		}

		// 最后返回index.html，用于SPA路由
		data, indexErr := distDir.ReadFile("frontend/dist/index.html")
		if indexErr == nil {
			c.Data(http.StatusOK, "text/html", data)
		} else {
			c.String(http.StatusNotFound, "Not Found")
		}
	})
}
