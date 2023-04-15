package main

import (
	"log"
	"net/http"
	"syscall"
	"time"

	"github.com/fvbock/endless"
	"github.com/gin-gonic/gin"
)

const defAddress = ":8080"

func main() {
	// Membuat Gin Engine dengan default configuration
	router := gin.Default()

	// Membuat HTTP Routing
	router.GET("/welcome", func(ctx *gin.Context) {
		time.Sleep(10 * time.Second)
		ctx.String(http.StatusOK, "Hello Gin Server")
	})

	// Memulai HTTP Server

	// Membuat graceful shutdown dengan library endless
	server := endless.NewServer(defAddress, router)
	server.BeforeBegin = func(add string) {
		// Proses yang dilakukan sebelum memulai layanan
		log.Printf("Actual pid is %d \n", syscall.Getpid())
		log.Println("Service ready to accept request, listen address: [", add, "]")
	}

	if err := server.ListenAndServe(); err != nil {
		log.Fatalln(err)
	}

	/*
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
	*/
}
