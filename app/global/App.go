package global

/*
和 wails App 全局状态相关的函数

*/

import (
	"app.local/app/utils/flog"
)

func SysInit() {
	flog.LogInit() // 初始化日志系统
}
