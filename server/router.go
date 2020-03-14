package server

import (
	"DuckyGo/api"
	"DuckyGo/api/v1"
	"DuckyGo/api/v2"
	"DuckyGo/middleware"
	"fmt"
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
	r.GET("/", api.Index)
	r.GET("ping", api.Ping)

	// v1 最基本网站需要
	if os.Getenv("V1") == "on" {
		sessionGroup := r.Group("/api/v1")
		{

			if os.Getenv("RIM") != "use" {
				panic(fmt.Sprintf("v1 Session验证必须依赖于MySQL以及Redis, 请在环境变量设置RIM为'use', 并且配置MySQL和Redis的连接"))
			}

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

	// v2 特殊情况需要 列如: 微信小程序等无法使用session维持会话的场景
	if os.Getenv("V2") == "on" {

		// 因为v2必须依赖用户模型和Redis, 所以判断是否开启了Redis和MySQL
		if os.Getenv("RIM") != "use" {
			panic(fmt.Sprintf("v2 JWT验证必须依赖于MySQL以及Redis, 请在环境变量设置RIM为'use', 并且配置MySQL和Redis的连接"))
		}


		jwtGroup := r.Group("/api/v2")
		{
			// 注册
			jwtGroup.POST("user/register", v2.UserRegister)

			// 登录
			jwtGroup.POST("user/login", v2.UserLogin)

			// 使用中间件验证.
			jwt := jwtGroup.Group("")
			jwt.Use(middleware.JwtRequired())
			{
				// 查看个人信息
				jwt.GET("user/me", v2.UserMe)
				// 修改密码
				jwt.PUT("user/changepassword", v2.ChangePassword)
				// 注销
				jwt.DELETE("user/logout", v2.Logout)
			}

		}
	}

	return r
}
