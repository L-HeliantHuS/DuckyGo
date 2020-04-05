package middleware

import (
	"DuckyGo/model"
	"DuckyGo/serializer"
	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
	"net/http"
)

// CurrentUser 获取登录用户
func CurrentUser() gin.HandlerFunc {
	return func(c *gin.Context) {
		session := sessions.Default(c)
		uid := session.Get("user_id")
		if uid != nil {
			user, err := model.GetUser(uid)
			if err == nil {
				c.Set("user", &user)
			}
		}
		c.Next()
	}
}

// AuthRequired 需要登录
func AuthRequired() gin.HandlerFunc {
	return func(c *gin.Context) {
		if user, _ := c.Get("user"); user != nil {
			if _, ok := user.(*model.User); ok {
				c.Next()
				return
			}
		}

		c.JSON(http.StatusOK, serializer.Response{
			Code: serializer.UserNotPermissionError,
			Msg:  "需要登录",
		}.Result())
		c.Abort()
	}
}

// 必须为管理员
func AuthAdmin() gin.HandlerFunc {
	return func(c *gin.Context) {
		if user, _ := c.Get("user"); user != nil {
			if user.(*model.User).SuperUser {
				c.Next()
				return
			}
		}
		c.JSON(http.StatusOK, serializer.Response{
			Code: serializer.UserNotPermissionError,
			Msg:  "你没有权限进行此操作.",
		}.Result())
		c.Abort()
	}
}
