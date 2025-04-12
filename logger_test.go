package logger

import (
	"testing"
	"time"
)

func TestZapLogger(t *testing.T) {
	logger, err := NewLoggerWithType(ZapLogger)
	if err != nil {
		t.Fatal(err)
	}
	logger.Info("Info:hello world")
	logger.Infof("Infof:%v", "hello world")
	logger.Warn("warn:", "hello world")
	logger.Warnf("warnf:%v", "hello world")
	logger.Error("error:", "hello world")
	logger.Errorf("errorf:%v", "hello world")
}

func TestZapLoggerWithLevel(t *testing.T) {
	logger, err := NewLoggerWithType(ZapLogger, WithLevel(ErrorLevel))
	if err != nil {
		t.Fatal(err)
	}
	logger.Info("Info:hello world")
	logger.Infof("Infof:%v", "hello world")
	logger.Warn("warn:", "hello world")
	logger.Warnf("warnf:%v", "hello world")
	logger.Error("error:", "hello world")
	logger.Errorf("errorf:%v", "hello world")
}

func TestZapLoggerWithAddSource(t *testing.T) {
	logger, err := NewLoggerWithType(ZapLogger, WithAddSource())
	if err != nil {
		t.Fatal(err)
	}
	logger.Info("Info:hello world")
	logger.Infof("Infof:%v", "hello world")
	logger.Warn("warn:", "hello world")
	logger.Warnf("warnf:%v", "hello world")
	logger.Error("error:", "hello world")
	logger.Errorf("errorf:%v", "hello world")
}

func TestZapLoggerWithFileOutput(t *testing.T) {
	logger, err := NewLoggerWithType(ZapLogger, WithFileOutput("./zap.log"))
	if err != nil {
		t.Fatal(err)
	}
	logger.Info("Info:hello world")
	logger.Infof("Infof:%v", "hello world")
	logger.Warn("warn:", "hello world")
	logger.Warnf("warnf:%v", "hello world")
	logger.Error("error:", "hello world")
	logger.Errorf("errorf:%v", "hello world")
}

func TestZapLoggerWithJSONFormat(t *testing.T) {
	logger, err := NewLoggerWithType(ZapLogger, WithJSONFormat())
	if err != nil {
		t.Fatal(err)
	}
	logger.Info("Info:hello world")
	logger.Infof("Infof:%v", "hello world")
	logger.Warn("warn:", "hello world")
	logger.Warnf("warnf:%v", "hello world")
	logger.Error("error:", "hello world")
	logger.Errorf("errorf:%v", "hello world")
}

func TestZapLoggerWithTimeFormat(t *testing.T) {
	logger, err := NewLoggerWithType(ZapLogger, WithTimeFormat(time.DateTime))
	if err != nil {
		t.Fatal(err)
	}
	logger.Info("Info:hello world")
	logger.Infof("Infof:%v", "hello world")
	logger.Warn("warn:", "hello world")
	logger.Warnf("warnf:%v", "hello world")
	logger.Error("error:", "hello world")
	logger.Errorf("errorf:%v", "hello world")
}

func TestZapLoggerWithErrorOutPut(t *testing.T) {
	logger, err := NewLoggerWithType(ZapLogger, WithErrorOutPut("./sss"))
	if err != nil {
		t.Fatal(err)
	}
	logger.Info("Info:hello world")
	logger.Infof("Infof:%v", "hello world")
	logger.Warn("warn:", "hello world")
	logger.Warnf("warnf:%v", "hello world")
	logger.Error("error:", "hello world")
	logger.Errorf("errorf:%v", "hello world")
}

func TestZapLoggerWithColor(t *testing.T) {
	logger, err := NewLoggerWithType(ZapLogger, WithColor(), WithFileOutput("./logger.log"))
	if err != nil {
		t.Fatal(err)
	}
	logger.Info("Info:hello world")
	logger.Infof("Infof:%v", "hello world")
	logger.Warn("warn:", "hello world")
	logger.Warnf("warnf:%v", "hello world")
	logger.Error("error:", "hello world")
	logger.Errorf("errorf:%v", "hello world")
}

