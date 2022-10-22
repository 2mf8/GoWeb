package middleware

import (
	"time"
	"github.com/dgtijalva/jwt-go"
)

var jwtKey = []byte("@#$%^&2mf8kequ._AFJK")

type JwtClaims struct{
	UserId uint64
}

func CreateJwt(username, passward string, timeout int) (string, bool) {
	expiresAt := time.Now().Add(time.Hour * time.Duration(timeout)).Unix()

	claims := JwtClaims{}
}