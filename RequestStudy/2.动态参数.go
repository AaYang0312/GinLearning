package main

import (
	"fmt"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	r.GET("users/:id/:name", func(c *gin.Context) {
		id := c.Param("id")
		name := c.Param("name")
		fmt.Println(id, "/", name)
	})
	r.Run(":8080")
}
