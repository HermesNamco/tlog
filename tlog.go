package tlog

import (
	"github.com/natefinch/lumberjack"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

var tlog *zap.SugaredLogger

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
		EncodeTime:     zapcore.ISO8601TimeEncoder,
		EncodeLevel:    zapcore.LowercaseLevelEncoder,
		EncodeDuration: zapcore.StringDurationEncoder,
		EncodeCaller:   zapcore.ShortCallerEncoder,
	})
	core := zapcore.NewCore(encoder, getLogWriter(), zapcore.InfoLevel)
	logger := zap.New(core, zap.AddCaller(), zap.AddCallerSkip(1))
	tlog = logger.Sugar()
}

func getLogWriter() zapcore.WriteSyncer {
	return zapcore.AddSync(&lumberjack.Logger{
		Filename:   "./tLog",
		MaxSize:    10,
		MaxAge:     7,
		MaxBackups: 10,
		Compress:   false,
	})
}

func Infof(template string, args ...interface{}) {
	tlog.Infof(template, args...)
}

func Warnf(template string, args ...interface{}) {
	tlog.Warnf(template, args...)
}

func Errorf(template string, args ...interface{}) {
	tlog.Errorf(template, args...)
}
