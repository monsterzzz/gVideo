package PreService

import (
	"gVideoServer/e"
	"gVideoServer/model"
	"gVideoServer/util/serializer"
)

type UserLoginService struct {
	Username string `form:"username" json:"username" binding:"required,min=4,max=40"`
	Password string `form:"password" json:"password" binding:"required,min=6,max=40"`
}

func (s *UserLoginService) LoginCheck() (*model.User, *serializer.Response) {
	user, err := model.ExistsUserName(s.Username)
	if user == nil {
		code := 20001
		return user, &serializer.Response{
			Status: code,
			Msg:    e.GetErrorMessage(code),
			Error:  err.Error(),
		}
	}
	if err := user.CheckPassword(s.Password); err != nil {
		code := 20001
		return user, &serializer.Response{
			Status: code,
			Msg:    e.GetErrorMessage(code),
			Error:  err.Error(),
		}
	}
	return user, nil

}
