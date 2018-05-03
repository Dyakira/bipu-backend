package controller

import "github.com/gin-gonic/gin"

const (
	RET_STATUS = "status"
	RET_MSG    = "msg"
	RET_DATA   = "data"
	RET_OTHERS = "others"
)

// 构造通用的返回信息结构
func ReturnStruct(status int, msg string, data interface{}, others interface{}) gin.H {
	return gin.H{RET_STATUS: status, RET_MSG: msg, RET_DATA: data, RET_OTHERS: others}
}

// 正常返回；返回码为0，无信息，无附加数据
func OkNormal(c *gin.Context, data interface{}) {
	c.JSON(200, ReturnStruct(0, "", data, ""))
}
