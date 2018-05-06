package controller

import (
	"bipu-backend/src/main/model"
	"fmt"
	"strconv"
	"time"

	"github.com/gin-gonic/contrib/sessions"
	"github.com/gin-gonic/gin"
)

func QueryVolumesAll(c *gin.Context) {
	volumes := model.QueryVolumesAll()
	OkNormal(c, volumes)
}

func QueryVolumesByUid(c *gin.Context) {
	uid, err := strconv.Atoi(c.Params.ByName("uid"))
	if nil != err {
		ErrInternalNormal(c)
		return
	}
	volumes := model.QueryVolumesByUid(uid)

	OkNormal(c, volumes)
}

func QueryVolumesByVid(c *gin.Context) {
	vid, err := strconv.Atoi(c.Params.ByName("vid"))
	if nil != err {
		ErrInternalNormal(c)
		return
	}
	volume := model.QueryVolumesByVid(vid)

	OkNormal(c, volume)
}

// 创建一个谱册
func InsertVolume(c *gin.Context) {
	session := sessions.Default(c)
	fmt.Println(c.Keys)
	uid, ok := session.Get("uid").(int)
	if !ok {
		ErrInternalNormal(c)
		return
	}
	if 0 == uid {
		ErrParamNormal(c, "未登录")
		return
	}
	var volume model.Volume
	volume.Uid = uid
	// TODO 确认下状态码
	volume.Status = 0
	volume.CreateTime = time.Now()

	OkNormal(c, uid)
}
