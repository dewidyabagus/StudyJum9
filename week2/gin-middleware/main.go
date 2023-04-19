package main

import (
	"context"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/gin-gonic/gin"
)

const defAddress = ":8080"

func main() {
	// Membuat Gin Engine dengan default configuration
	router := gin.Default()

	// Tanpa autentikasi
	router.GET("/welcome", func(ctx *gin.Context) {
		ctx.String(http.StatusOK, "Welcome To Home page")
	})

	// Autentikasi menggunakan basic auth
	basicAuth := router.Group("/route2", gin.BasicAuth(basicAuthAccounts))
	basicAuth.POST("/generate-token", func(ctx *gin.Context) {
		username, _, _ := ctx.Request.BasicAuth()

		token, err := GenerateToken(username)
		if err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
			return
		}

		ctx.JSON(http.StatusOK, gin.H{"token": token})
	})

	// Autentikasi menggunakan JWT
	jwtAuth := router.Group("/route3", JWTMiddlewareAuth(jwtSecretKey))
	jwtAuth.POST("/verify", func(ctx *gin.Context) {
		ctx.JSON(http.StatusOK, gin.H{"message": "OK"})
	})

	// Memulai HTTP Server dan setup graceful shutdown
	server := &http.Server{
		Addr:    defAddress,
		Handler: router,
	}

	// Memulai HTTP server
	go func() {
		if err := server.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Fatalln("listen and serve failed:", err.Error())
		}
	}()

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, os.Interrupt, syscall.SIGTERM)

	<-quit

	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()

	if err := server.Shutdown(ctx); err != nil {
		log.Fatalln("shutdown server failed:", err.Error())
	}
}
