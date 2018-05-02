package controller

import (
	"bipu-backend/src/main/model"

	"github.com/gin-gonic/gin"
)

func QueryVolumesAll(c *gin.Context) {
	volumes := model.QueryVolumesAll()

	c.JSON(200, gin.H{"volumes": volumes})
}
