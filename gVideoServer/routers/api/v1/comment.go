package v1

import (
	"gVideoServer/model"
	"gVideoServer/util/serializer"

	"github.com/gin-gonic/gin"
)

func AddComment(c *gin.Context) {
	user := serializer.GetSessionUser(c)

	type ReceiveComment struct {
		VideoUuid string `json:"video_uuid" from:"video_uuid"`
		Content   string `json:"content" from:"content" binding:"required,min=2,max=255"`
	}
	var comment ReceiveComment
	if err := c.ShouldBindJSON(&comment); err != nil {
		c.JSON(200, &serializer.Response{
			Status: 3,
			Data:   "nil comment",
			Error:  err.Error(),
		})
		return
	}
	videoUuid := comment.VideoUuid
	content := comment.Content
	if videoUuid == "" {
		c.JSON(200, &serializer.Response{
			Status: 3,
			Data:   "nil comment",
		})
		return
	}
	var video model.Video
	model.DB.Where("uuid = ?", videoUuid).Find(&video)

	model.DB.Create(&model.Comment{
		UserID:  int(user.ID),
		VideoID: int(video.ID),
		Content: content,
	})
	c.JSON(200, &serializer.Response{
		Status: 4,
		Data:   "success",
	})
}

func GetComments(c *gin.Context) {
	uuid := c.Param("uuid")
	if uuid == "" {
		c.JSON(404, &serializer.Response{})
		return
	}
	var video model.Video

	type Comment struct {
		User    model.User    `json:"user"`
		Comment model.Comment `json:"comment"`
	}

	model.DB.Where("uuid = ?", uuid).Find(&video)
	//model.DB.Model(&video).Related(&video.Comment).Find(&video.Comment)
	model.DB.Model(&video).Related(&video.Comment)
	var comments []Comment

	for i := 0; i < len(video.Comment); i++ {
		var tmpUser model.User
		model.DB.Where("id = ?", video.Comment[i].UserID).Find(&tmpUser)
		comments = append(comments, Comment{User: tmpUser, Comment: video.Comment[i]})
	}

	c.JSON(200, &serializer.Response{
		Data: comments,
	})

}
