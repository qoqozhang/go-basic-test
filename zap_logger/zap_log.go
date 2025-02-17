package zap_logger

import (
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

var logger *zap.Logger

// 运行时动态修改日志级别的全局级别控制
var level = zap.NewAtomicLevel()

func init() {
	// 创建基础的RotateFileLog
	rotateLog := &RotateFileLog{
		FilePrefix: "user",
		FilePath:   "",
		MaxAge:     1,
	}

	// 创建zap的WriteSyncer
	writeSyncer := zapcore.AddSync(rotateLog)

	// 设置日志编码器
	encoderConfig := zapcore.EncoderConfig{
		TimeKey:        "time",
		LevelKey:       "level",
		NameKey:        "logger",
		CallerKey:      "caller",
		MessageKey:     "msg",
		StacktraceKey:  "stacktrace",
		LineEnding:     zapcore.DefaultLineEnding,
		EncodeLevel:    zapcore.CapitalLevelEncoder,                        // 使用大写字母记录日志级别
		EncodeTime:     zapcore.TimeEncoderOfLayout("2006-01-02 15:04:05"), // 更易读的时间格式
		EncodeDuration: zapcore.StringDurationEncoder,                      // 更易读的持续时间格式
		EncodeCaller:   zapcore.ShortCallerEncoder,
	}

	// 创建核心配置
	core := zapcore.NewCore(
		zapcore.NewConsoleEncoder(encoderConfig), // 使用Console编码器替代JSON编码器
		writeSyncer,
		level,
	)
	level.SetLevel(zap.InfoLevel)

	// 创建logger
	logger = zap.New(core, zap.AddCaller())
}
