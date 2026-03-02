package res

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type Response struct {
	Code int    `json:"code"`
	Msg  string `json:"msg"`
	Data any    `json:"data"`
}

var codeMap = map[int]string{
	1001: "权限错误",
	1002: "角色错误",
}

func response(c *gin.Context, code int, data any, msg string) {
	c.JSON(http.StatusOK, Response{
		Code: code,
		Msg:  msg,
		Data: data,
	})
}
func OK(c *gin.Context, data any, msg string) {
	response(c, 0, data, msg)
}
func OKWithData(c *gin.Context, data any) {
	OK(c, data, "成功")
}
func OKWithMsg(c *gin.Context, msg string) {
	OK(c, gin.H{}, msg)
}
func Fail(c *gin.Context, code int, data any, msg string) {
	response(c, code, data, msg)
}
func FailWithMsg(c *gin.Context, msg string) {
	response(c, 1001, gin.H{}, msg) // code 设一个固定值
}
func FailWithCode(c *gin.Context, code int) {
	msg, OK := codeMap[code]
	if !OK {
		msg = "服务错误"
	}
	response(c, code, gin.H{}, msg)
}
