package e

import "gVideoServer/util/serializer"

var msg = map[int]string{
	ERROR_PARMA_LOST:       "参数缺失!",
	ERROR_LOGIN_FAIL:       "帐号或密码错误",
	ERROR_NICKNAME_EXIST:   "昵称已存在",
	ERROR_NAME_EXIST:       "帐号已存在",
	ERROR_PASSWORD_CONFIRM: "密码确认错误",
	SUCCESS:                "成功",
}

var ParamLostError = &serializer.Response{
	Status: ERROR_PARMA_LOST,
	Msg:    GetErrorMessage(ERROR_PARMA_LOST),
}

func GetErrorMessage(code int) string {
	return msg[code]
}
