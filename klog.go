package logger

import (
	"flag"
	"fmt"
	"k8s.io/klog/v2"
	"os"
	"time"
)

type klogLogger struct {
	level       Level
	filePath    string
	timeZone    *time.Location
	addSource   bool
	colorScheme *ColorScheme
	errorOutput string
}

var _ Logger = (*klogLogger)(nil)

func newKlogLogger(opts Options) (Logger, error) {
	location, err := time.LoadLocation(opts.TimeZone)
	if err != nil {
		return nil, err
	}
	if opts.FilePath != "" {
		klog.SetOutput(getOutput(opts.FilePath))
	}
	klog.LogToStderr(false)
	if opts.ErrorOutput != "" {
		klog.SetOutputBySeverity("ERROR", getOutput(opts.ErrorOutput))
	}
	return &klogLogger{
		level:       opts.Level,
		filePath:    opts.FilePath,
		addSource:   opts.AddSource,
		errorOutput: opts.ErrorOutput,
		colorScheme: opts.ColorScheme,
		timeZone:    location,
	}, nil
}

func (l *klogLogger) log(level Level, msg string, args ...any) {
	if l.level > level {
		return
	}

	// 追加时间戳到日志消息
	timestamp := time.Now().In(l.timeZone).Format(time.DateTime)
	msg = fmt.Sprintf("[%s] %s", timestamp, msg)

	if l.colorScheme != nil {
		msg = l.colorScheme.Colorize(level, msg)
	}

	kvs := make([]any, 0, len(args)*2)
	for i := 0; i < len(args); i += 2 {
		if i+1 >= len(args) {
			break
		}
		key, ok := args[i].(string)
		if !ok {
			continue
		}
		kvs = append(kvs, key, args[i+1])
	}

	switch level {
	case DebugLevel:
		klog.V(5).Infof(msg, kvs...)
	case InfoLevel:
		klog.Infof(msg, kvs...)
	case WarnLevel:
		klog.Warningf(msg, kvs...)
	case ErrorLevel:
		klog.Errorf(msg, kvs...)
	default:
		klog.Infof(msg, kvs...)
	}
	defer klog.Flush()
}

func (l *klogLogger) Debug(args ...any) {
	l.log(DebugLevel, fmt.Sprint(args...))
}

func (l *klogLogger) Debugf(format string, args ...any) {
	l.log(DebugLevel, fmt.Sprintf(format, args...))
}

func (l *klogLogger) Info(args ...any) {
	l.log(InfoLevel, fmt.Sprint(args...))
}

func (l *klogLogger) Infof(format string, args ...any) {
	l.log(InfoLevel, fmt.Sprintf(format, args...))
}

func (l *klogLogger) Warn(args ...any) {
	l.log(WarnLevel, fmt.Sprint(args...))
}

func (l *klogLogger) Warnf(format string, args ...any) {
	l.log(WarnLevel, fmt.Sprintf(format, args...))
}

func (l *klogLogger) Error(args ...any) {
	l.log(ErrorLevel, fmt.Sprint(args...))
}

func (l *klogLogger) Errorf(format string, args ...any) {
	l.log(ErrorLevel, fmt.Sprintf(format, args...))
}

func (l *klogLogger) Fatal(args ...any) {
	l.log(ErrorLevel, fmt.Sprint(args...))
	os.Exit(1)
}

func (l *klogLogger) Fatalf(format string, args ...any) {
	l.log(ErrorLevel, fmt.Sprintf(format, args...))
	klog.Flush()
	os.Exit(1)
}

func (l *klogLogger) SetLevel(level Level) {
	l.level = level
	// klog级别需要通过flags重新设置
	fs := flag.NewFlagSet("klog", flag.ContinueOnError)
	klog.InitFlags(fs)
	switch level {
	case DebugLevel:
		fs.Set("v", "4")
	case InfoLevel:
		fs.Set("v", "2")
	case WarnLevel:
		fs.Set("v", "1")
	case ErrorLevel, FatalLevel:
		fs.Set("v", "0")
	}
}

func (l *klogLogger) WithFields(fields map[string]any) Logger {
	// klog不支持结构化日志，返回新实例但保留字段
	return &klogLogger{
		level:       l.level,
		filePath:    l.filePath,
		addSource:   l.addSource,
		colorScheme: l.colorScheme,
		errorOutput: l.errorOutput,
	}
}
