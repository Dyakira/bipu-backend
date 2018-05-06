package model

import (
	"bipu-backend/src/main/util"
	"time"
)

// 谱册表
type Volume struct {
	Vid        int       `json:"vid" gorm:"column:f_vid;primary_key;AUTO_INCREMENT"`
	Name       string    `json:"name" gorm:"column:f_name"`
	Cover      string    `json:"cover" gorm:"column:f_cover"`
	Desc       string    `json:"desc" gorm:"column:f_desc"`
	Uid        int       `json:"uid" gorm:"column:f_uid"`
	Status     int       `json:"status" gorm:"column:f_status"`
	CreateTime time.Time `json:"createTime" gorm:"column:f_create_time"`
}

func (Volume) TableName() string {
	return "tab_volume"
}

func dbColumn(structFieldName string) string {
	return util.GetStructTagName4Gorm(&Volume{}, structFieldName, TAB_FIELD_KEY)
}

// 新增或更新一个谱册，更新只能更新部分字段
func InsertOrUpdateVolume(volume *Volume) {
	db := CommomGetDB()
	defer db.Close()

	if db.NewRecord(*volume) {
		// 新记录
		db.Create(volume)
	} else {
		// 已经存在这条记录了
		volumeInDB := QueryVolumesByVid(volume.Vid)
		// 更新部分字段
		volumeInDB.Name = volume.Name
		volumeInDB.Desc = volume.Desc
		volumeInDB.Cover = volume.Cover
		// 更新
		db.Save(&volumeInDB)
	}
}

// 查询所有谱册
func QueryVolumesAll() []Volume {
	volumes := make([]Volume, 0)
	db := CommomGetDB()
	defer db.Close()
	db.Find(&volumes)

	return volumes
}

// 查询一个用户创建的所有谱册
func QueryVolumesByUid(uid int) []Volume {
	volumes := make([]Volume, 0)
	db := CommomGetDB()
	defer db.Close()
	db.Where(dbColumn("Uid")+" = ?", uid).Find(&volumes)

	return volumes
}

// 通过id查询一个谱册
func QueryVolumesByVid(vid int) Volume {
	var volume Volume
	db := CommomGetDB()
	defer db.Close()
	db.Where(dbColumn("Vid")+" = ?", vid).First(&volume)

	return volume
}
