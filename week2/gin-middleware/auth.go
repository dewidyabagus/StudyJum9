package main

import (
	"time"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt"
)

const (
	jwtSecretKey = "jwt-secret-key"
	expires      = 10 * time.Minute
)

var basicAuthAccounts = gin.Accounts{
	"admin": "admin123",
	"user":  "user123",
}

func GenerateToken(username string) (token string, err error) {
	eJWT := jwt.NewWithClaims(
		jwt.SigningMethodHS256,
		jwt.MapClaims{
			"username": username,
			"exp":      time.Now().Add(expires).Unix(),
		},
	)

	return eJWT.SignedString([]byte(jwtSecretKey))
}
