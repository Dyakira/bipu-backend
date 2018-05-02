package model

import (
	"time"
)

// 谱册和谱子的多对多关系表
type VolumeSheet struct {
	Vsid       string    `json:"vsid"  gorm:"column:f_vsid"`
	Vid        string    `json:"vid"  gorm:"column:f_vid"`
	Sid        string    `json:"sid"  gorm:"column:f_sid"`
	CreateTime time.Time `json:"createTime"  gorm:"column:f_create_time"`
}

func (VolumeSheet) TableName() string {
	return "tab_volume_sheet"
}
