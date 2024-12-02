package logger

import (
	"io"
	"os"
	"time"

	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"gopkg.in/natefinch/lumberjack.v2"
)

func Init() {
	core := zapcore.NewCore(getEncoder(), getRollingLogWriter(), zapcore.DebugLevel)
	logger := zap.New(core, zap.AddCaller())
	zap.ReplaceGlobals(logger)
}

func getRollingLogWriter() zapcore.WriteSyncer {
	return zapcore.AddSync(GetLogMultiWriter())
}

func GetLogMultiWriter() io.Writer {
	//conf := config.Conf.LogConfig
	// 归档名xxx-2024-09-13T15-52-45.302.log.gz
	// 归档名xxx-2024-09-13T15-52-45.302.log
	l := &lumberjack.Logger{
		Filename:   "app.log",
		MaxSize:    200, // megabytes
		MaxBackups: 3,
		MaxAge:     30,    // days
		Compress:   false, // 压缩归档
	}
	ws := io.MultiWriter(l, os.Stdout)
	return ws
}

func getEncoder() zapcore.Encoder {
	//return zapcore.NewConsoleEncoder(zap.NewProductionEncoderConfig())

	//TimeKey:        "ts",
	//LevelKey:       "level",
	//NameKey:        "logger",
	//CallerKey:      "caller",
	//FunctionKey:    zapcore.OmitKey,
	//MessageKey:     "msg",
	//StacktraceKey:  "stacktrace",
	//LineEnding:     zapcore.DefaultLineEnding,
	//EncodeLevel:    zapcore.LowercaseLevelEncoder,
	//EncodeTime:     zapcore.EpochTimeEncoder,
	//EncodeDuration: zapcore.SecondsDurationEncoder,
	//EncodeCaller:   zapcore.ShortCallerEncoder,
	encoderConfig := zap.NewProductionEncoderConfig()
	//encoderConfig.EncodeTime = zapcore.ISO8601TimeEncoder
	encoderConfig.EncodeTime = zapcore.TimeEncoderOfLayout(time.DateTime)
	return zapcore.NewJSONEncoder(encoderConfig)
}
