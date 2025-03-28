package logger

import (
	"flag"
	"fmt"
	"io"
	"k8s.io/klog/v2"
	"os"
)

type klogLogger struct {
	level       Level
	filePath    string
	addSource   bool
	colorScheme *ColorScheme
	errorOutput string
}

var _ Logger = (*klogLogger)(nil)

func newKlogLogger(opts Options) (Logger, error) {
	klog.InitFlags(flag.CommandLine)
	var ioWriters []io.Writer

	if opts.FilePath != "" {
		ioWriters = append(ioWriters, getOutput(opts.FilePath))
		klog.SetOutput(getOutput(opts.FilePath))
	}
	var logRotation *LogRotation
	// 设置日志轮转
	if opts.LogRotation != nil {
		logRotation = initLogRotation(opts.LogRotation.FilePath,
			opts.LogRotation.MaxSize,
			opts.LogRotation.MaxAge,
			opts.LogRotation.MaxBackups,
			opts.LogRotation.Compress)
		ioWriters = append(ioWriters, logRotation.logger)
		multiWriter := io.MultiWriter(ioWriters...)
		klog.SetOutput(multiWriter)
	}
	klog.LogToStderr(false)
	if opts.ErrorOutput != "" {
		if logRotation != nil {
			multiWriter := io.MultiWriter(getOutput(opts.ErrorOutput), logRotation.logger)
			klog.SetOutputBySeverity("ERROR", multiWriter)
		}
		klog.SetOutputBySeverity("ERROR", getOutput(opts.ErrorOutput))
	}
	if err := flag.CommandLine.Set("one_output", "true"); err != nil {
		return nil, err
	}

	if !flag.Parsed() {
		flag.Parse()
	}
	return &klogLogger{
		level:       opts.Level,
		filePath:    opts.FilePath,
		addSource:   opts.AddSource,
		errorOutput: opts.ErrorOutput,
		colorScheme: opts.ColorScheme,
	}, nil
}

func (l *klogLogger) log(level Level, msg string, args ...any) {
	// 设置日志级别
	if l.level > level {
		return
	}
	if l.colorScheme != nil {
		msg = l.colorScheme.Colorize(level, msg)
	}
	// 转换 args 为 klog 结构化日志格式
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

	// 根据日志级别调用不同的 klog 方法
	switch level {
	case DebugLevel:
		klog.V(5).InfoSDepth(2, msg, kvs...)
	case InfoLevel:
		if l.addSource {
			klog.InfoSDepth(2, msg, kvs...)
			break
		}
		klog.InfoS(msg, kvs...)
	case WarnLevel:
		if l.addSource {
			klog.WarningfDepth(2, "%s %v", msg, kvs) // Warningf 只能格式化
			break
		}
		klog.Warningf("%s %v", msg, kvs) // Warningf 只能格式化
	case ErrorLevel:
		if l.addSource {
			klog.ErrorfDepth(2, "%s %v", msg, kvs)
			break
		}
		klog.ErrorS(nil, msg, kvs...)
	default:
		if l.addSource {
			klog.InfoSDepth(2, msg, kvs...)
			break
		}
		klog.InfoS(msg, kvs...)
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
