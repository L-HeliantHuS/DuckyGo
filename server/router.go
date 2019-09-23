package server

import (
	"DuckyGo/api"
	"DuckyGo/middleware"
	"os"

	"github.com/gin-gonic/gin"
)

// NewRouter 路由配置
func NewRouter() *gin.Engine {
	r := gin.Default()

	r.StaticFile("/favicon.ico", "static/favicon.ico")

	// 中间件, 顺序不能改
	r.Use(middleware.SessionCookie(os.Getenv("SESSION_SECRET")))
	r.Use(middleware.Cors())
	r.Use(middleware.CurrentUser())

	// 主页.
	r.Any("/", api.Index)

	// v1 最基本网站需要
	v1 := r.Group("/api/v1")
	{
		v1.GET("ping", api.Ping)

		// 用户注册
		v1.POST("user/register", api.UserRegister)

		// 用户登录
		v1.POST("user/login", api.UserLogin)

		// 需要登录保护的
		auth := v1.Group("")
		auth.Use(middleware.AuthRequired())
		{
			// User Routing
			auth.GET("user/me", api.UserMe)
			auth.DELETE("user/logout", api.UserLogout)
			auth.PUT("user/changepassword", api.ChangePassword)

			// 需要是管理员
			admin := auth.Group("")
			admin.Use(middleware.AuthAdmin())
			{

			}
		}
	}

	// v2 特殊情况需要 列如: 微信小程序等无法使用session维持会话的场景
	v2 := r.Group("/api/v2")
	{
		// 获得token
		v2.GET("sign", api.GetJwtToken)

		// 使用中间件验证.
		jwt := v2.Group("")
		jwt.Use(middleware.JwtRequired())
		{
			jwt.GET("ping", api.HelloJwt)
		}

	}

	return r
}
