package main

import (
	"context"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	zapLog "learning/go-build-app/logger"

	"github.com/gin-gonic/gin"
)

type TraceType string

const defAddress = ":8080"
const TraceIDKey TraceType = "trace_id"

var mlog *zapLog.Logger

func main() {
	mlog, _ = zapLog.New("info", "stdout")

	r := gin.New()
	r.Use(TrackingID())
	r.Use(CustomLogger())
	// gin.SetMode(gin.ReleaseMode)

	r.GET("/welcome", func(ctx *gin.Context) {
		traceID := ctx.GetHeader(keyRequestID)

		HandleFunc1(
			context.WithValue(ctx, TraceIDKey, traceID),
		)

		ctx.String(http.StatusOK, "Hello Gin Server")
	})

	server := &http.Server{
		Addr:    defAddress,
		Handler: r,
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
