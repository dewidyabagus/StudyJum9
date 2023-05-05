package main

import (
	"context"

	zapLog "learning/go-build-app/logger"
)

func HandleFunc1(ctx context.Context) {
	traceID := ctx.Value(TraceIDKey).(string)
	mlog.Info("pesan dari fungsi HandleFunc1", zapLog.String("trace_id", traceID))

	HandleFunc2(ctx)
}

// HandleFunc2(ctx context.Context, params ....) konversi penulisan params dengan context
func HandleFunc2(ctx context.Context) {
	traceID := ctx.Value(TraceIDKey).(string)
	mlog.Info("pesan dari fungsi HandleFunc2", zapLog.String("trace_id", traceID))
}
