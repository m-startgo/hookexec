package routes

import (
	"encoding/json"
	"io"
	"net/http"

	"app.local/app/utils/flog"
	"github.com/gin-gonic/gin"
)

// GitHubEvent 表示基本的GitHub webhook事件结构
type GitHubEvent struct {
	Action       string          `json:"action,omitempty"`
	Repository   json.RawMessage `json:"repository,omitempty"`
	Sender       json.RawMessage `json:"sender,omitempty"`
	Installation json.RawMessage `json:"installation,omitempty"`
}

// GitHubWebhookHandler 处理GitHub webhook事件
func GitHubWebhookHandler(c *gin.Context) {
	// 读取请求体
	body, err := io.ReadAll(c.Request.Body)
	if err != nil {
		flog.App.Error("GitHubWebhookHandler", "读取请求体失败", err)
		c.JSON(http.StatusBadRequest, gin.H{"error": "无法读取请求体"})
		return
	}

	// 获取请求头信息
	event := c.GetHeader("X-GitHub-Event")

	flog.App.Info("GitHubWebhookHandler", "收到GitHub webhook事件", "event", event, body)

	// 返回成功响应
	c.JSON(http.StatusOK, gin.H{})
}
