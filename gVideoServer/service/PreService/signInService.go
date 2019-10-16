package PreService

import (
	"fmt"
	"gVideoServer/e"
	"gVideoServer/model"
	"gVideoServer/util/serializer"
)

type UserSignInService struct {
	Username        string `form:"username" json:"username" binding:"required,min=4,max=40"`
	Password        string `form:"password" json:"password" binding:"required,min=6,max=16"`
	Nickname        string `form:"nickname" json:"nickname" binding:"required,min=2,max=10"`
	PasswordConfirm string `form:"password_confirm" json:"password_confirm" binding:"required,min=6,max=16"`
}

func (s *UserSignInService) Valid() *serializer.Response {

	user, _ := model.ExistsUserName(s.Username)
	fmt.Println(user)
	if user != nil {
		return &serializer.Response{
			Status: 30001,
			Msg:    e.GetErrorMessage(30001),
			//Error: err.Error(),
		}
	}
	user, _ = model.ExistsNickname(s.Nickname)
	if user != nil {
		return &serializer.Response{
			Status: 30002,
			Msg:    e.GetErrorMessage(30002),
			//Error: err.Error(),
		}
	}

	if s.Password != s.PasswordConfirm {
		return &serializer.Response{
			Status: 30003,
			Msg:    e.GetErrorMessage(30003),
		}
	}

	return nil
}

func (s *UserSignInService) CreateUser() *serializer.Response {

	user := model.User{
		Username: s.Username,
		Nickname: s.Nickname,
	}

	err := user.SetPassword(s.Password)
	if err != nil {
		return &serializer.Response{
			Status: 30007,
			Msg:    e.GetErrorMessage(30007),
			Error:  err.Error(),
		}
	}

	e := model.CreateUser(&user)
	if e != nil {
		return &serializer.Response{
			Status: 30005,
			Msg:    "Unknown error",
			Error:  e.Error(),
		}
	}

	return &serializer.Response{
		Status: 30006,
		Msg:    "success",
	}
}
