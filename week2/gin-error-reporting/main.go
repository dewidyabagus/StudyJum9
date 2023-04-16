package main

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/gin-gonic/gin"
	errs "github.com/go-errors/errors"
	"github.com/joho/godotenv"
)

const defAddress = ":8080"

type Divide struct {
	N int `form:"n" binding:"min=0"`
	M int `form:"m" binding:"min=0"`
}

func main() {
	if err := godotenv.Load(); err != nil {
		log.Println("load file .env:", err.Error())
	}

	// Membuat object slack client
	slackClient := NewSlackClient(os.Getenv("SLACK_CHANNEL_ID"), os.Getenv("SLACK_TOKEN"))

	// Membuat Gin Engine dengan default configuration
	router := gin.Default()

	// Setup middleware untuk custom recovery dan mengirimkan error message ke slack channel
	var recovery = func(ctx *gin.Context, err any) {
		if _, ok := err.(error); !ok {
			return
		}

		wrapErr := errs.Wrap(err, 4)
		slackClient.ErrorReporting(ctx, wrapErr.Error(), wrapErr.ErrorStack())

		ctx.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": wrapErr.Error()})
	}
	router.Use(gin.CustomRecovery(recovery))

	// Membuat HTTP Routing
	router.GET("/divide", func(ctx *gin.Context) {
		var query Divide
		if err := ctx.ShouldBind(&query); err != nil {
			ctx.JSON(http.StatusOK, gin.H{"message": err.Error()})
			return
		}

		result := query.N / query.M

		ctx.JSON(http.StatusOK, gin.H{"result": fmt.Sprintf("%d / %d = %d", query.N, query.M, result)})
	})

	/*
		Recovery error pastikan dibagian bawah sendiri untuk deskripsi middleware

		router.Use(apiKey)
		router.Use(modifyHeader)
		router.Use(validateJWT)
		router.Use(recoveryError)
	*/

	// Memulai HTTP Server
	// Membuat graceful shutdown menggunakan package net/http
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
