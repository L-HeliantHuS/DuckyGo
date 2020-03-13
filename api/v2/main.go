package v2

import (
	"DuckyGo/serializer"
	"github.com/gin-gonic/gin"
	"net/http"
)

// HelloJwt 通过JwtToken验证查看接口
func HelloJwt(c *gin.Context) {
	c.JSON(http.StatusOK, serializer.Response{
		Msg: "Hello!",
	}.Result())
}

