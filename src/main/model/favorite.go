package model

import (
	"time"
)

// 用户和谱册的多对多关系表
type Favorite struct {
	Fid        int       `json:"fid"  gorm:"column:f_fid"`
	Uid        int       `json:"uid"  gorm:"column:f_uid"`
	Vid        int       `json:"vid"  gorm:"column:f_vid"`
	CreateTime time.Time `json:"createTime"  gorm:"column:f_create_time"`
}

func (Favorite) TableName() string {
	return "tab_favorite"
}
