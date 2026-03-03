package main

import (
	"fmt"
	"io"
	"os"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()

	r.POST("/users", func(c *gin.Context) {
		name := c.PostForm("name")
		age, ok := c.GetPostForm("age")
		fmt.Println(name, age, ok)
	})
	r.POST("/files", func(c *gin.Context) {
		Header, err := c.FormFile("头像")
		if err != nil {
			fmt.Println(err)
			return
		}
		File, err := Header.Open()
		byteData, err := io.ReadAll(File)
		if err != nil {
			fmt.Println(err)
			return
		}
		err = os.WriteFile("xxx.jpg", byteData, 0666)
		if err != nil {
			fmt.Println(err)
			return
		}
	})
	r.POST("/filesEasy", func(c *gin.Context) {
		header, err := c.FormFile("头像")
		if err != nil {
			fmt.Println(err)
			return
		}

		err = c.SaveUploadedFile(header, "upload/"+header.Filename, 0666)
		if err != nil {
			fmt.Println(err)
			return
		}

	})

	r.Run(":8080")
}
