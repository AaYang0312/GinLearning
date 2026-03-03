package main

import (
	"bytes"
	"fmt"
	"io"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()

	r.POST("", func(context *gin.Context) {
		byteData, _ := io.ReadAll(context.Request.Body)
		fmt.Println(string(byteData)) // 这样读body阅后即焚
		// 阅后可以再次阅读的写法
		context.Request.Body = io.NopCloser(bytes.NewReader(byteData))
		//name, _ := context.GetPostForm("name")
		//fmt.Println(name)
	})
	r.Run(":8080")
}
