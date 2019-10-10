package api

import (
	"DuckyGo/serializer"
	"DuckyGo/service"
	"net/http"

	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
)

// UserRegister 用户注册接口
func UserRegister(c *gin.Context) {
	var service service.UserRegisterService
	if err := c.ShouldBind(&service); err == nil {
		if user, err := service.Register(); err != nil {
			c.JSON(http.StatusOK, err.Result())
		} else {
			res := serializer.Response{Data: serializer.BuildUserResponse(user)}
			c.JSON(http.StatusOK, res.Result())
		}
	} else {
		c.JSON(http.StatusOK, ErrorResponse(err).Result())
	}
}

// UserLogin 用户登录接口
func UserLogin(c *gin.Context) {
	var service service.UserLoginService
	if err := c.ShouldBind(&service); err == nil {
		if user, err := service.Login(); err != nil {
			c.JSON(http.StatusOK, err)
		} else {
			// 设置Session
			s := sessions.Default(c)
			s.Clear()
			s.Set("user_id", user.ID)
			s.Save()

			res := serializer.Response{Data: serializer.BuildUserResponse(user)}
			c.JSON(http.StatusOK, res.Result())
		}
	} else {
		c.JSON(http.StatusOK, ErrorResponse(err).Result())
	}
}

// UserMe 用户详情
func UserMe(c *gin.Context) {
	user := CurrentUser(c)
	res := serializer.Response{Data: serializer.BuildUserResponse(*user)}
	c.JSON(http.StatusOK, res.Result())
}

// ChangePassword 修改密码
func ChangePassword(c *gin.Context) {
	user := CurrentUser(c)
	var service service.ChangePassword
	if err := c.ShouldBind(&service); err == nil {
		res := service.Change(user)
		c.JSON(http.StatusOK, res.Result())
	} else {
		c.JSON(http.StatusOK, ErrorResponse(err).Result())
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
