package v1

import (
	"gVideoServer/model"
	"gVideoServer/util/serializer"
	"net/http"
	"os"
	"strings"

	"github.com/gin-gonic/gin"
	uuid "github.com/satori/go.uuid"
)

const (
	STATIC_IMG_URI = "/static/img/"
)

func UploadImage(c *gin.Context) {
	videoUuid := c.PostForm("videoUuid")
	file, err := c.FormFile("file")
	if err != nil {
		c.JSON(http.StatusBadRequest, &serializer.Response{
			Status: 50001,
			Msg:    err.Error(),
		})
		return
	}
	savePath := os.Getenv("IMAGE_SAVE_PATH")
	randomName := strings.ReplaceAll(uuid.NewV4().String(), "-", "")
	fullName := savePath + randomName + ".jpg"
	err = c.SaveUploadedFile(file, fullName)
	if err != nil {
		c.JSON(200, &serializer.Response{
			Status: 50002,
			Msg:    err.Error(),
		})
		return
	}
	var video model.Video
	model.DB.Where("uuid = ?", videoUuid).Find(&video)
	imgPath := STATIC_IMG_URI + randomName + ".jpg"
	video.Picture = imgPath
	model.DB.Save(&video)
	c.JSON(200, &serializer.Response{
		Status: 50000,
		Data:   imgPath,
		Msg:    "image successfully saved",
	})
}

func AvatarUpload(c *gin.Context) {
	file, err := c.FormFile("file")
	if err != nil {
		c.JSON(http.StatusBadRequest, &serializer.Response{
			Status: 50001,
			Msg:    err.Error(),
		})
		return
	}
	savePath := os.Getenv("IMAGE_SAVE_PATH")
	randomName := strings.ReplaceAll(uuid.NewV4().String(), "-", "")
	fullName := savePath + randomName + ".jpg"
	imgPath := STATIC_IMG_URI + randomName + ".jpg"
	err = c.SaveUploadedFile(file, fullName)
	if err != nil {
		c.JSON(200, &serializer.Response{
			Status: 50002,
			Msg:    err.Error(),
		})
		return
	}
	user := serializer.GetSessionUser(c)
	user.Avatar = imgPath
	model.DB.Model(&user).Update("avatar", imgPath)
	c.JSON(200, &serializer.Response{
		Status: 50000,
		Data:   gin.H{"path": imgPath},
		Msg:    "image successfully saved",
	})
}
