package main

import (
	"fmt"

	"github.com/gin-gonic/gin"
)

func UserView(c *gin.Context) {
	path := c.Request.URL
	fmt.Println(path)
	c.String(200, "Hello Path")
}
func main() {
	gin.SetMode(gin.ReleaseMode)
	r := gin.Default()

	//r.POST()   // 提交数据
	//r.GET()    // 获取数据
	//r.PUT()    // 全量更新
	//r.PATCH()  // 选择更新
	//r.DELETE() // 删除数据
	//
	//r.Any() // 支持所有

	apiGroup := r.Group("api") // 创建一个分组
	UserGroup(apiGroup)        // /api 下辖
	r.Run(":8080")
}
func UserGroup(group *gin.RouterGroup) {
	group.GET("users", UserView)
	group.POST("users", UserView)
	// ...

}
