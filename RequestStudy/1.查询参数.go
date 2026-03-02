package main

import (
	"fmt"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	r.GET("", func(c *gin.Context) {
		name := c.Query("name")
		age := c.DefaultQuery("age", "25")
		keyList := c.QueryArray("key")
		fmt.Println(name, age, keyList) // fengfeng 123 [123 124]
	})
	r.Run(":8080")
}
