package v1

import (
	"DuckyGo/api"
	"DuckyGo/serializer"
	"DuckyGo/service/v1"
	"net/http"

	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
)

// UserRegister 用户注册接口
func UserRegister(c *gin.Context) {
	var service v1.UserRegisterService
	if err := c.ShouldBind(&service); err == nil {
		res := service.Register()
		c.JSON(http.StatusOK, res.Result())
	} else {
		c.JSON(http.StatusOK, api.ErrorResponse(err).Result())
	}
}

// UserLogin 用户登录接口
func UserLogin(c *gin.Context) {
	var service v1.UserLoginService
	if err := c.ShouldBind(&service); err == nil {
		res := service.Login()
		if user, ok := res.Data.(serializer.UserResponse); ok {
			// 设置Session
			s := sessions.Default(c)
			s.Clear()
			s.Set("user_id", user.Data.ID)
			s.Save()
		}

		c.JSON(http.StatusOK, res.Result())

	} else {
		c.JSON(http.StatusOK, api.ErrorResponse(err).Result())
	}
}

// UserMe 用户详情
func UserMe(c *gin.Context) {
	user := api.CurrentUser(c)
	res := serializer.Response{Data: serializer.BuildUserResponse(*user)}
	c.JSON(http.StatusOK, res.Result())
}

// ChangePassword 修改密码
func ChangePassword(c *gin.Context) {
	user := api.CurrentUser(c)
	var service v1.ChangePassword
	if err := c.ShouldBind(&service); err == nil {
		res := service.Change(user)
		c.JSON(http.StatusOK, res.Result())
	} else {
		c.JSON(http.StatusOK, api.ErrorResponse(err).Result())
	}
}

// UserLogout 用户登出
func UserLogout(c *gin.Context) {
	s := sessions.Default(c)
	s.Clear()
	s.Save()
	c.JSON(http.StatusOK, serializer.Response{
		Msg: "登出成功",
	}.Result())
}
