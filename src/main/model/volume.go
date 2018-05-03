package model

import (
	"bipu-backend/src/main/util"
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

func dbColumn(structFieldName string) string {
	return util.GetStructTagName4Gorm(&Volume{}, structFieldName, TAB_FIELD_KEY)
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
	volumes := make([]Volume, 0)
	db := CommomGetDB()
	db.Where(dbColumn("Uid")+" = ?", uid).Find(&volumes)

	return volumes
}

// 查询一个用户创建的所有谱册
func QueryVolumesByVid(vid string) []Volume {
	volumes := make([]Volume, 0)
	db := CommomGetDB()
	db.Where(dbColumn("Vid")+" = ?", vid).Find(&volumes)

	return volumes
}
