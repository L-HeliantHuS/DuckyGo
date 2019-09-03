package util

import (
	"fmt"
	"os"
	"time"
)

// 自增iota, 依次为 0 1 2 3
const (
	// LevelError 错误
	LevelError = iota
	// LevelWarning 警告
	LevelWarning
	// LevelInformational 提示
	LevelInformational
	// LevelDebug 除错
	LevelDebug
)

var logger *Logger

// Logger 日志
type Logger struct {
	level int
}

func (ll *Logger) saveLog(log string, level int) error {
	logLevelMap := map[int]string{
		LevelError:         "ERROR",
		LevelWarning:       "WARNING",
		LevelInformational: "INFO",
		LevelDebug:         "DEBUG",
	}
	filename := fmt.Sprintf("log/%s.log", logLevelMap[level])
	file, err := os.OpenFile(filename, os.O_WRONLY, 0644)
	if err != nil {
		panic(err)
	}
	n, _ := file.Seek(0, 2)
	_, err = file.WriteAt([]byte(log), n)
	return err
}

// Println 打印
func (ll *Logger) Println(msg string, level int) {
	s := fmt.Sprintf("%s %s \n", time.Now().Format("2006-01-02 15:04:05 -0700"), msg)
	_ = ll.saveLog(s, level)
	fmt.Println(s)
}

// Panic 极端错误
func (ll *Logger) Panic(format string) {
	if LevelError > ll.level {
		return
	}
	msg := fmt.Sprintf("[Panic] " + format)
	ll.Println(msg, LevelError)
	os.Exit(0)
}

// Error 错误
func (ll *Logger) Error(format string) {
	if LevelError > ll.level {
		return
	}
	msg := fmt.Sprintf("[E] " + format)
	ll.Println(msg, LevelError)
}

// Warning 警告
func (ll *Logger) Warning(format string) {
	if LevelWarning > ll.level {
		return
	}
	msg := fmt.Sprintf("[W] " + format)
	ll.Println(msg, LevelWarning)
}

// Info 信息
func (ll *Logger) Info(format string) {
	if LevelInformational > ll.level {
		return
	}
	msg := fmt.Sprintf("[I] " + format)
	ll.Println(msg, LevelInformational)
}

// Debug 校验
func (ll *Logger) Debug(format string) {
	if LevelDebug > ll.level {
		return
	}
	msg := fmt.Sprintf("[D] " + format)
	ll.Println(msg, LevelDebug)
}

// BuildLogger 构建logger
func BuildLogger(level string) {
	intLevel := LevelError
	switch level {
	case "ERROR":
		intLevel = LevelError
	case "WARNING":
		intLevel = LevelWarning
	case "INFO":
		intLevel = LevelInformational
	case "DEBUG":
		intLevel = LevelDebug
	}
	l := Logger{
		level: intLevel,
	}
	logger = &l
}

// Log 返回日志对象
func Log() *Logger {
	if logger == nil {
		l := Logger{
			level: LevelDebug,
		}
		logger = &l
	}
	return logger
}
