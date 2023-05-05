package main

import (
	"encoding/json"
	"fmt"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

const keyRequestID = "X-Request-ID"

type logger struct {
	Level    string `json:"lvl"`
	Time     string `json:"time"`
	Msg      string `json:"msg"`
	Method   string `json:"method"`
	Path     string `json:"path"`
	Status   int    `json:"status"`
	Latency  string `json:"latency"`
	ClientIP string `json:"client_ip"`
	TraceID  string `json:"trace_id"`
	ErrorMsg string `json:"error,omitempty"`
}

func CustomLogger() gin.HandlerFunc {
	return gin.LoggerWithFormatter(func(params gin.LogFormatterParams) string {
		b, _ := json.Marshal(logger{
			Level:    "info",
			Time:     time.Now().Format("2006-01-02 15:04:05.999"),
			Msg:      "Request Completed",
			Method:   params.Method,
			Path:     params.Path,
			Status:   params.StatusCode,
			Latency:  params.Latency.String(),
			ClientIP: params.ClientIP,
			TraceID:  params.Request.Header.Get(keyRequestID),
			ErrorMsg: params.ErrorMessage,
		})

		return fmt.Sprintf("%s\n", string(b))
	})
}

func TrackingID() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		var traceID = ctx.GetHeader(keyRequestID)
		if strings.TrimSpace(traceID) == "" {
			ctx.Request.Header.Set(keyRequestID, uuid.NewString())
		}

		ctx.Next()
	}
}
