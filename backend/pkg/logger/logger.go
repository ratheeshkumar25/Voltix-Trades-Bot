// By Ratheesh Kumar Golang Developer
package logger

import (
	"log"
	"os"

	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

type Logger struct {
	Logger *zap.SugaredLogger
}

var globalLogger *Logger

// NewLogger initializes and returns a new logger instance
func NewLogger() *Logger {
	var logger *zap.Logger
	var err error

	env := os.Getenv("ENV")
	if env == "production" {
		logger, err = zap.NewProduction()
	} else {
		// Development config with better readability
		config := zap.NewDevelopmentConfig()
		config.EncoderConfig.EncodeLevel = zapcore.CapitalColorLevelEncoder
		logger, err = config.Build()
	}

	if err != nil {
		log.Fatalf("Failed to initialize logger: %v", err)
	}

	defer logger.Sync() // Flush any buffered log entries

	sugaredLogger := logger.Sugar()
	globalLogger = &Logger{Logger: sugaredLogger}

	return globalLogger
}

// GetLogger returns the global logger instance
func GetLogger() *Logger {
	if globalLogger == nil {
		globalLogger = NewLogger()
	}
	return globalLogger
}

// Info logs an info level message
func (l *Logger) Info(msg string, fields ...interface{}) {
	l.Logger.Infow(msg, fields...)
}

// Error logs an error level message
func (l *Logger) Error(msg string, fields ...interface{}) {
	l.Logger.Errorw(msg, fields...)
}

// Debug logs a debug level message
func (l *Logger) Debug(msg string, fields ...interface{}) {
	l.Logger.Debugw(msg, fields...)
}

// Warn logs a warning level message
func (l *Logger) Warn(msg string, fields ...interface{}) {
	l.Logger.Warnw(msg, fields...)
}

// Fatal logs a fatal level message and exits
func (l *Logger) Fatal(msg string, fields ...interface{}) {
	l.Logger.Fatalw(msg, fields...)
}

// InfoF logs formatted info message
func (l *Logger) InfoF(format string, args ...interface{}) {
	l.Logger.Infof(format, args...)
}

// ErrorF logs formatted error message
func (l *Logger) ErrorF(format string, args ...interface{}) {
	l.Logger.Errorf(format, args...)
}

// DebugF logs formatted debug message
func (l *Logger) DebugF(format string, args ...interface{}) {
	l.Logger.Debugf(format, args...)
}

// WarnF logs formatted warning message
func (l *Logger) WarnF(format string, args ...interface{}) {
	l.Logger.Warnf(format, args...)
}