func TestZapLoggerWithTimeZone(t *testing.T) {
	logger, err := NewLoggerWithType(ZapLogger, WithTimeZone(JSTTime))
	if err != nil {
		t.Fatal(err)
	}
	logger.Info("Info:hello world")
	logger.Infof("Infof:%v", "hello world")
	logger.Warn("warn:", "hello world")
	logger.Warnf("warn:%v", "hello world")
	logger.Error("error:", "hello world")
	logger.Errorf("error:%v", "hello world")
}

func TestZapLoggerWithMark(t *testing.T) {
	logger, err := NewLoggerWithType(ZapLogger, WithMark())
	if err != nil {
		t.Fatal(err)
	}
	logger.Info("xxx", "pwd", 12234)
}

func TestZapLoggerWithMarkRules(t *testing.T) {
	logger, err := NewLoggerWithType(ZapLogger, WithMark(&AddressMask{}))
	if err != nil {
		t.Fatal(err)
	}
	logger.Info("xxx", "pwd", 12234)
	logger.Info("xxx", "address", "xxxx")
}

func TestLogrusLogger(t *testing.T) {
	logger, err := NewLoggerWithType(LogrusLogger)
	if err != nil {
		t.Fatal(err)
	}
	logger.Info("Info:hello world")
	logger.Infof("Infof:%v", "hello world")
	logger.Warn("warn:", "hello world")
	logger.Warnf("warnf:%v", "hello world")
	logger.Error("error:", "hello world")
	logger.Errorf("errorf:%v", "hello world")
}

func TestLogrusLoggerWithLevel(t *testing.T) {
	logger, err := NewLoggerWithType(LogrusLogger, WithLevel(ErrorLevel))
	if err != nil {
		t.Fatal(err)
	}
	logger.Info("Info:hello world")
	logger.Infof("Infof:%v", "hello world")
	logger.Warn("warn:", "hello world")
	logger.Warnf("warnf:%v", "hello world")
	logger.Error("error:", "hello world")
	logger.Errorf("errorf:%v", "hello world")
}

func TestLogrusLoggerWithAddSource(t *testing.T) {
	logger, err := NewLoggerWithType(LogrusLogger, WithAddSource())
	if err != nil {
		t.Fatal(err)
	}
	logger.Info("Info:hello world")
	logger.Infof("Infof:%v", "hello world")
	logger.Warn("warn:", "hello world")
	logger.Warnf("warnf:%v", "hello world")
	logger.Error("error:", "hello world")
	logger.Errorf("errorf:%v", "hello world")
}

func TestLogrusLoggerWithFileOutput(t *testing.T) {
	logger, err := NewLoggerWithType(LogrusLogger, WithFileOutput("./logger2.log"))
	if err != nil {
		t.Fatal(err)
	}
	logger.Info("Info:hello world")
	logger.Infof("Infof:%v", "hello world")
	logger.Warn("warn:", "hello world")
	logger.Warnf("warnf:%v", "hello world")
	logger.Error("error:", "hello world")
	logger.Errorf("errorf:%v", "hello world")
}

func TestLogrusLoggeWithJSONFormat(t *testing.T) {
	logger, err := NewLoggerWithType(ZapLogger, WithJSONFormat())
	if err != nil {
		t.Fatal(err)
	}
	logger.Info("Info:hello world")
	logger.Infof("Infof:%v", "hello world")
	logger.Warn("warn:", "hello world")
	logger.Warnf("warnf:%v", "hello world")
	logger.Error("error:", "hello world")
	logger.Errorf("errorf:%v", "hello world")
}

func TestLogrusLoggrWithTimeFormat(t *testing.T) {
	logger, err := NewLoggerWithType(LogrusLogger, WithTimeFormat(time.DateTime))
	if err != nil {
		t.Fatal(err)
	}
	logger.Info("Info:hello world")
	logger.Infof("Infof:%v", "hello world")
	logger.Warn("warn:", "hello world")
	logger.Warnf("warnf:%v", "hello world")
	logger.Error("error:", "hello world")
	logger.Errorf("errorf:%v", "hello world")
}

