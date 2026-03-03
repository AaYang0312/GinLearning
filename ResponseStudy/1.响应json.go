package main

import (
	"ResponseStudy/res"

	"github.com/gin-gonic/gin"
)

//	func Index(c *gin.Context) {
//		c.JSON(http.StatusOK,
//			gin.H{"code": 0, "msg": "Success", "data": gin.H{}})
//	}
func main() {
	r := gin.Default()

	r.GET("/index", func(c *gin.Context) {
		res.OKWithMsg(c, "登陆成功")
	})
	r.GET("/users", func(c *gin.Context) {
		res.OKWithData(c, map[string]any{
			"username":   "Louis",
			"ip_address": "127.0.0.1",
		})
	})
	r.POST("/users", func(c *gin.Context) {
		res.FailWithMsg(c, "参数错误")
	})
	r.GET("/login", func(c *gin.Context) {
		res.FailWithMsg(c, "登陆失败")
	})

	r.Run(":8080")
}
