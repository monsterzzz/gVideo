package VideoService

import (
	"gVideoServer/model"
	"strings"

	uuid "github.com/satori/go.uuid"
)

type UploadVideoService struct {
	Title       string `form:"title" json:"title" binding:"required,min=2,max=255"`
	Description string `form:"description" json:"description" binding:"required,min=2,max=255"`
}

func (s *UploadVideoService) Valid() error {
	return nil
}

func (s *UploadVideoService) Save(userId uint) (*model.Video, error) {
	var video model.Video
	video.Uuid = strings.ReplaceAll(uuid.NewV4().String(), "-", "")
	video.Title = s.Title
	video.Description = s.Description
	video.UserID = userId
	result := model.DB.Create(&video)
	return &video, result.Error
}