func TestLogrusLoggerWithFields(t *testing.T) {
	logger, err := NewLoggerWithType(LogrusLogger)
	if err != nil {
		t.Fatal(err)
	}
	loggerFiled := logger.WithFields(map[string]any{"key": "value"})
	logger.Info("Info:hello world")
	logger.Infof("Infof:%v", "hello world")
	loggerFiled.Warn("warn:", "hello world")
	loggerFiled.Warnf("warn:%v", "hello world")
	logger.Error("error:", "hello world")
	logger.Errorf("errorf:%v", "hello world")
}

func TestLogrusLoggerWithErrorOutPut(t *testing.T) {
	logger, err := NewLoggerWithType(LogrusLogger, WithErrorOutPut("./loggerout.log"), WithFileOutput("./sss"))
	if err != nil {
		t.Fatal(err)
	}
	logger.Info("Info:hello world")
	logger.Infof("Infof:%v", "hello world")
	logger.Warn("warn:", "hello world")
	logger.Warnf("warnf:%v", "hello world")
	logger.Error("error:", "hello world")
	logger.Errorf("errorf:%v", "hello world")
}

func TestLogrusLoggerWithColor(t *testing.T) {
	logger, err := NewLoggerWithType(LogrusLogger, WithColor(), WithFileOutput("./logger.log"))
	if err != nil {
		t.Fatal(err)
	}
	logger.Info("Info:hello world")
	logger.Infof("Infof:%v", "hello world")
	logger.Warn("warn:", "hello world")
	logger.Warnf("warnf:%v", "hello world")
	logger.Error("error:", "hello world")
	logger.Errorf("errorf:%v", "hello world")
}

func TestLogrusLoggerWithTimeZone(t *testing.T) {
	logger, err := NewLoggerWithType(LogrusLogger, WithTimeZone(JSTTime))
	if err != nil {
		t.Fatal(err)
	}
	logger.Info("Info:hello world")
	logger.Infof("Infof:%v", "hello world")
	logger.Warn("warn:", "hello world")
	logger.Warnf("warn:%v", "hello world")
	logger.Error("error:", "hello world")
	logger.Errorf("error:%v", "hello world")
}

func TestLogrusLoggerWithMark(t *testing.T) {
	logger, err := NewLoggerWithType(LogrusLogger, WithMark())
	if err != nil {
		t.Fatal(err)
	}
	logger.Info("xxx", "pwd", 12234)
}

func TestLogrusLoggerWithMarkRules(t *testing.T) {
	logger, err := NewLoggerWithType(LogrusLogger, WithMark(&AddressMask{}))
	if err != nil {
		t.Fatal(err)
	}
	logger.Info("xxx", "pwd", 12234)
	logger.Info("xxx", "address", "xxxx")
}

func TestKlogLogger(t *testing.T) {
	logger, err := NewLoggerWithType(KlogLogger)
	if err != nil {
		t.Fatal(err)
	}
	logger.Info("Info:hello world")
	logger.Infof("Infof:%v", "hello world")
	logger.Warn("warn:", "hello world")
	logger.Warnf("warnf:%v", "hello world")
	logger.Error("error:", "hello world")
	logger.Errorf("errorf:%v", "hello world")
}

func TestKlogLoggerWithLevel(t *testing.T) {
	logger, err := NewLoggerWithType(KlogLogger, WithLevel(ErrorLevel))
	if err != nil {
		t.Fatal(err)
	}
	logger.Info("Info:hello world")
	logger.Infof("Infof:%v", "hello world")
	logger.Warn("warn:", "hello world")
	logger.Warnf("warnf:%v", "hello world")
	logger.Error("error:", "hello world")
	logger.Errorf("errorf:%v", "hello world")
}

