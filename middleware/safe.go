package middleware

import (
	"go.uber.org/zap"
	"runtime"
	"runtime/debug"
)

// getCurrentGoroutineStack 获取
func getCurrentGoroutineStack() string {
	var buf [4096]byte
	runtime.Stack(buf[:], false)
	return string(buf[:])
}

// SafeFunc 安全方法
func SafeFunc(fn func()) {
	defer func() {
		if err := recover(); err != nil {
			logger.Error(
				"safe func recover",
				zap.Any("err", err),
				zap.String("stack-1", getCurrentGoroutineStack()),
				zap.ByteString("stack-2", debug.Stack()))
		}
	}()

	fn()
}
