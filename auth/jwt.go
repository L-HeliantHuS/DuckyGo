package auth

import (
	"DuckyGo/model"
	"github.com/dgrijalva/jwt-go"
)

// Jwt 需要编码的结构体
type Jwt struct {
	jwt.StandardClaims
	Data model.User
}
