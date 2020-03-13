package v1

import (
	"DuckyGo/serializer"
	"github.com/gin-gonic/gin"
	"net/http"
)

// Index 主页
func Index(c *gin.Context) {
	c.String(http.StatusOK, "================   Welcome to DuckyGo Restful API Index Page!     https://github.com/L-HeliantHuS/DuckyGo   ================")
}

// Ping 状态检查页面
func Ping(c *gin.Context) {
	c.JSON(http.StatusOK, serializer.Response{
		Msg: "Pong",
	}.Result())
}