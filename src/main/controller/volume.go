package controller

import (
	"bipu-backend/src/main/model"

	"github.com/gin-gonic/gin"
)

func QueryVolumesAll(c *gin.Context) {
	volumes := model.QueryVolumesAll()
	OkNormal(c, volumes)
}

func QueryVolumesByUid(c *gin.Context) {
	uid := c.Params.ByName("uid")
	volumes := model.QueryVolumesByUid(uid)

	OkNormal(c, volumes)
}

func QueryVolumesByVid(c *gin.Context) {
	vid := c.Params.ByName("vid")
	volumes := model.QueryVolumesByVid(vid)

	OkNormal(c, volumes)
}
