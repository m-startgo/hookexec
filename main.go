package main

import (
	"embed"
	"log"
	"net/http"
	"path/filepath"

	"app.local/app/global"
	"app.local/app/utils/flog"
	"github.com/gin-gonic/gin"
	"github.com/m-startgo/go-utils/mstr"
)

//go:embed all:frontend/dist
var FrontendDist embed.FS

// 初始化静态文件服务
func initStaticFiles() {
	// 首先处理根路径的index.html
	global.Router.GET("/", func(c *gin.Context) {
		data, err := FrontendDist.ReadFile("frontend/dist/index.html")
		if err != nil {
			c.String(http.StatusInternalServerError, "无法读取index.html")
			return
		}
		c.Data(http.StatusOK, "text/html", data)
	})

	// 处理所有其他非API路径，用于SPA路由支持
	global.Router.NoRoute(func(c *gin.Context) {
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

func main() {
	global.SysInit() // 初始化全局系统设置

	// 设置Gin模式
	gin.SetMode(gin.ReleaseMode)
	global.Router = gin.Default()

	global.SetApiRoutes()

	// 初始化静态文件服务（后定义通配符路由）
	initStaticFiles()

	// 启动服务器
	port := "9900"
	flog.App.Info("服务器启动", mstr.Join("http://localhost:", port))

	err := global.Router.Run(":" + port)
	if err != nil {
		log.Fatalf("启动服务器失败: %v", err)
	}
}
