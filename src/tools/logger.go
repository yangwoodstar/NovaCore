package tools

import (
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"gopkg.in/natefinch/lumberjack.v2"
	"sync"
)

var (
	Logger *zap.Logger
	once   sync.Once
)

func InitLogger(logPath string, loglevel string) {
	// 日志分割
	once.Do(func() {
		hook := lumberjack.Logger{
			Filename:   logPath, // 日志文件路径，默认 os.TempDir()
			MaxSize:    100,     // 每个日志文件保存10M，默认 100M
			MaxBackups: 10,      // 保留30个备份，默认不限
			MaxAge:     7,       // 保留7天，默认不限
			Compress:   false,   // 是否压缩，默认不压缩
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
		switch loglevel {
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
		core := zapcore.NewCore(
			// zapcore.NewConsoleEncoder(encoderConfig),
			zapcore.NewJSONEncoder(encoderConfig),
			//zapcore.NewMultiWriteSyncer(zapcore.AddSync(os.Stdout), zapcore.AddSync(write)), // 打印到控制台和文件
			//multiWriteSyncer,
			fileWriteSyncer,
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
