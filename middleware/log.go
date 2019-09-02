package middleware

import (
	"github.com/gin-gonic/gin"
	"DuckyGo/util"
)

func SaveLog() gin.HandlerFunc {
	return func(c *gin.Context) {
		go func() {
			url := c.Request.URL.Path
			ua := c.Request.Header["User-Agent"][0]
			util.Log().Info(ua + " | " + url)

		}()
		c.Next()
	}
}
