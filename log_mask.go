package logger

import (
	"regexp"
	"sync"
)

// MaskHandler 脱敏处理器接口
type MaskHandler interface {
	// Mask 对输入字符串进行脱敏处理
	// fieldName 是字段名，可用于识别特定字段
	// value 是原始值
	// 返回脱敏后的值
	Mask(fieldName string, value any) any
}

// MaskProcessor 脱敏处理器
type MaskProcessor struct {
	maskHandlers []MaskHandler
	mu           sync.RWMutex
}

// NewMaskProcessor 创建新的脱敏处理器
func NewMaskProcessor(handlers ...MaskHandler) *MaskProcessor {
	return &MaskProcessor{
		maskHandlers: handlers,
	}
}

// RegisterHandler 注册脱敏处理器
func (p *MaskProcessor) RegisterHandler(handler ...MaskHandler) {
	p.mu.Lock()
	defer p.mu.Unlock()
	p.maskHandlers = append(p.maskHandlers, handler...)
}

// Process 执行脱敏处理
func (p *MaskProcessor) Process(args ...any) []any {
	p.mu.RLock()
	defer p.mu.RUnlock()

	for i := 0; i < len(args); i += 2 {
		if key, ok := args[i].(string); ok {
			for _, masker := range p.maskHandlers {
				args[i+1] = masker.Mask(key, args[i+1])
			}
		}
	}
	return args
}

/*
 * 以下是一些常见的脱敏处理器
 * 你可以根据需要添加更多的处理器
 * 也可以自行实现 MaskHandler 接口
 */

// PasswordMark 脱敏处理器，用于标记密码字段
type PasswordMark struct{}

func (p *PasswordMark) Mask(fieldName string, value any) any {
	if fieldName == "password" || fieldName == "pwd" {
		return "[****]"
	}
	return value
}

// PhoneMask 脱敏处理器，用于隐藏手机号中间四位
type PhoneMask struct{}

func (m *PhoneMask) Mask(fieldName string, value any) any {
	if fieldName == "phone" {
		if s, ok := value.(string); ok {
			return regexp.MustCompile(`(\d{3})\d{4}(\d{4})`).ReplaceAllString(s, "$1****$2")
		}
	}
	return value
}
