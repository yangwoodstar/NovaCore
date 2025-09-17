package tools

import (
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"gopkg.in/natefinch/lumberjack.v2"
	"os"
	"sync"
)

var (
	Logger *zap.Logger
	once   sync.Once
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

func GetLogger() *zap.Logger {
	return Logger
}

func InitLogger(config LoggerConfig) {
	// 日志分割
	once.Do(func() {
		hook := lumberjack.Logger{
			Filename:   config.LogPath,    // 日志文件路径，默认 os.TempDir()
			MaxSize:    config.MaxSize,    // 每个日志文件保存10M，默认 100M
			MaxBackups: config.MaxBackups, // 保留30个备份，默认不限
			MaxAge:     config.MaxAge,     // 保留7天，默认不限
			Compress:   config.Compress,   // 是否压缩，默认不压缩
		}
		fileWriteSyncer := zapcore.AddSync(&hook)
		//consoleWriteSyncer := zapcore.AddSync(os.Stdout)
		//multiWriteSyncer := zapcore.NewMultiWriteSyncer(fileWriteSyncer, consoleWriteSyncer)
		// 设置日志级别
		// debug 可以打印出 info debug warn
		// info  级别可以打印 warn info
		// warn  只能打印 warn
		// debug->info->warn->error
		var level zapcore.Level
		switch config.LogLevel {
		case "debug":
			level = zap.DebugLevel
		case "info":
			level = zap.InfoLevel
		case "error":
			level = zap.ErrorLevel
		default:
			level = zap.InfoLevel
		}
		encoderConfig := zapcore.EncoderConfig{
			TimeKey:        "time",
			LevelKey:       "level",
			NameKey:        "logger",
			CallerKey:      "linenum",
			MessageKey:     "msg",
			StacktraceKey:  "stacktrace",
			LineEnding:     zapcore.DefaultLineEnding,
			EncodeLevel:    zapcore.LowercaseLevelEncoder,  // 小写编码器
			EncodeTime:     zapcore.ISO8601TimeEncoder,     // ISO8601 UTC 时间格式
			EncodeDuration: zapcore.SecondsDurationEncoder, //
			EncodeCaller:   zapcore.FullCallerEncoder,      // 全路径编码器
			EncodeName:     zapcore.FullNameEncoder,
		}
		// 设置日志级别
		atomicLevel := zap.NewAtomicLevel()
		atomicLevel.SetLevel(level)
		writer := fileWriteSyncer
		if config.IsStdout == true {
			writer = zapcore.NewMultiWriteSyncer(zapcore.AddSync(os.Stdout), zapcore.AddSync(fileWriteSyncer))
		}
		core := zapcore.NewCore(
			// zapcore.NewConsoleEncoder(encoderConfig),
			zapcore.NewJSONEncoder(encoderConfig),
			//zapcore.NewMultiWriteSyncer(zapcore.AddSync(os.Stdout), zapcore.AddSync(write)), // 打印到控制台和文件
			//multiWriteSyncer,
			writer,
			level,
		)
		// 开启开发模式，堆栈跟踪
		caller := zap.AddCaller()
		// 开启文件及行号
		development := zap.Development()
		// 设置初始化字段,如：添加一个服务器名称
		//filed := zap.Fields(zap.String("V5Live", "backend"))
		// 构造日志
		//Logger = zap.New(core, caller, development, filed)
		Logger = zap.New(core, caller, development)
		Logger.Info("Live backend Logger init success")
	})
}

func WithTraceID(traceID string) *zap.Logger {
	if traceID != "" {
		return Logger.With(zap.String("trace_id", traceID))
	}
	return Logger
}
