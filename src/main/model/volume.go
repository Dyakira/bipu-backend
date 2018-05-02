package model

import (
	"time"
)

// 谱册表
type Volume struct {
	Vid        string    `json:"vid" gorm:"column:f_vid"`
	Name       string    `json:"name" gorm:"column:f_name"`
	Cover      string    `json:"cover" gorm:"column:f_cover"`
	Desc       string    `json:"desc" gorm:"column:f_desc"`
	Uid        string    `json:"uid" gorm:"column:f_uid"`
	Status     int       `json:"status" gorm:"column:f_status"`
	CreateTime time.Time `json:"createTime" gorm:"column:f_create_time"`
}

func (Volume) TableName() string {
	return "tab_volume"
}

// 查询所有谱册
func QueryVolumesAll() []Volume {
	volumes := make([]Volume, 0)
	db := CommomGetDB()
	db.Find(&volumes)

	return volumes
}

// 查询一个用户创建的所有谱册
func QueryVolumesByUid(uid string) []Volume {
	// TODO 待实现
	return nil
}
