package v1

import (
	"gVideoServer/model"
	"gVideoServer/service/ContentService/VideoService"
	"gVideoServer/util/serializer"

	"github.com/gin-gonic/gin"
)

func VideoFormUpload(c *gin.Context) {
	var uploadForm VideoService.UploadVideoService
	err := c.ShouldBindJSON(&uploadForm)
	if err != nil {
		c.JSON(200, &serializer.Response{
			Status: 300,
			Data:   err,
		})
		return
	}
	user := serializer.GetSessionUser(c)
	video, err := uploadForm.Save(user.ID)
	if err != nil {
		c.JSON(200, &serializer.Response{
			Status: 300,
			Data:   err,
		})
		return
	}
	c.JSON(200, &serializer.Response{
		Status: 400,
		Data:   video.Uuid,
	})
	return
}

func CheckMd5(c *gin.Context) {
	var video1 model.Video
	var video2 model.Video
	md5 := c.Param("md5")
	uuid := c.Query("uuid")
	model.DB.Where("md5 = ?", md5).Find(&video1)
	model.DB.Where("uuid = ?", uuid).Find(&video2)
	if video1.ID > 0 {
		video2.Md5 = video1.Md5
		video2.Path = video1.Path
		video2.IsComplete = 1
		model.DB.Save(&video2)
		c.JSON(200, &serializer.Response{
			Status: 500,
			Data:   1,
		})
	} else {
		c.JSON(200, &serializer.Response{
			Status: 500,
			Data:   0,
		})
	}
}
