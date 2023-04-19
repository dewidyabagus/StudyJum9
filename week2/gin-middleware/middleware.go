package main

import (
	"errors"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt"
)

func JWTMiddlewareAuth(jwtSecretKey string) gin.HandlerFunc {
	return func(c *gin.Context) {
		token := strings.Replace(
			c.GetHeader("Authorization"), "Bearer ", "", 1,
		)

		if token = strings.TrimSpace(token); token == "" {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"message": "invalid token"})
			return
		}

		res, err := jwt.Parse(token, func(t *jwt.Token) (interface{}, error) {
			if method, ok := t.Method.(*jwt.SigningMethodHMAC); !ok || method != jwt.SigningMethodHS256 {
				return nil, errors.New("signing method invalid")
			}

			return []byte(jwtSecretKey), nil
		})
		if err != nil {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"message": err.Error()})
			return
		}

		c.Set("user", res)
		c.Next()
	}
}
