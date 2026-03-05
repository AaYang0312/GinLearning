// /3.中间件.go
package main

import (
	"fmt"

	"github.com/gin-gonic/gin"
)

func Home(c *gin.Context) {
	fmt.Println("Home")
	c.String(200, "Home")
}
func M1(c *gin.Context) {
	fmt.Println("M1 请求部分")
	c.Abort()
	fmt.Println("M1 响应部分")
}
func M2(c *gin.Context) {
	fmt.Println("M2 请求部分")
	c.Next()
	fmt.Println("M2 响应部分")
}

func GM1(c *gin.Context) {
	fmt.Println("GM1 请求部分")
	c.Abort()               // 阻止后续中间件和处理器执行
	fmt.Println("GM1 响应部分") // 响应部分按注册顺序逆序执行
}

func GM2(c *gin.Context) {
	fmt.Println("GM2 请求部分")
	c.Next()
	fmt.Println("GM2 响应部分")
}

func AuthMiddleware(c *gin.Context) {

}

func main() {
	r := gin.Default()
	g := r.Group("api")
	g.Use(GM1, GM2)
	g.GET("users", Home)
	r.Run(":8080")
}
