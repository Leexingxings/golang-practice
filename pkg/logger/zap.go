package logger

import (
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"gopkg.in/natefinch/lumberjack.v2"
	"os"
	"path/filepath"
	"time"
)

var sugarLogger *zap.SugaredLogger

func init() {
	writeSyncer := getLogWriter()
	encoder := getEncoder()
	core := zapcore.NewCore(encoder, writeSyncer, zapcore.DebugLevel)

	// AddCaller() 日志调用方信息
	sugarLogger = zap.New(core, zap.AddCaller()).Sugar()
}

// getEncoder 日志编码格式，支持普通文本
func getEncoder() zapcore.Encoder {
	encoderConfig := zapcore.EncoderConfig{
		TimeKey:       "time",
		LevelKey:      "level",
		NameKey:       "logger",
		CallerKey:     "caller",
		FunctionKey:   zapcore.OmitKey,
		MessageKey:    "msg",
		StacktraceKey: "stacktrace",
		LineEnding:    zapcore.DefaultLineEnding,
		EncodeLevel: func(level zapcore.Level, encoder zapcore.PrimitiveArrayEncoder) {
			encoder.AppendString("[" + level.String() + "]") // 日志等级序列为小写字符串，如:InfoLevel被序列化为 "info"
		},
		EncodeTime: func(t time.Time, encoder zapcore.PrimitiveArrayEncoder) {
			encoder.AppendString(t.Format(time.DateTime))
		},
		EncodeDuration: zapcore.SecondsDurationEncoder, // 时间序列化，Duration为经过的浮点秒数
		EncodeCaller: func(caller zapcore.EntryCaller, encoder zapcore.PrimitiveArrayEncoder) {
			encoder.AppendString(caller.TrimmedPath()) // 日志行号显示
		},
	}

	return zapcore.NewConsoleEncoder(encoderConfig)
}

func getLogWriter() zapcore.WriteSyncer {
	// file, _ := os.OpenFile("./test.log", os.O_CREATE|os.O_APPEND|os.O_RDWR, 0744)

	separator := string(filepath.Separator)
	rootDir, _ := os.Getwd()

	logFIlePath := rootDir + separator + "logs" + separator + "practice.log"

	lumberJackLogger := &lumberjack.Logger{
		Filename:   logFIlePath,
		MaxSize:    10,    // 大小，单位：MB
		MaxBackups: 5,     // 最大备份数量
		MaxAge:     30,    // 最大备份天数
		Compress:   false, // 是否压缩
	}

	var writes = []zapcore.WriteSyncer{zapcore.AddSync(lumberJackLogger), zapcore.AddSync(os.Stdout)}

	return zapcore.NewMultiWriteSyncer(writes...)
}

func Debug(args ...any) {
	sugarLogger.Debug(args...)
}

func Debugf(template string, args ...any) {
	sugarLogger.Debugf(template, args...)
}

func Info(args ...any) {
	sugarLogger.Info(args...)
}

func Infof(template string, args ...any) {
	sugarLogger.Infof(template, args...)
}

func Warn(args ...any) {
	sugarLogger.Warn(args...)
}

func Warnf(template string, args ...any) {
	sugarLogger.Warnf(template, args...)
}

func Error(args ...any) {
	sugarLogger.Error(args...)
}

func Errorf(template string, args ...any) {
	sugarLogger.Errorf(template, args...)
}

func Panic(args ...any) {
	sugarLogger.Panic(args...)
}

func Panicf(template string, args ...any) {
	sugarLogger.Panicf(template, args...)
}
