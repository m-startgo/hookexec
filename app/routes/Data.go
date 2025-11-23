package routes

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/m-startgo/go-utils/mtime"
)

func GetData(c *gin.Context) {
	time := mtime.NowDefaultString()
	// 获取全部的查询参数（ShouldBindQuery 无法把查询参数绑定到 map[string]any）
	params := map[string]any{}
	for k, v := range c.Request.URL.Query() {
		if len(v) == 1 {
			params[k] = v[0]
		} else {
			params[k] = v
		}
	}

	var data map[string]any
	c.ShouldBindJSON(&data)

	c.JSON(http.StatusOK, gin.H{
		"message": "Hello from API",
		"data": gin.H{
			"timestamp": time,
			"method":    c.Request.Method,
			"params":    params,
			"data":      data,
		},
	})
}
