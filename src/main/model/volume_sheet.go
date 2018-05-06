package model

import (
	"time"
)

// 谱册和谱子的多对多关系表
type VolumeSheet struct {
	Vsid       int       `json:"vsid"  gorm:"column:f_vsid"`
	Vid        int       `json:"vid"  gorm:"column:f_vid"`
	Sid        int       `json:"sid"  gorm:"column:f_sid"`
	CreateTime time.Time `json:"createTime"  gorm:"column:f_create_time"`
}

func (VolumeSheet) TableName() string {
	return "tab_volume_sheet"
}
