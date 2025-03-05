/**
 * @Author: LiuShuXin
 * @Description:
 * @File:  logrus_test
 * Software: Goland
 * @Date: 2025/3/5 16:41
 */

package logger

import (
	"fmt"
	"testing"
)

//
//// Logger 自定义日志对象（每个实例对应一个日志文件）
//type DayLogger struct {
//	entry *logrus.Entry
//}
//
//// NewDayLogger 创建独立日志实例
//func NewDayLogger(logName string) *DayLogger {
//	log := logrus.New()
//
//	// 设置日志级别
//	level, err := logrus.ParseLevel(settings.GetConf().Level)
//	if err != nil {
//		level = logrus.InfoLevel // 默认 info 级别
//	}
//	log.SetLevel(level)
//
//	// 输出到 终端 及 文件中
//	stderr := os.Stdout
//	logFile := &lumberjack.Logger{
//		Filename:   filepath.Join(LOG_DIR, fmt.Sprintf("%s.log", logName)),
//		MaxSize:    settings.GetConf().MaxSize,
//		MaxBackups: settings.GetConf().MaxBackups,
//		MaxAge:     settings.GetConf().MaxAge,
//	}
//	log.SetOutput(io.MultiWriter(stderr, logFile))
//
//	// 设置 JSON 格式 + 调用者信息
//	//log.SetFormatter(&logrus.JSONFormatter{
//	//	CallerPrettyfier: func(f *runtime.Frame) (string, string) {
//	//		filename := filepath.Base(f.File)
//	//		return "", fmt.Sprintf("%s:%d", filename, f.Line)
//	//	},
//	//})
//	//log.SetReportCaller(true) // 显示调用位置
//
//	log.SetFormatter(&logrus.TextFormatter{})
//
//	return &DayLogger{
//		entry: log.WithFields(logrus.Fields{}), // 初始无额外字段
//	}
//}
//
//// ----- 通用日志方法 -----
//func (dl *DayLogger) Info(args ...any) {
//	dl.entry.Info(args...)
//}
//
//func (dl *DayLogger) Infof(format string, args ...any) {
//	dl.entry.Infof(format, args...)
//}
//
//func (dl *DayLogger) Error(args ...any) {
//	dl.entry.Error(args...)
//}
//
//func (dl *DayLogger) Errorf(format string, args ...any) {
//	dl.entry.Errorf(format, args...)
//}
//
//// WithField 添加字段（返回新实例）
//func (dl *DayLogger) WithField(key string, value any) *DayLogger {
//	return &DayLogger{entry: dl.entry.WithField(key, value)}
//}
//
//// WithFields 添加多个字段（返回新实例）
//func (dl *DayLogger) WithFields(fields map[string]any) *DayLogger {
//	return &DayLogger{entry: dl.entry.WithFields(fields)}
//}
//
//// GetWriter 获取 io.Writer（用于 Gin 等框架的默认日志输出）
//func (dl *DayLogger) GetWriter() io.Writer {
//	return dl.entry.Writer()
//}

func TestDayLogger(t *testing.T) {
	log := NewDayLogger("test", 3)
	fmt.Println(log)
	//log.Info("hello world")
	//log.Infof("hello %s", "world")
}
