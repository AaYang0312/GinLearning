package main

import (
	"fmt"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()

	r.POST("/files", func(c *gin.Context) {
		forms, err := c.MultipartForm()
		if err != nil {
			fmt.Println(err)
			return
		}
		for _, headers := range forms.File {
			for _, header := range headers {
				c.SaveUploadedFile(header, "upload/"+header.Filename, 0666)
			}
		}
	})
	r.Run(":8080")
}
