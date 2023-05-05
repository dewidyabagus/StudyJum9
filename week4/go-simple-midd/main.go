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

func validationHeaders() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		log.Println("Middleware 1")

		contentType := ctx.GetHeader("Content-Type")
		if contentType != "application/json" {
			ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"message": "Content-Type nilai harus application/json"})
			return
		}

		ctx.Next()
	}
}

func middleware2() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		log.Println("Middleware 2")

		ctx.Next()
	}
}

func main() {
	router := gin.Default()

	// Description pattern
	router.Use(validationHeaders())
	router.Use(middleware2())

	router.POST("/messages", func(ctx *gin.Context) {
		ctx.String(http.StatusOK, "Hello Gin Server")
	})

	server := &http.Server{
		Addr:    defAddress,
		Handler: router,
	}
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
