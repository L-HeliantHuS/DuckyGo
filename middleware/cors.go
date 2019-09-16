package middleware

import (
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

// Cors 跨域配置
func Cors() gin.HandlerFunc {
	config := cors.DefaultConfig()
	config.AllowMethods = []string{"GET", "POST", "PUT", "PATCH", "DELETE", "HEAD", "OPTIONS"}
	config.AllowHeaders = []string{"Origin", "Content-Length", "Content-Type", "Cookie"}
	// 运行再Release模式下才会进行跨域保护 保证开发过程中不会被跨域困扰~
	if gin.Mode() == gin.ReleaseMode {
		config.AllowOrigins = []string{"http://www.example.com"}
	} else {
		config.AllowAllOrigins = true
	}
	config.AllowCredentials = true
	return cors.New(config)
}
