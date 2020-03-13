package v2

import (
	"DuckyGo/auth"
	"DuckyGo/conf"
	"DuckyGo/model"
	"DuckyGo/serializer"
	"github.com/dgrijalva/jwt-go"
	"time"
)

// UserLoginService 管理用户登录的服务
type UserLoginService struct {
	UserName string `form:"user_name" json:"user_name" binding:"required,min=5,max=30"`
	Password string `form:"password" json:"password" binding:"required,min=8,max=18"`
}

func GenerateToken(user model.User) (string, error) {
	claims := auth.Jwt{
		jwt.StandardClaims{
			ExpiresAt: time.Now().Add(time.Hour * time.Duration(720)).Unix(),
			IssuedAt:  time.Now().Unix(),
		},
		user,
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	jwtString, err := token.SignedString(conf.SigningKey)
	return jwtString, err
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

	token, err := GenerateToken(user)
	if err != nil {
		return &serializer.Response{
			Code:  serializer.ServerPanicError,
			Error: err.Error(),
		}
	}

	return &serializer.Response{
		Data: token,
	}
}
