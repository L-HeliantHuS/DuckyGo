package api

import (
	"DuckyGo/service"
	"github.com/gin-gonic/gin"
	"net/http"
)

// GetJwtToken 获得Token
func GetJwtToken(c *gin.Context) {
	var service service.GetJwtTokenService
	if err := c.ShouldBind(&service); err == nil {
		res := service.Get()
		c.JSON(http.StatusOK, res.Result())
	} else {
		c.JSON(http.StatusOK, ErrorResponse(err).Result())
	}
}
