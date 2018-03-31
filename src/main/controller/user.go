package controller

import (
	"net/http"
	"github.com/gin-gonic/contrib/sessions"
	"bipu-backend/src/main/model"
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

