package main

import (
	"fmt"

	"github.com/gin-gonic/gin"
)

func main() {
	gin.SetMode(gin.ReleaseMode)
	r := gin.Default()
	r.GET("", func(c *gin.Context) {
		type User struct {
			Name string `form:"name"`
			Age  int    `form:"age"`
		}
		var user User
		err := c.ShouldBindQuery(&user)
		if err != nil {
			fmt.Println(err)
			return
		}
		fmt.Println(user)
	})
	r.GET("/users/:id/:name", func(c *gin.Context) {
		type User struct {
			Name string `uri:"name"`
			ID   int    `uri:"id"`
		}
		var user User
		err := c.ShouldBindUri(&user)
		fmt.Println(user, err)
	})
	r.POST("/form", func(c *gin.Context) {
		type User struct {
			Name string `form:"name"`
			Age  int    `form:"age"`
		}
		var user User
		err := c.ShouldBind(&user)
		fmt.Println(user, err)
	})
	r.POST("/userjson", func(c *gin.Context) {
		type Address struct {
			Ave      string `json:"ave"`
			Building string `json:"building"`
		}
		type User struct {
			Name    string    `json:"name"`
			Age     int       `json:"age"`
			Address []Address `json:"address" binding:"dive,required"`
		}
		var user User
		err := c.ShouldBindJSON(&user)
		fmt.Println(user, err)
	})
	r.Run(":8080")
}
