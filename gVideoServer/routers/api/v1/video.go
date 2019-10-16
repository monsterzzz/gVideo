package v1

import (
	"fmt"
	"gVideoServer/model"
	"gVideoServer/util/fileUtil"
	"gVideoServer/util/serializer"
	"os"
	"os/exec"
	"strconv"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
)

func GetAllVideo(c *gin.Context) {

	type EasyVideo struct {
		Uuid        string `json:"uuid"`
		Picture     string `json:"avatar"`
		Title       string `json:"title"`
		Description string `json:"description"`
	}
	var videos []EasyVideo
	model.DB.Table("videos").Select("uuid,picture,title,description").Scan(&videos)
	c.JSON(200, &serializer.Response{
		Status: 54545,
		Msg:    "dddd",
		Data:   videos,
	})
}

func UploadVideo(c *gin.Context) {
	file, err := c.FormFile("file")
	md5 := c.PostForm("md5")
	currentChunk := c.PostForm("currentChunk")
	totalChunk := c.PostForm("totalChunks")
	fileType := c.PostForm("fileType")
	videoUuid := c.PostForm("uuid")

	intTotalChunks, err := strconv.Atoi(totalChunk)
	if err != nil {
		fmt.Println(err)
	}
	intCurrentChunk, err := strconv.Atoi(currentChunk)
	if err != nil {
		fmt.Println(err)
	}

	dirPath := fmt.Sprintf("./videoRepo/tmp/%s", md5)
	if intCurrentChunk == 0 {
		os.Mkdir(dirPath, os.ModePerm)
	}
	c.SaveUploadedFile(file, dirPath+fmt.Sprintf("/chunk_%s.tmp", currentChunk))

	if intCurrentChunk == intTotalChunks-1 {
		fileUtil.MergeFiles(dirPath, fileType)
		var video model.Video
		var tmp model.Video
		model.DB.Where("uuid = ?", videoUuid).Find(&video)

		tmp.Path = "./videoRepo/fullFile/" + md5 + "." + fileType
		tmp.Md5 = md5
		tmp.IsComplete = 1
		model.DB.Model(&video).Updates(tmp)

		c.JSON(200, &serializer.Response{
			Status: 3,
			Data:   video.Uuid,
		})
		return
	}
	c.JSON(200, &serializer.Response{
		Status: 5,
	})

}

func TestVideo(c *gin.Context) {
	uuid := c.Param("uuid")
	var video model.Video
	model.DB.Where("uuid = ?", uuid).Find(&video)
	var owner model.User
	type UserResponse struct {
		Nickname    string `json:"nickname"`
		Avatar      string `json:"avatar"`
		Id          uint   `json:"id"`
		Description string `json:"description"`
	}
	var ures UserResponse
	model.DB.Where("id = ?", video.UserID).Find(&owner).Scan(&ures)
	videoPath := video.Path
	outputPath := "./videoRepo/fullFile/chunkTs/" + video.Md5 + "/" + video.Md5 + ".m3u8"
	err := os.Mkdir("./videoRepo/fullFile/chunkTs/"+video.Md5, os.ModePerm)
	if err == nil {
		command := fmt.Sprintf("ffmpeg -i %s -vcodec copy -acodec copy -hls_time 60 -hls_list_size 0 %s", videoPath, outputPath)
		commandList := strings.Split(command, " ")
		cmder := exec.Command(commandList[0], commandList[1:]...)
		err = cmder.Run()
		if err != nil {
			fmt.Println(err)
		}
	}

	type Result struct {
		User        UserResponse `json:"user"`
		Title       string       `json:"title"`
		Description string       `json:"description"`
		CreatedAt   time.Time    `json:"-"`
		Time        string       `json:"time"`
		M3u8        string       `json:"m3u8"`
		PlayCount   int          `json:"play_count"`
		LikesCount  int          `json:"likes_count"`
	}
	var result = &Result{
		User:        ures,
		Title:       video.Title,
		Description: video.Description,
		CreatedAt:   video.CreatedAt,
		M3u8:        "/static1/fullFile/chunkTs/" + video.Md5 + "/" + video.Md5 + ".m3u8",
		Time:        video.CreatedAt.Format("2006-01-02 15:04:05"),
		PlayCount:   video.PlayCount,
		LikesCount:  video.LikesCount,
	}
	c.JSON(200, &serializer.Response{
		Status: 200,
		Data:   result,
	})
	//c.Redirect(302, "/api/static1/fullFile/chunkTs/"+video.Md5+"/"+video.Md5+".m3u8")
}
