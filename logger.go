package logger

import (
	"fmt"
	"time"
)

// Level 定义日志级别类型 / Define log level type
type Level int

const (
	defaultLevel       = InfoLevel         // 默认日志级别 / Default log level
	defaultJSONFormat  = false             // 默认非 JSON 格式 / Default is non-JSON format
	defaultAddSource   = false             // 默认不打印调用信息 / Default: no caller info
	defaultLogFile     = "./app.log"       // 默认日志文件路径 / Default log file path
	defaultTimeFormat  = time.DateTime     // 默认时间格式 / Default time format
	defaultErrorOutput = "./app_error.log" // 默认错误日志文件 / Default error log file
)

// 日志级别定义 / Log level definitions
const (
	DebugLevel Level = iota // 调试级别 / Debug level
	InfoLevel               // 信息级别 / Info level
	WarnLevel               // 警告级别 / Warning level
	ErrorLevel              // 错误级别 / Error level
	FatalLevel              // 致命错误级别 / Fatal level
)

// 常见时区定义 / Common timezone constants
const (
	// 亚洲时区 / Asia timezones
	CSTTime = "Asia/Shanghai" // 中国标准时间 / China Standard Time (UTC+8)
	JSTTime = "Asia/Tokyo"    // 日本时间 / Japan Standard Time (UTC+9)
	ISTTime = "Asia/Kolkata"  // 印度时间 / India Standard Time (UTC+5:30)
	KSTTime = "Asia/Seoul"    // 韩国时间 / Korea Standard Time (UTC+9)

	// 美洲时区 / Americas
	ESTTime   = "America/New_York"    // 美国东部时间 / Eastern Standard Time (UTC-5/UTC-4 DST)
	CSTUSTime = "America/Chicago"     // 美国中部时间 / Central Standard Time (UTC-6/UTC-5 DST)
	MSTTime   = "America/Denver"      // 美国山区时间 / Mountain Standard Time (UTC-7/UTC-6 DST)
	PSTTime   = "America/Los_Angeles" // 美国太平洋时间 / Pacific Standard Time (UTC-8/UTC-7 DST)
	BRTTime   = "America/Sao_Paulo"   // 巴西时间 / Brazil Time (UTC-3)

	// 欧洲时区 / Europe
	UTC     = "UTC"           // 协调世界时 / Coordinated Universal Time (UTC+0)
	GMT     = "Europe/London" // 英国时间 / Greenwich Mean Time (UTC+0, DST+1)
	CETTime = "Europe/Paris"  // 中欧时间 / Central European Time (UTC+1/DST+2)
	EETTime = "Europe/Athens" // 东欧时间 / Eastern European Time (UTC+2/DST+3)

	// 澳大利亚及大洋洲 / Oceania
	AESTTime = "Australia/Sydney" // 澳大利亚东部时间 / Australian Eastern Time (UTC+10/DST+11)
	NZTime   = "Pacific/Auckland" // 新西兰时间 / New Zealand Time (UTC+12/DST+13)

	// 非洲时区 / Africa
	SATime = "Africa/Johannesburg" // 南非时间 / South Africa Time (UTC+2)
	EATime = "Africa/Nairobi"      // 东非时间 / East Africa Time (UTC+3)
)

// Logger 接口定义
type Logger interface {
	Debug(msg string, args ...any)
	Debugf(format string, args ...any)

	Info(msg string, args ...any)
	Infof(format string, args ...any)

	Warn(msg string, args ...any)
	Warnf(format string, args ...any)

	Error(msg string, args ...any)
	Errorf(format string, args ...any)

	Fatal(msg string, args ...any)
	Fatalf(format string, args ...any)

	SetLevel(level Level)
	WithFields(fields map[string]any) Logger
}

// LoggerType defines the supported logger types
// LoggerType 定义支持的日志类型
type LoggerType string

const (
	SlogLogger   LoggerType = "slog"   // 默认使用slog / Default logger
	ZapLogger    LoggerType = "zap"    // Zap 日志 / Zap logger
	LogrusLogger LoggerType = "logrus" // Logrus 日志 / Logrus logger
	KlogLogger   LoggerType = "klog"   // Kubernetes 日志 / Kubernetes logger
)

// NewLogger creates a new logger instance with default type
// NewLogger 创建默认类型的日志实例（使用 slog）
func NewLogger(options ...Option) (Logger, error) {
	return NewLoggerWithType(SlogLogger, options...)
}

// DefaultLogger creates a default logger with pre-defined options
// DefaultLogger 创建默认配置的日志实例
func DefaultLogger() (Logger, error) {
	opts := applyOptions(
		WithAddSource(),
		WithLevel(defaultLevel),
		WithTimeFormat(defaultTimeFormat),
		WithFileOutput(defaultLogFile),
		WithErrorOutPut(defaultErrorOutput),
	)
	return newSlogLogger(opts)
}

// DefaultMaskHandler 创建默认的脱敏处理器
// DefaultMaskHandler creates a default mask handler
// 目前支持脱敏规则：
// 1. 密码脱敏
// 2. 手机号脱敏
func DefaultMaskHandler() []MaskHandler {
	return []MaskHandler{
		&PasswordMark{}, &PhoneMask{},
	}
}

