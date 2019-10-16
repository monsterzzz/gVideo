package v1

import (
	"fmt"
	"gVideoServer/e"
	"gVideoServer/model"
	"gVideoServer/service/PreService"
	"gVideoServer/util/jwt"
	"gVideoServer/util/serializer"
	"net/http"

	"github.com/gin-contrib/sessions"

	"github.com/gin-gonic/gin"
)

func UserLogin(c *gin.Context) {
	var loginService PreService.UserLoginService
	type Result struct {
		Id          uint   `json:"id"`
		Username    string `json:"username"`
		Nickname    string `json:"nickname"`
		Avatar      string `json:"avatar"`
		Token       string `json:"token"`
		Description string `json:"description"`
	}
	if err := c.ShouldBindJSON(&loginService); err == nil {
		user, err := loginService.LoginCheck()
		if err != nil {
			c.JSON(200, gin.H{
				"token": "",
				"data":  err,
			})
			return
		}
		var result Result
		result.Username = user.Username
		result.Nickname = user.Nickname
		result.Avatar = user.Avatar
		result.Id = user.ID
		result.Description = user.Description
		result.Token = jwt.GenerateToken(jwt.ClaimsMaker(int(user.ID), user.Username, user.Password))
		//c.JSON(200, gin.H{
		//	"token": jwt.GenerateToken(jwt.ClaimsMaker(int(user.ID), user.Username, user.Password)),
		//	"data":  user,
		//})
		session := sessions.Default(c)

		session.Set("user", user)
		session.Save()

		c.JSON(200, &serializer.Response{
			Status: 200,
			Data:   result,
		})

	} else {
		fmt.Println(err)
		c.JSON(http.StatusBadRequest, err.Error())
	}

}

func UserSign(c *gin.Context) {
	var userSignInService PreService.UserSignInService
	err := c.ShouldBindJSON(&userSignInService)

	if err == nil {
		if e := userSignInService.Valid(); e != nil {
			c.JSON(200, e)
		} else {
			e := userSignInService.CreateUser()
			c.JSON(200, e)
		}
	} else {
		c.JSON(200, &serializer.Response{
			Status: 99999,
			Msg:    e.GetErrorMessage(99999),
			Error:  err.Error(),
		})
	}

}

func UpdateNickName(c *gin.Context) {
	type Nickname struct {
		Nickname string `json:"nickname";binding:"required,min=2,max=255"`
	}
	var nickname Nickname
	if err := c.ShouldBindJSON(&nickname); err != nil {
		c.JSON(200, &serializer.Response{
			Status: 66661,
			Data:   "error nickname!",
		})
		return
	}
	newNickName := nickname.Nickname

	var tempUser model.User
	user := serializer.GetSessionUser(c)

	model.DB.Where("nickname = ?", newNickName).Find(&tempUser)

	if tempUser.ID > 0 && tempUser.ID != user.ID {
		c.JSON(200, &serializer.Response{
			Status: 66661,
			Data:   "nickname exists!",
		})
		return
	}
	model.DB.Model(&user).Update("nickname", newNickName)
	c.JSON(200, &serializer.Response{
		Status: 66666,
		Data:   "success",
	})

}

func UpdateDescription(c *gin.Context) {
	type Description struct {
		Description string `json:"description";binding:"required,max=255"`
	}
	var description Description
	if err := c.BindJSON(&description); err != nil {
		c.JSON(200, &serializer.Response{
			Status: 77776,
			Data:   "error",
			Error:  err.Error(),
		})
		return
	}
	newDescription := description.Description

	if len(newDescription) == 0 {
		c.JSON(200, &serializer.Response{
			Status: 77777,
			Data:   "success",
		})
		return
	}
	user := serializer.GetSessionUser(c)

	result := model.DB.Model(&user).Update("description", newDescription)
	var e string
	if result.Error != nil {
		e = result.Error.Error()
	}
	c.JSON(200, &serializer.Response{
		Status: 77777,
		Data:   "success",
		Error:  e,
	})
}

func GetVideoRecords(c *gin.Context) {
	user := serializer.GetSessionUser(c)
	model.DB.Model(&user).Related(&user.Video)
	c.JSON(200, &serializer.Response{
		Status: 1,
		Data:   user.Video,
	})

}
