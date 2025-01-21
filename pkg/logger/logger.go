/**
 * @Author: LiuShuXin
 * @Description:
 * @File:  logger
 * Software: Goland
 * @Date: 2025/1/21 9:43
 */

package logger

import (
	"app-server/settings"
	"encoding/json"
	"fmt"
	"os"
	"time"

	"github.com/natefinch/lumberjack"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

type Level int8

func (l Level) String() string {
	switch l {
	case LevelDebug:
		return "debug"
	case LevelInfo:
		return "info"
	case LevelWarn:
		return "warn"
	case LevelError:
		return "error"
	case LevelFatal:
		return "fatal"
	case LevelPanic:
		return "panic"
	}
	return ""
}

const (
	LevelDebug Level = iota
	LevelInfo
	LevelWarn
	LevelError
	LevelFatal
	LevelPanic
)

var lg *zap.Logger

// Init 初始化全局日志 todo
func Init(cfg *settings.LogConfig, mode string) (err error) {
	writeSyncer := getLogWriter(cfg.Filepath+"/"+cfg.Filename, cfg.MaxSize, cfg.MaxBackups, cfg.MaxAge) // 写入器
	encoder := getEncoder()                                                                             // 编码器
	var l = new(zapcore.Level)
	err = l.UnmarshalText([]byte(cfg.Level))
	if err != nil {
		return
	}
	var core zapcore.Core
	if mode == "debug" {
		// 进入开发模式，日志输出到终端
		consoleEncoder := zapcore.NewConsoleEncoder(zap.NewDevelopmentEncoderConfig())
		core = zapcore.NewTee(
			zapcore.NewCore(encoder, writeSyncer, l),
			zapcore.NewCore(consoleEncoder, zapcore.Lock(os.Stdout), zapcore.DebugLevel),
		)
	} else {
		core = zapcore.NewCore(encoder, writeSyncer, l)
	}

	lg = zap.New(core, zap.AddCaller())

	zap.ReplaceGlobals(lg)
	zap.L().Info("init logger success")
	return
}

func getEncoder() zapcore.Encoder {
	encoderConfig := zap.NewProductionEncoderConfig()
	encoderConfig.EncodeTime = zapcore.ISO8601TimeEncoder
	encoderConfig.TimeKey = "time"
	encoderConfig.EncodeLevel = zapcore.CapitalLevelEncoder
	encoderConfig.EncodeDuration = zapcore.SecondsDurationEncoder
	encoderConfig.EncodeCaller = zapcore.ShortCallerEncoder
	return zapcore.NewJSONEncoder(encoderConfig)
}

func getLogWriter(filename string, maxSize, maxBackup, maxAge int) zapcore.WriteSyncer {
	lumberJackLogger := &lumberjack.Logger{
		Filename:   filename,
		MaxSize:    maxSize,
		MaxBackups: maxBackup,
		MaxAge:     maxAge,
	}
	return zapcore.AddSync(lumberJackLogger)
}

//-----------------------------------------------

type Fields map[string]interface{}

type Logger struct {
	zl *zap.Logger
	//ctx       context.Context
	fields  Fields
	callers []string
}

// GetDayLogger 按天拆分的日志 todo
//func GetDayLogger(name string, numDay int) *Logger {
//	// 创建日志写入器
//	lumberJackLogger := &lumberjack.Logger{
//		Filename:   "storage/logs" + "/" + name + ".log",
//		MaxSize:    settings.GetConf().MaxSize, // MB
//		MaxBackups: settings.GetConf().MaxBackups,
//		MaxAge:     numDay,
//		Compress:   true, // 是否压缩
//		LocalTime:  true,
//	}
//	writeSyncer := zapcore.AddSync(lumberJackLogger)
//
//	// 创建日志编码器
//	encoderConfig := zap.NewProductionEncoderConfig()
//	encoderConfig.EncodeTime = zapcore.ISO8601TimeEncoder
//	encoderConfig.TimeKey = "time"
//	encoderConfig.EncodeLevel = zapcore.CapitalLevelEncoder
//	encoderConfig.EncodeDuration = zapcore.SecondsDurationEncoder
//	encoderConfig.EncodeCaller = zapcore.ShortCallerEncoder
//	encoder := zapcore.NewJSONEncoder(encoderConfig)
//
//	// 创建日志核心
//	core := zapcore.NewCore(encoder, writeSyncer, zapcore.DebugLevel)
//
//	// 创建日志记录器
//	logger := zap.New(core)
//
//	l := &Logger{
//		zl: logger,
//	}
//
//	return l
//}

func (l *Logger) JSONFormat(level Level, message string) map[string]interface{} {
	data := make(Fields, len(l.fields)+4)
	data["level"] = level.String()
	data["time"] = time.Now().Local().UnixNano()
	data["message"] = message
	data["callers"] = l.callers
	if len(l.fields) > 0 {
		for k, v := range l.fields {
			if _, ok := data[k]; !ok {
				data[k] = v
			}
		}
	}

	return data
}

func (l *Logger) Output(level Level, message string) {
	//body, _ := json.Marshal(l.JSONFormat(level, message))
	body, _ := json.Marshal(message)
	content := string(body)
	switch level {
	case LevelDebug:
		l.zl.Debug(content)
	case LevelInfo:
		l.zl.Info(content)
	case LevelWarn:
		l.zl.Warn(content)
	case LevelError:
		l.zl.Error(content)
	case LevelFatal:
		l.zl.Fatal(content)
	case LevelPanic:
		l.zl.Panic(content)
	}
}

func (l *Logger) Info(v ...interface{}) {
	l.Output(LevelInfo, fmt.Sprint(v...))
}

func (l *Logger) Infof(format string, v ...interface{}) {
	l.Output(LevelInfo, fmt.Sprintf(format, v...))
}

func (l *Logger) Fatal(v ...interface{}) {
	l.Output(LevelFatal, fmt.Sprint(v...))
}

func (l *Logger) Fatalf(format string, v ...interface{}) {
	l.Output(LevelFatal, fmt.Sprintf(format, v...))
}

func (l *Logger) Debug(v ...interface{}) {
	l.Output(LevelDebug, fmt.Sprint(v...))
}

func (l *Logger) Debugf(format string, v ...interface{}) {
	l.Output(LevelDebug, fmt.Sprintf(format, v...))
}

func (l *Logger) Warn(v ...interface{}) {
	l.Output(LevelWarn, fmt.Sprint(v...))
}

func (l *Logger) Warnf(format string, v ...interface{}) {
	l.Output(LevelWarn, fmt.Sprintf(format, v...))
}

func (l *Logger) Error(v ...interface{}) {
	l.Output(LevelError, fmt.Sprint(v...))
}

func (l *Logger) Errorf(format string, v ...interface{}) {
	l.Output(LevelError, fmt.Sprintf(format, v...))
}

func (l *Logger) Panic(v ...interface{}) {
	l.Output(LevelPanic, fmt.Sprint(v...))
}

func (l *Logger) Panicf(format string, v ...interface{}) {
	l.Output(LevelPanic, fmt.Sprintf(format, v...))
}
