package service

import (
	"DuckyGo/auth"
	"DuckyGo/conf"
	"DuckyGo/serializer"
	"github.com/dgrijalva/jwt-go"
	"time"
)

// GetJwtTokenService 获得Token的服务
type GetJwtTokenService struct {
}

// Get 获得Token
func (service *GetJwtTokenService) Get() *serializer.Response {
	claims := auth.Jwt{
		jwt.StandardClaims{
			ExpiresAt: time.Now().Add(time.Second * time.Duration(300)).Unix(),
			IssuedAt:  time.Now().Unix(),
		},
		"Default Message.",
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	jwtString, _ := token.SignedString(conf.SigningKey)

	return &serializer.Response{
		Data: jwtString,
	}
}