func TestKlogLoggerWithFileOutput(t *testing.T) {
	logger, err := NewLoggerWithType(KlogLogger, WithFileOutput("./app.log"))
	if err != nil {
		t.Fatal(err)
	}
	logger.Info("Info:hello world")
	logger.Infof("Infof:%v", "hello world")
	logger.Warn("warn:", "hello world")
	logger.Warnf("warnf:%v", "hello world")
	logger.Error("error:", "hello world")
	logger.Errorf("errorf:%v", "hello world")
}

func TestKlogLoggerWithErrorOutPut(t *testing.T) {
	logger, err := NewLoggerWithType(KlogLogger, WithErrorOutPut("./app_error.log"))
	if err != nil {
		t.Fatal(err)
	}
	logger.Info("Info:hello world")
	logger.Infof("Infof:%v", "hello world")
	logger.Warn("warn:", "hello world")
	logger.Warnf("warnf:%v", "hello world")
	logger.Error("error:", "hello world")
	logger.Errorf("errorf:%v", "hello world")
}

func TestKlogLoggerWithErrorOutPutAndFile(t *testing.T) {
	logger, err := NewLoggerWithType(KlogLogger, WithFileOutput("./app.log"), WithErrorOutPut("./app_error.log"))
	if err != nil {
		t.Fatal(err)
	}
	logger.Info("Info:hello world")
	logger.Infof("Infof:%v", "hello world")
	logger.Warn("warn:", "hello world")
	logger.Warnf("warnf:%v", "hello world")
	logger.Error("error:", "hello world")
	logger.Errorf("errorf:%v", "hello world")
}

func TestKlogLoggerWithColor(t *testing.T) {
	logger, err := NewLoggerWithType(KlogLogger, WithColor(), WithFileOutput("./logger.log"))
	if err != nil {
		t.Fatal(err)
	}
	logger.Info("Info:hello world")
	logger.Infof("Infof:%v", "hello world")
	logger.Warn("warn:", "hello world")
	logger.Warnf("warnf:%v", "hello world")
	logger.Error("error:", "hello world")
	logger.Errorf("errorf:%v", "hello world")
}

func TestKlogLoggerWithLogRotation(t *testing.T) {
	logger, err := NewLoggerWithType(SlogLogger, WithLogRotation("app.log", 1, 5, 0, true))
	if err != nil {
		t.Fatal(err)
	}
	for i := 0; i < 1000; i++ {
		logger.Info("Info:hello world")
		logger.Infof("Infof:%v", "hello world")
		logger.Warn("warn:", "hello world")
		logger.Warnf("warnf:%v", "hello world")
		logger.Error("error:", "hello world")
		logger.Errorf("errorf:%v", "hello world")
	}
}

func TestKlogLoggerWithTimeZone(t *testing.T) {
	logger, err := NewLoggerWithType(KlogLogger, WithTimeZone(JSTTime))
	if err != nil {
		t.Fatal(err)
	}
	logger.Info("Info:hello world")
	logger.Infof("Infof:%v", "hello world")
	logger.Warn("warn:", "hello world")
	logger.Warnf("warnf:%v", "hello world")
	logger.Error("error:", "hello world")
	logger.Errorf("error:%v", "hello world")
}

func TestKlogLoggerWithMark(t *testing.T) {
	logger, err := NewLoggerWithType(KlogLogger, WithMark())
	if err != nil {
		t.Fatal(err)
	}
	logger.Error("xxx", "pwd", 12234)
}

func TestKlogLoggerWithMarkRules(t *testing.T) {
	logger, err := NewLoggerWithType(KlogLogger, WithMark(&AddressMask{}))
	if err != nil {
		t.Fatal(err)
	}
	logger.Error("xxx", "address", "xxxx")
}

func TestSlogLogger(t *testing.T) {
	logger, err := NewLoggerWithType(SlogLogger)
	if err != nil {
		t.Fatal(err)
	}
	logger.Info("Info:hello world")
	logger.Infof("Infof:%v", "hello world")
	logger.Warn("warn:%v", "hello world")
	logger.Warnf("warnf:%v", "hello world")
	logger.Error("error:", "hello world")
	logger.Errorf("errorf:%v", "hello world")
}

