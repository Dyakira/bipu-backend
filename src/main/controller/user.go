package controller

import (
	"bipu-backend/src/main/model"
	"bipu-backend/src/main/util"
	"net/http"

	"github.com/gin-gonic/contrib/sessions"
	"github.com/gin-gonic/gin"
)

func UserLogoutHandler(c *gin.Context) {
	session := sessions.Default(c)
	session.Delete("uid")
	session.Save()
	c.JSON(http.StatusOK, gin.H{
		"status":  "success",
		"message": "logout",
	})
}

func UserLoginHandler(c *gin.Context) {
	session := sessions.Default(c)
	session.Delete("uid")
	session.Save()
	message := "uid is logout"
	c.JSON(http.StatusOK, gin.H{
		"status":  "success",
		"message": message,
	})
}

func UserInfo(c *gin.Context) {
	session := sessions.Default(c)
	uid := session.Get("uid").(int)
	user := model.QueryUserInfo(uid)
	c.JSON(http.StatusOK, gin.H{
		"status":  "success",
		"message": user,
	})
}

// 一般的用户注册
func InsertUser(c *gin.Context) {
	session := sessions.Default(c)
	var user model.User

	c.Bind(&user)
	// TODO 参数校验

	// 1 用户名不能重复
	userInDB := model.QueryUserByLogin(user.Login)
	if 0 != userInDB.Uid {
		ErrParamNormal(c, "用户名重复")
		return
	}
	// 2 写数据库
	// 转化密码
	user.Password = util.Md5StringWithSalt(user.Password)
	model.InsertUserInfo(&user)

	// 3 用户存入session
	session.Set("uid", user.Uid)
	session.Save()

	// 4 返回user
	// 去除密码
	user.Password = ""
	OkNormal(c, user)
}

// 一般的用户登陆
func UserLogin(c *gin.Context) {
	session := sessions.Default(c)
	var user model.User

	c.Bind(&user)

	userInDB := model.QueryUserByLogin(user.Login)

	if 0 == userInDB.Uid {
		ErrParamNormal(c, "用户不存在")
		return
	}
	if userInDB.Password != util.Md5StringWithSalt(user.Password) {
		ErrParamNormal(c, "密码错误")
		return
	}

	// 用户存入session
	session.Set("uid", user.Uid)
	session.Save()

	// 去除密码
	userInDB.Password = ""
	//返回user
	OkMsgNormal(c, "登陆成功", userInDB)
}
