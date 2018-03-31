package model

import (
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"time"
	"bipu-backend/src/main/middleware"
	"fmt"
)

var (
	dbConf    middleware.Db
	dbConnect string
)

func InitDB() {
	dbConf = middleware.Config.Db
	var dbInfo = "%s:%s@tcp(%s:%s)/%s?charset=utf8&parseTime=True"
	dbConnect = fmt.Sprintf(dbInfo, dbConf.User, dbConf.Pass, dbConf.Host, dbConf.Port, dbConf.Database)
	fmt.Println("dbConnect: " + dbConnect)
}

type User struct {
	Uid       int       `form:"uid" json:"uid" gorm:"primary_key;AUTO_INCREMENT"`
	Login     string    `form:"login" json:"login"`
	Password  string    `form:"password" json:"password"`
	AvatarUrl string    `form:"avatar_url" json:"avatar_url"`
	Name      string    `form:"name" json:"name"`
	Bio       string    `form:"bio" json:"bio"`
	Email     string    `form:"email" json:"email"`
	Status    int       `form:"status" json:"status"`
	Sex       int8      `form:"sex" json:"sex"`
	Openid    int64     `form:"openid" json:"openid" gorm:"not null" `
	OpenToken string    `form:"open_token" json:"open_token" gorm:"not null" `
	Source    string    `form:"source" json:"source" gorm:"not null" `
	CreateAt  time.Time `form:"create_at" json:"create_at"`
	UpdateAt  time.Time `form:"update_at" json:"update_at"`
}

func (User) TableName() string {
	return "user"
}

func InsertUserInfo(user *User) (err error) {

	db, err := gorm.Open("mysql", dbConnect)
	if err != nil {
		panic("failed to connect database")
	}
	db.LogMode(true)

	if db.NewRecord(user) {
		db.Create(user)
	} else {
		db.Model(user).Where("uid= ?", user.Uid).Update("open_token", user.OpenToken)
	}
	defer db.Close()

	return nil
}

func QueryUserInfo(userId int) (User) {
	var err error
	db, err := gorm.Open("mysql", dbConnect)
	if err != nil {
		panic("failed to connect database")
	}
	//defer db.Close()
	db.LogMode(true)
	defer db.Close()

	var user User
	db.Where("uid = ?", userId).First(&user)
	return user
}
