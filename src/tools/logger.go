package tools

import (
	"errors"
	"fmt"
	"os"
	"path/filepath"
	"strings"
	"sync"

	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"gopkg.in/natefinch/lumberjack.v2"
)

var (
	Logger         *zap.Logger
	initOnce       sync.Once
	mu             sync.RWMutex
	defaultLogPath = "logs/app.log"
)

type LoggerConfig struct {
	LogPath    string
	LogLevel   string
	MaxSize    int
	MaxBackups int
	MaxAge     int
	Compress   bool
	IsStdout   bool
}

// parseLogLevel 解析日志级别字符串
func parseLogLevel(levelStr string) zapcore.Level {
	levelStr = strings.ToLower(strings.TrimSpace(levelStr))
	switch levelStr {
	case "debug":
		return zap.DebugLevel
	case "info":
		return zap.InfoLevel
	case "warn":
		return zap.WarnLevel
	case "error":
		return zap.ErrorLevel
	default:
		return zap.InfoLevel
	}
}

// validateConfig 验证并设置默认值
func validateConfig(config *LoggerConfig) error {
	if config.LogPath == "" {
		return errors.New("LogPath cannot be empty")
	}
	if config.LogLevel == "" {
		config.LogLevel = "info"
	}
	if config.MaxSize <= 0 {
		config.MaxSize = 100
	}
	if config.MaxBackups < 0 {
		config.MaxBackups = 10
	}
	if config.MaxAge <= 0 {
		config.MaxAge = 7
	}

	logDir := filepath.Dir(config.LogPath)
	if logDir != "." && logDir != "" {
		if err := os.MkdirAll(logDir, 0755); err != nil {
			return fmt.Errorf("failed to create log directory: %w", err)
		}
		// 检查可写性
		testFile := filepath.Join(logDir, ".write_test")
		if f, err := os.Create(testFile); err != nil {
			return fmt.Errorf("log directory is not writable: %w", err)
		} else {
			f.Close()
			os.Remove(testFile)
		}
	}
	return nil
}

// buildLogger 构建 logger（公共函数）
func buildLogger(config LoggerConfig) *zap.Logger {
	hook := lumberjack.Logger{
		Filename:   config.LogPath,
		MaxSize:    config.MaxSize,
		MaxBackups: config.MaxBackups,
		MaxAge:     config.MaxAge,
		Compress:   config.Compress,
	}

	fileSyncer := zapcore.AddSync(&hook)
	level := parseLogLevel(config.LogLevel)

	encoderConfig := zapcore.EncoderConfig{
		TimeKey:        "time",
		LevelKey:       "level",
		NameKey:        "logger",
		CallerKey:      "linenum",
		MessageKey:     "msg",
		StacktraceKey:  "stacktrace",
		LineEnding:     zapcore.DefaultLineEnding,
		EncodeLevel:    zapcore.LowercaseLevelEncoder,
		EncodeTime:     zapcore.ISO8601TimeEncoder,
		EncodeDuration: zapcore.SecondsDurationEncoder,
		EncodeCaller:   zapcore.FullCallerEncoder,
		EncodeName:     zapcore.FullNameEncoder,
	}

	writer := fileSyncer
	if config.IsStdout {
		writer = zapcore.NewMultiWriteSyncer(zapcore.AddSync(os.Stdout), fileSyncer)
	}

	core := zapcore.NewCore(zapcore.NewJSONEncoder(encoderConfig), writer, level)
	return zap.New(core, zap.AddCaller(), zap.Development())
}

// initDefaultLogger 使用默认配置初始化
func initDefaultLogger() {
	config := LoggerConfig{
		LogPath:    defaultLogPath,
		LogLevel:   "info",
		MaxSize:    100,
		MaxBackups: 10,
		MaxAge:     7,
		Compress:   false,
		IsStdout:   true,
	}
	_ = validateConfig(&config) // 忽略错误，使用默认值

	mu.Lock()
	Logger = buildLogger(config)
	mu.Unlock()

	if Logger != nil {
		Logger.Info("Logger auto-initialized", zap.String("log_level", config.LogLevel))
	}
}

// GetLogger 获取 Logger，未初始化则自动初始化
func GetLogger() *zap.Logger {
	mu.RLock()
	if Logger != nil {
		mu.RUnlock()
		return Logger
	}
	mu.RUnlock()

	initOnce.Do(initDefaultLogger)
	return Logger
}

// MustGetLogger 获取 Logger，未初始化则自动初始化
func MustGetLogger() *zap.Logger {
	return GetLogger()
}

// InitLogger 初始化日志器
func InitLogger(config LoggerConfig) error {
	mu.RLock()
	if Logger != nil {
		mu.RUnlock()
		return errors.New("logger has already been initialized")
	}
	mu.RUnlock()

	var initErr error
	initOnce.Do(func() {
		if err := validateConfig(&config); err != nil {
			initErr = err
			return
		}

		mu.Lock()
		Logger = buildLogger(config)
		mu.Unlock()

		if Logger != nil {
			Logger.Info("Logger initialized",
				zap.String("log_path", config.LogPath),
				zap.String("log_level", config.LogLevel),
			)
		}
	})
	return initErr
}

// WithTraceID 返回带 trace_id 的 Logger
func WithTraceID(traceID string) *zap.Logger {
	logger := GetLogger()
	if traceID != "" {
		return logger.With(zap.String("trace_id", traceID))
	}
	return logger
}

// Sync 同步日志缓冲区
func Sync() error {
	if logger := GetLogger(); logger != nil {
		return logger.Sync()
	}
	return nil
}
