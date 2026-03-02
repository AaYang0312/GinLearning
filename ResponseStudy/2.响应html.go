package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	gin.SetMode(gin.ReleaseMode)
	r := gin.Default()
	r.LoadHTMLGlob("ResponseStudy/static/*") // 重要步骤
	//r.LoadHTMLFiles()                                // 需要写具体文件名
	r.GET("", func(c *gin.Context) {
		c.HTML(http.StatusOK, "浏览器下载.html", gin.H{ // 不常用，需要结合前端，但可以直接修改标题
			"title": "百度知道",
		})
	})
	r.Run(":8080")
}
