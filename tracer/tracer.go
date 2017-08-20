package tracer

import (
	"time"

	"go.uber.org/zap"
)

var logger, _ = zap.NewProduction()

func Trace(s string) (string, time.Time) {
	return s, time.Now()
}

func Un(name string, startTime time.Time) {
	defer logger.Sync()
	endTime := time.Now()

	logger.Info("Execution Time",
		zap.String("elapsedTime", endTime.Sub(startTime).String()),
		zap.String("name", name),
	)
}
