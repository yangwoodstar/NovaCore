package tools

import (
	"context"
	"go.uber.org/zap"
	"runtime/debug"
)

// SafeGo 安全协程（自动捕获 panic 并支持 context 控制）
func SafeGo(ctx context.Context, handler func(context.Context)) {
	go func() {
		defer func() {
			if r := recover(); r != nil {
				Logger.Error("SafeGo Panic", zap.Any("error", r), zap.String("stack", string(debug.Stack())))
			}
		}()

		handler(ctx) // 将 context 传递给业务逻辑
	}()
}

// SafeGoWithParams 支持参数的安全协程
func SafeGoWithParams(ctx context.Context, handler func(context.Context, ...interface{}), params ...interface{}) {
	go func() {
		defer func() {
			if r := recover(); r != nil {
				Logger.Error("SafeGoWithParams Panic", zap.Any("error", r), zap.String("stack", string(debug.Stack())))
			}
		}()

		handler(ctx, params...) // 传递 context 和参数
	}()
}
