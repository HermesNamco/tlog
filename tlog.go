package tlog

import (
	"github.com/natefinch/lumberjack"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

var Logger *zap.SugaredLogger

func init() {
	encoder := zapcore.NewConsoleEncoder(zapcore.EncoderConfig{
		MessageKey:     "msg",
		LevelKey:       "level",
		NameKey:        "name",
		TimeKey:        "ts",
		CallerKey:      "caller",
		FunctionKey:    "func",
		StacktraceKey:  "stacktrace",
		LineEnding:     "\n",
		EncodeTime:     zapcore.EpochTimeEncoder,
		EncodeLevel:    zapcore.LowercaseLevelEncoder,
		EncodeDuration: zapcore.SecondsDurationEncoder,
		EncodeCaller:   zapcore.ShortCallerEncoder,
	})
	core := zapcore.NewCore(encoder, getLogWriter(), zapcore.DebugLevel)
	logger := zap.New(core, zap.AddCaller())
	Logger = logger.Sugar()
}

func getLogWriter() zapcore.WriteSyncer {
	return zapcore.AddSync(&lumberjack.Logger{
		Filename:   "./myLog",
		MaxSize:    10,
		MaxAge:     7,
		MaxBackups: 10,
		Compress:   false,
	})
}

func Infof(template string, args ...interface{}) {
	Logger.Infof(template, args...)
}

func Warnf(template string, args ...interface{}) {
	Logger.Warnf(template, args...)
}

func Errorf(template string, args ...interface{}) {
	Logger.Errorf(template, args...)
}
