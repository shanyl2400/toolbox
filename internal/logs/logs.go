package logs

import "go.uber.org/zap"

var (
	_logger *zap.Logger
)

func init() {
	_logger, _ = zap.NewProduction()
}

func Debug(msg string, fields ...zap.Field) {
	_logger.Debug(msg, fields...)
}

func Info(msg string, fields ...zap.Field) {
	_logger.Info(msg, fields...)
}

func Warn(msg string, fields ...zap.Field) {
	_logger.Warn(msg, fields...)
}

func Error(msg string, fields ...zap.Field) {
	_logger.Error(msg, fields...)
}
