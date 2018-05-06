package model

import (
	"bipu-backend/src/main/middleware"
	"fmt"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

var (
	commonDbConf    middleware.Db
	commonDbConnect string
)

// 初始化数据库连接字符串
func CommomInitDB() {
	commonDbConf = middleware.Config.Db
	var dbInfo = "%s:%s@tcp(%s:%s)/%s?charset=utf8&parseTime=True"
	commonDbConnect = fmt.Sprintf(dbInfo, commonDbConf.User, commonDbConf.Pass, commonDbConf.Host, commonDbConf.Port, commonDbConf.Database)
	fmt.Println("[model][common]dbConnect: " + commonDbConnect)
}

// 获取db的通用方法
func CommomGetDB() (db *gorm.DB) {
	db, err := gorm.Open(DB_TYPE, commonDbConnect)
	if err != nil {
		panic("[model][common]failed to connect database")
	}
	db.LogMode(true)
	return db
}
