package logger

import (
	"flag"
	"fmt"
	"io"
	"os"
	"time"

	"k8s.io/klog/v2"
)

type klogLogger struct {
	level       Level
	filePath    string
	timeZone    *time.Location
	addSource   bool
	colorScheme *ColorScheme
	errorOutput string
	maskLogger  *MaskProcessor
}

var _ Logger = (*klogLogger)(nil)

func newKlogLogger(opts Options) (Logger, error) {
	klog.InitFlags(flag.CommandLine)
	var ioWriters []io.Writer

	location, err := time.LoadLocation(opts.TimeZone)
	if err != nil {
		return nil, err
	}
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
	klogLogger := &klogLogger{
		level:       opts.Level,
		filePath:    opts.FilePath,
		addSource:   opts.AddSource,
		errorOutput: opts.ErrorOutput,
		colorScheme: opts.ColorScheme,
		timeZone:    location,
	}
	if opts.MaskEnable {
		klogLogger.maskLogger = NewMaskProcessor(opts.maskRules...)
	}
	return klogLogger, nil
}

func (l *klogLogger) log(level Level, msg string, args ...any) {
	defer klog.Flush()
	if l.level > level {
		return
	}

	// 追加时间戳到日志消息
	timestamp := time.Now().In(l.timeZone).Format(time.DateTime)
	msg = fmt.Sprintf("[%s] %s", timestamp, msg)

	if l.colorScheme != nil {
		msg = l.colorScheme.Colorize(level, msg)
	}
	if l.maskLogger != nil {
		args = l.maskLogger.Process(args...)
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
}

func (l *klogLogger) Debug(msg string, args ...any) {
	l.log(DebugLevel, msg, args...)
}

func (l *klogLogger) Debugf(format string, args ...any) {
	l.log(DebugLevel, fmt.Sprintf(format, args...))
}

func (l *klogLogger) Info(msg string, args ...any) {
	l.log(InfoLevel, msg, args...)
}

func (l *klogLogger) Infof(format string, args ...any) {
	l.log(InfoLevel, fmt.Sprintf(format, args...))
}

func (l *klogLogger) Warn(msg string, args ...any) {
	l.log(WarnLevel, msg, args...)
}

func (l *klogLogger) Warnf(format string, args ...any) {
	l.log(WarnLevel, fmt.Sprintf(format, args...))
}

func (l *klogLogger) Error(msg string, args ...any) {
	l.log(ErrorLevel, msg, args...)
}

func (l *klogLogger) Errorf(format string, args ...any) {
	l.log(ErrorLevel, fmt.Sprintf(format, args...))
}

func (l *klogLogger) Fatal(msg string, args ...any) {
	l.log(ErrorLevel, msg, args...)
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
