package controller

import (
	"bipu-backend/src/main/model"

	"github.com/gin-gonic/gin"
)

func QueryVolumesAll(c *gin.Context) {
	volumes := model.QueryVolumesAll()

	c.JSON(200, gin.H{"volumes": volumes})
}

func QueryVolumesByUid(c *gin.Context) {
	uid := c.Params.ByName("uid")
	volumes := model.QueryVolumesByUid(uid)

	c.JSON(200, gin.H{"volumes": volumes})
}
