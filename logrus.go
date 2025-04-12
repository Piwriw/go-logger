package logger

import (
	"fmt"
	"io"
	"os"
	"path"
	"runtime"
	"time"

	"github.com/sirupsen/logrus"
)

var defaultCallerPrettyfierFunc = func(f *runtime.Frame) (string, string) {
	// 自定义文件名 + 行号格式
	filename := path.Base(f.File)
	return "", fmt.Sprintf("%s:%d", filename, f.Line)
}

// logrusLogger 实现

type logrusLogger struct {
	logger      *logrus.Logger
	errorLogger *logrus.Logger
	maskLogger  *MaskProcessor
	level       Level
	fields      logrus.Fields
	colorScheme *ColorScheme
	AddSource   bool
}

type customTextFormatter struct {
	logrus.TextFormatter
	location *time.Location
}

func (f *customTextFormatter) Format(entry *logrus.Entry) ([]byte, error) {
	entry.Time = entry.Time.In(f.location)
	return f.TextFormatter.Format(entry)
}

type customJSONFormatter struct {
	logrus.JSONFormatter
	location *time.Location
}

func (f *customJSONFormatter) Format(entry *logrus.Entry) ([]byte, error) {
	entry.Time = entry.Time.In(f.location)
	return f.JSONFormatter.Format(entry)
}

var _ Logger = (*logrusLogger)(nil)

func newLogrusLogger(opts Options) (Logger, error) {
	logger := logrus.New()
	errorLogger := logrus.New()
	location, err := time.LoadLocation(opts.TimeZone)
	if err != nil {
		return nil, err
	}
	if opts.AddSource {
		logger.SetReportCaller(true)
		errorLogger.SetReportCaller(true)
	}

	if opts.JSONFormat {
		customFmt := &customJSONFormatter{
			JSONFormatter: logrus.JSONFormatter{
				CallerPrettyfier: defaultCallerPrettyfierFunc,
				TimestampFormat:  opts.TimeFormat,
			},
			location: location,
		}

		logger.SetFormatter(customFmt)
		errorLogger.SetFormatter(customFmt)
	} else {
		customFmt := &customTextFormatter{
			TextFormatter: logrus.TextFormatter{
				CallerPrettyfier: defaultCallerPrettyfierFunc,
				TimestampFormat:  opts.TimeFormat,
			},
			location: location,
		}
		logger.SetFormatter(customFmt)
		errorLogger.SetFormatter(customFmt)
	}
	logger.SetLevel(ToLogrusLoggerLevel(opts.Level))
	errorLogger.SetLevel(ToLogrusLoggerLevel(ErrorLevel))
	// 设置控制台和文件输出
	multiWriter := io.MultiWriter(os.Stdout, getOutput(opts.FilePath))
	logger.SetOutput(multiWriter)
	errorLogger.SetOutput(getOutput(opts.ErrorOutput))
	logrusLogger := &logrusLogger{
		logger:      logger,
		errorLogger: errorLogger,
		level:       opts.Level,
	}
	// 设置颜色输出
	if opts.ColorEnabled {
		logrusLogger.colorScheme = opts.ColorScheme
	}
	if opts.AddSource {
		logrusLogger.AddSource = true
	}
	if opts.MaskEnable {
		logrusLogger.maskLogger = NewMaskProcessor(opts.maskRules...)
	}
	return logrusLogger, nil
}

func ToLogrusLoggerLevel(level Level) logrus.Level {
	switch level {
	case DebugLevel:
		return logrus.DebugLevel
	case InfoLevel:
		return logrus.InfoLevel
	case WarnLevel:
		return logrus.WarnLevel
	case ErrorLevel:
		return logrus.ErrorLevel
	case FatalLevel:
		return logrus.FatalLevel
	default:
		return logrus.InfoLevel
	}
}

func FromLogrusLoggerLevel(level logrus.Level) Level {
	switch level {
	case logrus.DebugLevel:
		return DebugLevel
	case logrus.InfoLevel:
		return InfoLevel
	case logrus.WarnLevel:
		return WarnLevel
	case logrus.ErrorLevel:
		return ErrorLevel
	case logrus.FatalLevel:
		return FatalLevel
	default:
		return defaultLevel
	}
}

func (l *logrusLogger) log(level logrus.Level, msg string, args ...any) {
	// 创建基础 fields
	fields := make(logrus.Fields)

	// 添加颜色格式化
	if l.colorScheme != nil {
		msg = l.colorScheme.Colorize(FromLogrusLoggerLevel(level), msg)
	}

	// 添加固定字段
	if l.fields != nil {
		for k, v := range l.fields {
			fields[k] = v
		}
	}

	// 添加调用源信息
	if l.AddSource {
		fields["source"] = getCaller(3)
	}

	// 处理 KV 参数
	if len(args) > 0 {
		// 确保参数是偶数个
		if len(args)%2 != 0 {
			args = append(args, "MISSING_VALUE")
		}

		// 将 KV 参数转换为 fields
		for i := 0; i < len(args); i += 2 {
			if key, ok := args[i].(string); ok {
				// 如果有脱敏处理器，先处理值
				if l.maskLogger != nil {
					fields[key] = l.maskLogger.Process(args[i], args[i+1])[1]
				} else {
					fields[key] = args[i+1]
				}
			}
		}
	}

	// 记录主日志
	entry := l.logger.WithFields(fields)
	entry.Log(level, msg)

	// 记录错误日志
	if l.errorLogger != nil && level >= logrus.ErrorLevel {
		errorFields := make(logrus.Fields)
		for k, v := range fields {
			errorFields[k] = v
		}
		// 确保错误日志有源信息
		if _, exists := errorFields["source"]; !exists {
			errorFields["source"] = getCaller(3)
		}
		l.errorLogger.WithFields(errorFields).Log(level, msg)
	}
}

func (l *logrusLogger) Debug(msg string, args ...any) { l.log(logrus.DebugLevel, msg, args...) }
func (l *logrusLogger) Debugf(format string, args ...any) {
	l.log(logrus.DebugLevel, fmt.Sprintf(format, args...))
}
func (l *logrusLogger) Info(msg string, args ...any) { l.log(logrus.InfoLevel, msg, args...) }
func (l *logrusLogger) Infof(format string, args ...any) {
	l.log(logrus.InfoLevel, fmt.Sprintf(format, args...))
}
func (l *logrusLogger) Warn(msg string, args ...any) {
	l.log(logrus.WarnLevel, msg, args...)
}
func (l *logrusLogger) Warnf(format string, args ...any) {
	l.log(logrus.WarnLevel, fmt.Sprintf(format, args...))
}
func (l *logrusLogger) Error(msg string, args ...any) { l.log(logrus.ErrorLevel, msg, args...) }
func (l *logrusLogger) Errorf(format string, args ...any) {
	l.log(logrus.ErrorLevel, fmt.Sprintf(format, args...))
}
func (l *logrusLogger) Fatal(msg string, args ...any) { l.log(logrus.FatalLevel, msg, args...) }
func (l *logrusLogger) Fatalf(format string, args ...any) {
	l.log(logrus.FatalLevel, fmt.Sprintf(format, args...))
}
func (l *logrusLogger) SetLevel(level Level) { l.logger.SetLevel(logrus.Level(level)) }

func (l *logrusLogger) WithFields(fields map[string]any) Logger {
	return &logrusLogger{
		logger: l.logger,
		level:  l.level,
		fields: logrus.Fields(fields),
	}
}
