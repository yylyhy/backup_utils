package logger

import (
	"io"
	"os"

	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"gopkg.in/natefinch/lumberjack.v2"
)

var (
	sugarLogger *zap.SugaredLogger
	logFileName = "test.log"
	maxSize     = 1
	maxBackups  = 5
	maxAge      = 30
	logLevel    = zapcore.DebugLevel
)

func init() {
	writeSyncer := getLogWriter()
	encoder := getEncoder()
	core := zapcore.NewCore(encoder, writeSyncer, logLevel)

	logger := zap.New(core, zap.AddCaller(), zap.AddCallerSkip(1))
	sugarLogger = logger.Sugar()
}

func getEncoder() zapcore.Encoder {
	encoderConfig := zap.NewProductionEncoderConfig()
	encoderConfig.EncodeTime = zapcore.ISO8601TimeEncoder
	encoderConfig.EncodeLevel = zapcore.CapitalLevelEncoder
	return zapcore.NewConsoleEncoder(encoderConfig)
}

func getLogWriter() zapcore.WriteSyncer {
	lumberJackLogger := &lumberjack.Logger{
		Filename:   logFileName,
		MaxSize:    maxSize,
		MaxBackups: maxBackups,
		MaxAge:     maxAge,
		Compress:   false,
		LocalTime:  true,
	}
	ws := io.MultiWriter(lumberJackLogger, os.Stdout)
	return zapcore.AddSync(ws)
}

func Debug(args ...interface{}) {
	sugarLogger.Debug(args...)
}
func Debugf(template string, args ...interface{}) {
	sugarLogger.Debugf(template, args...)
}
func Info(args ...interface{}) {
	sugarLogger.Info(args...)
}
func Infof(template string, args ...interface{}) {
	sugarLogger.Infof(template, args...)
}
func Warn(args ...interface{}) {
	sugarLogger.Warn(args...)
}
func Warnf(template string, args ...interface{}) {
	sugarLogger.Warnf(template, args...)
}
func Error(args ...interface{}) {
	sugarLogger.Error(args...)
}
func Errorf(template string, args ...interface{}) {
	sugarLogger.Errorf(template, args...)
}
func DPanic(args ...interface{}) {
	sugarLogger.DPanic(args...)
}
func DPanicf(template string, args ...interface{}) {
	sugarLogger.DPanicf(template, args...)
}
func Panic(args ...interface{}) {
	sugarLogger.Panic(args...)
}
func Panicf(template string, args ...interface{}) {
	sugarLogger.Panicf(template, args...)
}
func Fatal(args ...interface{}) {
	sugarLogger.Fatal(args...)
}
func Fatalf(template string, args ...interface{}) {
	sugarLogger.Fatalf(template, args...)
}
func Sync() {
	sugarLogger.Sync()
}
