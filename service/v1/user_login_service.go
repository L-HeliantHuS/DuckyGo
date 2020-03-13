package v1

import (
	"DuckyGo/model"
	"DuckyGo/serializer"
)

// UserLoginService 管理用户登录的服务
type UserLoginService struct {
	UserName string `form:"user_name" json:"user_name" binding:"required,min=5,max=30"`
	Password string `form:"password" json:"password" binding:"required,min=8,max=18"`
}

// Login 用户登录函数
func (service *UserLoginService) Login() *serializer.Response {
	var user model.User

	if err := model.DB.Where("user_name = ?", service.UserName).First(&user).Error; err != nil {
		return &serializer.Response{
			Code: serializer.UserNotFoundError,
			Msg:  "账号或者密码错误",
		}
	}

	if user.CheckPassword(service.Password) == false {
		return &serializer.Response{
			Code: serializer.UserPasswordError,
			Msg:  "账号或密码错误",
		}
	}
	return &serializer.Response{
		Data: serializer.BuildUserResponse(user),
	}
}
