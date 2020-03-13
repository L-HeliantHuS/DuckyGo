package server

import (
	"DuckyGo/api/v1"
	"DuckyGo/api/v2"
	"DuckyGo/middleware"
	"DuckyGo/model"
	"os"

	"github.com/gin-gonic/gin"
)

// NewRouter 路由配置
func NewRouter() *gin.Engine {
	r := gin.Default()

	r.StaticFile("/favicon.ico", "static/favicon.ico")

	// 中间件, 顺序不能改
	// 启动Redis的情况下将切换成Redis保存Session.
	if os.Getenv("RIM") == "use" {
		r.Use(middleware.SessionRedis(os.Getenv("SESSION_SECRET")))
	} else {
		r.Use(middleware.SessionCookie(os.Getenv("SESSION_SECRET")))
	}
	r.Use(middleware.Cors())
	r.Use(middleware.CurrentUser())

	// 主页.
	r.GET("/", v1.Index)

	// v1 最基本网站需要
	if os.Getenv("v1") == "on" {
		sessionGroup := r.Group("/api/v1")
		{
			sessionGroup.GET("ping", v1.Ping)

			// 如果没连接数据库就可以不用启动用户模型了.
			if model.DB != nil {
				// 用户注册
				sessionGroup.POST("user/register", v1.UserRegister)

				// 用户登录
				sessionGroup.POST("user/login", v1.UserLogin)

				// 需要登录保护的
				auth := sessionGroup.Group("")
				auth.Use(middleware.AuthRequired())
				{
					// User Routing
					auth.GET("user/me", v1.UserMe)
					auth.DELETE("user/logout", v1.UserLogout)
					auth.PUT("user/changepassword", v1.ChangePassword)

					// 需要是管理员
					admin := auth.Group("")
					admin.Use(middleware.AuthAdmin())
					{

					}
				}
			}
		}
	}

	// v2 特殊情况需要 列如: 微信小程序等无法使用session维持会话的场景
	if os.Getenv("v2") == "on" {
		jwtGroup := r.Group("/api/v2")
		{
			// 注册
			jwtGroup.POST("user/register", v2.UserRegister)

			// 登录获得Token
			jwtGroup.POST("user/login", v2.UserLogin)

			// 使用中间件验证.
			jwt := jwtGroup.Group("")
			jwt.Use(middleware.JwtRequired())
			{
				jwt.GET("user/me", v2.UserMe)
				jwt.PUT("user/changepassword", v2.ChangePassword)
				jwt.GET("ping", v2.HelloJwt)
			}

		}
	}

	return r
}
