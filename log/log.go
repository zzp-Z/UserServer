package log

import (
	"context"
	"fmt"
	"github.com/zeromicro/go-zero/core/logc"
)

type ErrorContent struct {
	Message   string // 错误内容，例如 "<张三"
	Error     error  // 错误消息，例如 "非法用户名"
	ErrorCode string // 错误代码，例如 "RVI901"
}

// Error 错误日志
func Error(ctx context.Context, c ErrorContent) {
	info := fmt.Sprintf("ErrorCode: [%s], Message: [%s], Error: [%v]", c.ErrorCode, c.Message, c.Error)
	logc.Error(ctx, info)
}