func TestSlogLoggerWithLevel(t *testing.T) {
	logger, err := NewLoggerWithType(SlogLogger, WithAddSource(), WithLevel(ErrorLevel))
	if err != nil {
		t.Fatal(err)
	}
	logger.Info("Info:hello world")
	logger.Infof("Infof:%v", "hello world")
	logger.Warn("warn:", "hello world")
	logger.Warnf("warnf:%v", "hello world")
	logger.Error("error:", "hello world")
	logger.Errorf("errorf:%v", "hello world")
}

func TestSlogLoggerWithAddSource(t *testing.T) {
	logger, err := NewLoggerWithType(SlogLogger, WithAddSource())
	if err != nil {
		t.Fatal(err)
	}
	logger.Info("Info:hello world")
	logger.Infof("Infof:%v", "hello world")
	logger.Warn("warn:", "hello world")
	logger.Warnf("warnf:%v", "hello world")
	logger.Error("error:", "hello world")
	logger.Errorf("errorf:%v", "hello world")
}

func TestSlogLoggerWithFileOutput(t *testing.T) {
	logger, err := NewLoggerWithType(SlogLogger, WithFileOutput("./logger.log"))
	if err != nil {
		t.Fatal(err)
	}
	logger.Info("Info:hello world")
	logger.Infof("Infof:%v", "hello world")
	logger.Warn("warn:", "hello world")
	logger.Warnf("warnf:%v", "hello world")
	logger.Error("error:", "hello world")
	logger.Errorf("errorf:%v", "hello world")
}

func TestSlogLoggerWithJSONFormat(t *testing.T) {
	logger, err := NewLoggerWithType(SlogLogger, WithJSONFormat())
	if err != nil {
		t.Fatal(err)
	}
	logger.Info("Info:hello world")
	logger.Infof("Infof:%v", "hello world")
	logger.Warn("warn:", "hello world")
	logger.Warnf("warnf:%v", "hello world")
	logger.Error("error:", "hello world")
	logger.Errorf("errorf:%v", "hello world")
}

func TestSlogLoggerWithTimeFormat(t *testing.T) {
	logger, err := NewLoggerWithType(SlogLogger, WithTimeFormat(time.Stamp))
	if err != nil {
		t.Fatal(err)
	}
	logger.Info("Info:hello world")
	logger.Infof("Infof:%v", "hello world")
	logger.Warn("warn:", "hello world")
	logger.Warnf("warnf:%v", "hello world")
	logger.Error("error:", "hello world")
	logger.Errorf("errorf:%v", "hello world")
}

func TestSlogLoggerWithFields(t *testing.T) {
	logger, err := NewLoggerWithType(SlogLogger, WithTimeFormat(time.Stamp))
	if err != nil {
		t.Fatal(err)
	}
	loggerFiled := logger.WithFields(map[string]any{"key": "value"})
	logger.Info("Info:hello world")
	logger.Infof("Infof:%v", "hello world")
	loggerFiled.Warn("warn:", "hello world")
	loggerFiled.Warnf("warn:%v", "hello world")
	logger.Error("error:", "hello world")
	logger.Errorf("errorf:%v", "hello world")
}

func TestSlogLoggerWithErrorOutPut(t *testing.T) {
	logger, err := NewLoggerWithType(SlogLogger, WithErrorOutPut("./loggerout.log"))
	if err != nil {
		t.Fatal(err)
	}
	loggerFiled := logger.WithFields(map[string]any{"key": "value"})
	logger.Info("Info:hello world")
	logger.Infof("Infof:%v", "hello world")
	loggerFiled.Warn("warn:", "hello world")
	loggerFiled.Warnf("warn:%v", "hello world")
	logger.Error("error:", "hello world")
	logger.Errorf("errorf:%v", "hello world")
}