// NewLoggerWithType creates a logger with specified type
// NewLoggerWithType 创建指定类型的日志实例
func NewLoggerWithType(loggerType LoggerType, options ...Option) (Logger, error) {
	opts := applyOptions(options...)

	switch loggerType {
	case SlogLogger:
		return newSlogLogger(opts)
	case ZapLogger:
		return newZapLogger(opts)
	case LogrusLogger:
		return newLogrusLogger(opts)
	case KlogLogger:
		return newKlogLogger(opts)
	default:
		return nil, fmt.Errorf("unknown logger type: %s", loggerType)
	}
}

// Option 是配置函数类型 / Option function type for configuration
type Option func(*Options)

// Options contains logger configuration
// Options 包含日志配置项
type Options struct {
	// Logging level
	// 设置日志级别
	Level Level
	// Use JSON format
	// 以JSON格式输出日志
	JSONFormat bool
	// Log file path
	// 设置日志文件路径
	FilePath string
	// Show caller information
	// 是否打印日志函数调用信息，默认不打印
	AddSource bool
	// Time format
	// 格式化日志打印时间
	TimeFormat string
	// Timezone
	// 时区
	TimeZone string
	// Error log output path
	// ErrorOutput 错误日志输出
	ErrorOutput string
	// Log rotation configuration
	// 日志轮转配置
	LogRotation *LogRotation
	// Enable color output
	// 设置颜色输出
	// 开启默认有一个默认的颜色方案，可通过 WithColorScheme 自定义颜色方案
	ColorEnabled bool
	// Custom color scheme
	// 主题颜色方案
	ColorScheme *ColorScheme
	//
	MaskEnable bool
	maskRules  []MaskHandler
	// 其他配置项...
}

// WithLevel sets the logging level
// WithLevel 设置日志级别
// 默认为 InfoLevel
func WithLevel(level Level) Option {
	return func(o *Options) {
		o.Level = level
	}
}

// WithJSONFormat enables JSON output format
// WithJSONFormat 以JSON格式输出日志
// Klog不支持
func WithJSONFormat() Option {
	return func(o *Options) {
		o.JSONFormat = true
	}
}

// WithFileOutput sets the log file path
// Default FilePath is ./app.log
// WithFileOutput 设置日志文件路径
// 默认为 ./app.log
func WithFileOutput(path string) Option {
	return func(o *Options) {
		if path == "" {
			path = defaultLogFile
		}
		o.FilePath = path
	}
}

// WithAddSource enables showing caller information
// WithAddSource 打印日志函数调用信息
// 默认为 false，不打印
// klog 不开启，会打印同一行日志，无法区分日志来源
func WithAddSource() Option {
	return func(o *Options) {
		o.AddSource = true
	}
}

// WithTimeFormat sets the time format
// WithTimeFormat 格式化日志打印时间
// 默认为 time.DateTime
func WithTimeFormat(format string) Option {
	return func(o *Options) {
		o.TimeFormat = format
	}
}

// WithTimeZone sets the timezone
// WithTimeZone 格式化日志打印时间
func WithTimeZone(timeZone string) Option {
	return func(o *Options) {
		o.TimeZone = timeZone
	}
}

// WithErrorOutPut sets the error log file path
// WithErrorOutPut 错误日志输出
// Zap 默认开启Error日志的堆栈打印，Logrus 默认关闭
func WithErrorOutPut(path string) Option {
	return func(o *Options) {
		if path == "" {
			path = defaultErrorOutput
		}
		o.ErrorOutput = path
	}
}

// WithLogRotation sets log rotation settings
// WithLogRotation 日志轮转配置
// 日志轮转配置，默认不开启
func WithLogRotation(filePath string, maxSize int, maxBackups int, maxAge int, isCompress bool) Option {
	return func(o *Options) {
		o.LogRotation = &LogRotation{
			FilePath:   filePath,
			MaxSize:    maxSize,
			MaxBackups: maxBackups,
			MaxAge:     maxAge,
			Compress:   isCompress,
		}
	}
}

// WithColor enables color output
// WithColor 启用颜色输出
// 启用颜色输出，默认不开启
// 注意：颜色输出会影响性能，建议在开发环境中使用
func WithColor() Option {
	return func(o *Options) {
		o.ColorEnabled = true
	}
}

// WithColorScheme sets a custom color scheme
// WithColorScheme 添加主题颜色方案
// 自定义颜色方案，默认使用默认颜色方案
// 注意：颜色输出会影响性能，建议在开发环境中使用
func WithColorScheme(scheme ColorScheme) Option {
	return func(o *Options) {
		o.ColorScheme = &scheme
	}
}

// WithMark 启用脱敏
// 不传入脱敏规则，默认使用默认脱敏规则
// 目前支持脱敏规则：
// 1. 密码脱敏
// 2. 手机号脱敏
func WithMark(maskRules ...MaskHandler) Option {
	return func(o *Options) {
		o.MaskEnable = true
		if len(maskRules) == 0 {
			maskRules = append(maskRules, DefaultMaskHandler()...)
		}
		o.maskRules = maskRules
	}
}

// applyOptions applies all options to the Options struct
// applyOptions 应用所有配置项
func applyOptions(opts ...Option) Options {
	options := Options{
		Level:      defaultLevel,
		JSONFormat: defaultJSONFormat,
		AddSource:  defaultAddSource,
		TimeFormat: time.DateTime,
		TimeZone:   time.Local.String(),
	}
	for _, opt := range opts {
		opt(&options)
	}
	if options.ColorEnabled && options.ColorScheme == nil {
		options.ColorScheme = DefaultFatihColorScheme
	}
	return options
}
