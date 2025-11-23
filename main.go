package main

import (
	"embed"
	"io"

	"app.local/app/global"
	"app.local/app/utils/flog"
	"github.com/gin-gonic/gin"
	"github.com/m-startgo/go-utils/mstr"
)

//go:embed all:frontend/dist
var FrontendDist embed.FS

func main() {
	global.SysInit() // 初始化全局系统设置

	// 设置Gin模式
	gin.SetMode(gin.ReleaseMode)
	gin.DefaultWriter = io.Discard // 丢弃标准输出日志
	// gin.DefaultErrorWriter = io.Discard // 丢弃错误输出日志

	global.Router = gin.Default()

	global.SetApiRoutes()

	// 初始化静态文件服务（后定义通配符路由）
	global.SetStaticServer(FrontendDist)

	// 启动服务器
	port := "9900"
	flog.App.Info("服务器启动", mstr.Join("http://localhost:", port))

	err := global.Router.Run(":" + port)
	if err != nil {
		flog.App.Error("服务器启动失败:", err.Error())
	}
}