func TestSlogLoggerWithLogRotation(t *testing.T) {
	logger, err := NewLoggerWithType(SlogLogger, WithLogRotation("app.log", 1, 5, 0, true))
	if err != nil {
		t.Fatal(err)
	}
	for i := 0; i < 10000; i++ {
		logger.Info("Info:hello world")
		logger.Infof("Infof:%v", "hello world")
		logger.Warn("warn:", "hello world")
		logger.Warnf("warnf:%v", "hello world")
		logger.Error("error:", "hello world")
		logger.Errorf("errorf:%v", "hello world")
	}
}

func TestSlogLoggeWithColor(t *testing.T) {
	logger, err := NewLoggerWithType(SlogLogger, WithColor(), WithFileOutput("./logger.log"))
	if err != nil {
		t.Fatal(err)
	}
	logger.Info("Info:hello world")
	logger.Infof("Infof:%v", "hello world")
	logger.Warn("warn:", "hello world")
	logger.Warnf("warnf:%v", "hello world")
	logger.Error("error:", "hello world")
	logger.Errorf("errorf:%v", "hello world")
}

func BenchmarkOriginalDebug(b *testing.B) {
	logger, _ := newSlogLogger(Options{Level: DebugLevel})
	b.ResetTimer()
	for i := 0; i < 100000; i++ {
		logger.Info("Info:hello world")
		logger.Infof("Infof:%v", "hello world")
		logger.Warn("warn:", "hello world")
		logger.Warnf("warnf:%v", "hello world")
		logger.Error("error:", "hello world")
		logger.Errorf("errorf:%v", "hello world")
	}
}

// BenchmarkOriginalDebug-8   	       1	1685798209 ns/op
// BenchmarkOriginalDebug-8   	1000000000	         0.5068 ns/op
// BenchmarkOriginalDebug-8   	1000000000	         0.4962 ns/op
// BenchmarkOriginalDebug-8   	 1144393	      1241 ns/op

func TestSlogLoggeWithColorTheme(t *testing.T) {
	theme := ColorScheme{
		CodeType: CodeTypeANSI,
		Info: &Color{
			ansi: "\u001B[35m",
		},
	}
	logger, err := NewLoggerWithType(SlogLogger, WithColor(), WithFileOutput("./logger.log"), WithColorScheme(theme))
	if err != nil {
		t.Fatal(err)
	}
	logger.Info("Info:hello world")
	logger.Infof("Infof:%v", "hello world")
	logger.Warn("warn:", "hello world")
	logger.Warnf("warnf:%v", "hello world")
	logger.Error("error:", "hello world")
	logger.Errorf("errorf:%v", "hello world")
}

func TestSlogLoggeWithTimeZone(t *testing.T) {
	logger, err := NewLoggerWithType(SlogLogger, WithTimeZone(JSTTime))
	if err != nil {
		t.Fatal(err)
	}
	logger.Info("Info:hello world")
	logger.Infof("Infof:%v", "hello world")
	logger.Warn("warn:", "hello world")
	logger.Warnf("warn:%v", "hello world")
	logger.Error("error:", "hello world")
	logger.Errorf("error:%v", "hello world")
}

func TestSlogLoggWithMark(t *testing.T) {
	logger, err := NewLoggerWithType(SlogLogger, WithMark())
	if err != nil {
		t.Fatal(err)
	}
	logger.Info("xxx", "pwd", 12234)
}

func TestSlogLoggWithMarkRules(t *testing.T) {
	logger, err := NewLoggerWithType(SlogLogger, WithMark(&AddressMask{}))
	if err != nil {
		t.Fatal(err)
	}
	logger.Info("xxx", "pwd", 12234)
	logger.Info("xxx", "address", "xxxx")
}

// AddressMask 地址脱敏处理器
type AddressMask struct{}

func (m *AddressMask) Mask(fieldName string, value any) any {
	if fieldName == "address" {
		if s, ok := value.(string); ok {
			if len(s) > 4 {
				return s[:4] + "****"
			}
			return "****"
		}
	}
	return value
}
