package log

import (
	"context"

	"github.com/spf13/cast"
	"gitlab.com/gobang/logger"
)

var pkgLogger logger.Logger

func formatLogs(messages ...interface{}) (logRecord []logger.Field) {
	for index, msg := range messages {
		logRecord = append(logRecord, logger.ToField("_message_"+cast.ToString(index), msg))
	}

	return
}

func Error(ctx context.Context, title string, messages ...interface{}) {
	fields := formatLogs(messages...)
	pkgLogger.Error(ctx, title, fields...)
}

func Info(ctx context.Context, title string, messages ...interface{}) {
	fields := formatLogs(messages...)
	pkgLogger.Info(ctx, title, fields...)
}
