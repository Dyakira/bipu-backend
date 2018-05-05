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

// 带信息返回；返回码为0，无附加数据
func OkMsgNormal(c *gin.Context, msg string, data interface{}) {
	c.JSON(200, ReturnStruct(0, msg, data, ""))
}

// 正常返回；返回码为0，无信息，无附加数据
func OkNormal(c *gin.Context, data interface{}) {
	OkMsgNormal(c, "", data)
}

// 正常错误；返回码为-1，无信息，无附加数据
func ErrParamNormal(c *gin.Context, msg string) {
	c.JSON(400, ReturnStruct(-1, msg, "", ""))
}
