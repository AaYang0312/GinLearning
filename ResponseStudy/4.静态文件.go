package main

import "github.com/gin-gonic/gin"

func main() {
	r := gin.Default()
	r.Static("st", "static")              // 第一个参数是别名，第二个是文件目录
	r.StaticFile("abc", "static/abc.txt") // 直接挂载文件
	r.Run(":8080")
}
