package api

import (
	"DuckyGo/service"
	"github.com/gin-gonic/gin"
)

// GetJwtToken 获得Token
func GetJwtToken(c *gin.Context) {
	var service service.GetJwtTokenService
	if err := c.ShouldBind(&service); err == nil {
		res := service.Get()
		c.JSON(200, res.Result())
	} else {
		c.JSON(200, ErrorResponse(err).Result())
	}
}
