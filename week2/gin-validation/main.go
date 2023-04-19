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

// Daftar tag binding yang digunakan (a.k.a validate tags) bisa dilihat
// di https://pkg.go.dev/github.com/go-playground/validator/v10
type Registration struct {
	Email           string `json:"email" binding:"required,email"`
	FirstName       string `json:"first_name" binding:"required,min=5,max=100"`
	LastName        string `json:"last_name" binding:"required,min=5,max=100"`
	Password        string `json:"password" binding:"required,min=6,max=12"`
	ConfirmPassword string `json:"confirm_password" binding:"eqfield=Password"`
	Options         []int  `json:"options" binding:"min=1,dive,oneof=1 2 3 4 5"`
}

func main() {
	// Membuat Gin Engine dengan default configuration
	router := gin.Default()

	// Memulai HTTP Server dan setup graceful shutdown
	server := &http.Server{
		Addr:    defAddress,
		Handler: router,
	}

	// Rotuing
	router.POST("/registry", func(ctx *gin.Context) {
		var payload Registration
		if err := ctx.ShouldBind(&payload); err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{"error": "bad request", "validation": TranslateFormError(err)})
			return
		}

		ctx.JSON(http.StatusCreated, gin.H{"message": "registrasi berhasil dilakukan"})
	})

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
