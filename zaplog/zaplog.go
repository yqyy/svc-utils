package zaplog

import (
	"strings"

	"go.uber.org/zap"
)

// NewZapLogger creates a logger
func NewZapLogger(name string, devMode bool, logLevel string, opts ...ModOptions) *zap.Logger {
	level := zap.DebugLevel
	switch strings.ToUpper(logLevel) {
	case "DEBUG":
		level = zap.DebugLevel
	case "INFO":
		level = zap.InfoLevel
	case "WARN":
		level = zap.WarnLevel
	case "ERROR":
		level = zap.ErrorLevel
	case "DPANIC":
		level = zap.DPanicLevel
	}

	var logOpts []ModOptions
	logOpts = append(logOpts, SetAppName(name), SetDevelopment(devMode), SetLevel(level))
	logOpts = append(logOpts, opts...)

	return NewLogger(logOpts...)
}
