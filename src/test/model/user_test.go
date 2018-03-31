package model

import "bipu-backend/src/main/model"
import (
	"testing"
)

func TestUser_InsertUserInfo(t *testing.T) {
	user := model.User{Openid: 1231445, Source: "abc"}
	model.InsertUserInfo(&user)
}


func TestUser_QueryUserInfo(t *testing.T) {
	user := model.QueryUserInfo(123123)
	println(user.Openid)
}
