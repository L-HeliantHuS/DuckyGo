package middleware

import (
	"DuckyGo/auth"
	"DuckyGo/conf"
	"DuckyGo/serializer"
	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"net/http"
)

// JwtRequired 需要在Header中传递token
func JwtRequired() gin.HandlerFunc {
	return func(c *gin.Context) {
		// 从请求头获得token
		userToken := c.Request.Header.Get("token")
		// 判断请求头中是否有token
		if userToken == "" {
			c.JSON(http.StatusOK, serializer.Response{
				Code: serializer.UserNotPermissionError,
				Msg:  "令牌不能为空！",
			}.Result())
			c.Abort()
			return
		}

		// 解码token值
		token, _ := jwt.ParseWithClaims(userToken, &auth.Jwt{}, func(token *jwt.Token) (interface{}, error) {
			return conf.SigningKey, nil
		})

		if token.Valid != true {
			// 过期或者非正确处理
			c.JSON(http.StatusOK, serializer.Response{
				Code: serializer.UserNotPermissionError,
				Msg:  "令牌错误！",
			}.Result())
			c.Abort()
		}
	}
}
