package main

import (
	"embed"
	"log"
	"net/http"
	"path/filepath"

	"github.com/gin-gonic/gin"
)

//go:embed all:frontend/dist
var FrontendDist embed.FS

func main() {
	// 设置Gin模式
	gin.SetMode(gin.ReleaseMode)

	// 创建Gin引擎
	router := gin.Default()

	// 初始化API路由（先定义具体路由）
	initApiRoutes(router)

	// 初始化静态文件服务（后定义通配符路由）
	initStaticFiles(router)

	// 启动服务器
	port := "9900"
	log.Printf("服务器启动在 http://localhost:%s", port)
	if err := router.Run(":" + port); err != nil {
		log.Fatalf("启动服务器失败: %v", err)
	}
}

// 初始化静态文件服务
func initStaticFiles(router *gin.Engine) {
	// 首先处理根路径的index.html
	router.GET("/", func(c *gin.Context) {
		data, err := FrontendDist.ReadFile("frontend/dist/index.html")
		if err != nil {
			c.String(http.StatusInternalServerError, "无法读取index.html")
			return
		}
		c.Data(http.StatusOK, "text/html", data)
	})

	// 处理所有其他非API路径，用于SPA路由支持
	router.NoRoute(func(c *gin.Context) {
		// 获取请求路径
		path := c.Request.URL.Path

		// 排除API路径
		if len(path) > 4 && path[:4] == "/api" {
			c.String(http.StatusNotFound, "API路径不存在")
			return
		}

		// 尝试直接提供文件
		if path != "/" {
			data, err := FrontendDist.ReadFile("frontend/dist" + path)
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
			data, err404 := FrontendDist.ReadFile("frontend/dist/404.html")
			if err404 == nil {
				c.Data(http.StatusNotFound, "text/html", data)
				return
			}
		}

		// 最后返回index.html，用于SPA路由
		data, indexErr := FrontendDist.ReadFile("frontend/dist/index.html")
		if indexErr == nil {
			c.Data(http.StatusOK, "text/html", data)
		} else {
			c.String(http.StatusNotFound, "Not Found")
		}
	})
}

// 初始化API路由
func initApiRoutes(router *gin.Engine) {
	// API路由组
	api := router.Group("/api")
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
