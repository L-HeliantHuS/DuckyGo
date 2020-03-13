package v2

import (
	"DuckyGo/model"
	"DuckyGo/serializer"
)

// ChangePassword 修改用户密码
type ChangePassword struct {
	Password        string `form:"password" json:"password" binding:"required,min=8,max=40"`
	PasswordConfirm string `form:"password_confirm" json:"password_confirm" binding:"required,min=8,max=40"`
}

// Valid 验证表单
func (service *ChangePassword) Valid() *serializer.Response {
	if service.PasswordConfirm != service.Password {
		return &serializer.Response{
			Code: serializer.UserInputError,
			Msg:  "两次输入的密码不相同",
		}
	}

	return nil
}

// Change 修改密码
func (service *ChangePassword) Change(user *model.User) *serializer.Response {

	// 表单验证
	if err := service.Valid(); err != nil {
		return err
	}

	// 加密密码
	if err := user.SetPassword(service.Password); err != nil {
		return &serializer.Response{
			Code: serializer.ServerPanicError,
			Msg:  "加密密码出现错误.",
		}
	}

	// 更新数据库
	if err := model.DB.Save(&user).Error; err != nil {
		return &serializer.Response{
			Code: serializer.DatabaseWriteError,
			Msg:  "更新数据库出现错误。",
		}
	}

	return &serializer.Response{
		Data: serializer.BuildUser(*user),
		Msg:  "修改密码成功！",
	}
}
