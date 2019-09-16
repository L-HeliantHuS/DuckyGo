package middleware

import (
	"github.com/gin-contrib/sessions"
	"github.com/gin-contrib/sessions/cookie"
	"github.com/gin-contrib/sessions/redis"
	"github.com/gin-gonic/gin"
	"os"
)

// 特地提供两种解决方案，一种是打算开发成大型项目要用Redis，另一种是不需要Redis和MySQL开发练习使用~

// SessionRedis 初始化Redis-session
func SessionRedis(secret string) gin.HandlerFunc {
	store, _ := redis.NewStore(10, "tcp", os.Getenv("REDIS_ADDR"), os.Getenv("REDIS_PW"), []byte(secret))
	//Also set Secure: true if using SSL, you should though
	store.Options(sessions.Options{HttpOnly: true, MaxAge: 7 * 86400, Path: "/"})
	return sessions.Sessions("gin-session", store)
}

// SessionCookie 初始化Cookie-session
func SessionCookie(secret string) gin.HandlerFunc {
	store := cookie.NewStore([]byte(secret))
	//Also set Secure: true if using SSL, you should though
	store.Options(sessions.Options{HttpOnly: true, MaxAge: 7 * 86400, Path: "/"})
	return sessions.Sessions("gin-session", store)
}
