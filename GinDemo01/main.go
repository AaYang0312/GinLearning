package main

import (
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

type Response struct {
	Code int            `json:"code"`
	Msg  string         `json:"msg"`
	Data map[string]any `json:"data"`
}

func Ping(c *gin.Context) {
	// 返回 JSON 响应
	c.JSON(http.StatusOK, gin.H{
		"message": "pong",
	})
}
func Index(c *gin.Context) {
	response := Response{
		Code: 200,
		Msg:  "success",
		Data: map[string]any{
			"timestamp": time.Now().Unix(),
		},
	}
	// 返回 JSON 响应
	c.JSON(http.StatusOK, response)
}
func main() {
	// 创建带默认中间件（日志与恢复）的 Gin 路由器
	r := gin.Default()

	// 定义简单的 GET 路由
	r.GET("/ping", Ping)

	r.POST("/index", Index)

	// 默认端口 8080 启动服务器
	// 监听 0.0.0.0:8080（Windows 下为 localhost:8080）
	r.Run(":8080")
}
