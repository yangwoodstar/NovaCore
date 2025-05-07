package tools

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"github.com/yangwoodstar/NovaCore/src/api"
	"github.com/yangwoodstar/NovaCore/src/modelStruct"
	"go.uber.org/zap"
	"time"
)

type RetryableRequest struct {
	MaxRetries    int                              // 最大重试次数（默认3）
	OperationName string                           // 操作名称（用于日志）
	Context       context.Context                  // 上下文控制
	ReqExecutor   func() error                     // 请求执行函数
	ErrChecker    func() error                     // 错误检查函数
	ErrCallback   func(err error, index int) error // 错误回调函数（可选）
	DelayStrategy func(int) time.Duration          // 延时策略（可选）
}

type SendConfig struct {
	Url    string
	Mobile string
	RoomID string
	Method string
	Logger *zap.Logger
}

// 通用重试执行器
func Retry(req *RetryableRequest) error {
	if req.MaxRetries <= 0 {
		req.MaxRetries = 3
	}
	if req.Context == nil {
		req.Context = context.Background()
	}
	if req.DelayStrategy == nil {
		req.DelayStrategy = func(int) time.Duration { return 0 } // 默认无延时
	}

	var lastErr error
	for attempt := 0; attempt < req.MaxRetries+1; attempt++ {
		// 执行请求
		err := req.ReqExecutor()

		// 检查错误
		checkErr := req.ErrChecker()
		if err == nil && checkErr == nil {
			return nil
		}

		// 记录错误
		if err != nil {
			lastErr = err
		} else {
			lastErr = checkErr
		}

		_ = req.ErrCallback(lastErr, attempt)

		// 终止条件检查
		if attempt >= req.MaxRetries {
			break
		}

		// 计算延时
		delay := req.DelayStrategy(attempt)
		if delay < 0 {
			delay = 0
		}

		// 执行延时或监听上下文取消
		if delay > 0 {
			select {
			case <-req.Context.Done():
				return fmt.Errorf("操作 %q 被取消: %v", req.OperationName, req.Context.Err())
			case <-time.After(delay):
				// 继续重试
			}
		} else {
			select {
			case <-req.Context.Done():
				return fmt.Errorf("操作 %q 被取消: %v", req.OperationName, req.Context.Err())
			default:
				// 立即重试
			}
		}
	}

	return fmt.Errorf("操作 %q 在 %d 次尝试后失败，最后错误: %v",
		req.OperationName, req.MaxRetries+1, lastErr)
}

type ConnectionRequest interface {
	Request(request *modelStruct.RequestModel, response *modelStruct.ResponseModel, ctx context.Context) error
}

func SendRequest(connectionRequest ConnectionRequest, request *modelStruct.RequestModel, sendConfig SendConfig) {
	response := &modelStruct.ResponseModel{}

	reqExecutor := func() error {
		ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
		defer cancel()
		return connectionRequest.Request(request, response, ctx)
	}

	errCallback := func(errMsg error, index int) error {
		joinRoomData, errMarshal := json.Marshal(request)
		if errMarshal != nil {
			sendConfig.Logger.Error("Failed to marshal join room data", zap.Error(errMarshal))
		}
		if joinRoomData != nil {
			message, err := api.SendWaningMessage(sendConfig.Url, fmt.Sprintf("send to liveRoom error: %v %v %v  %s \n retry time: %v \n", string(joinRoomData), sendConfig.RoomID, errMsg, sendConfig.Method, index), sendConfig.Mobile)
			if err != nil {
				sendConfig.Logger.Error("Failed to send message", zap.Error(err))
			} else {
				sendConfig.Logger.Info("Message sent successfully", zap.String("response", message))
			}
		} else {
			// 如果 joinRoomData 为空，则使用 request 的字符串表示
			message, err := api.SendWaningMessage(sendConfig.Url, fmt.Sprintf("send to liveRoom error: %v %v %v  %s \n retry time: %v \n", request, sendConfig.RoomID, errMsg, sendConfig.Method, index), sendConfig.Mobile)
			if err != nil {
				sendConfig.Logger.Error("Failed to send message", zap.Error(err))
			} else {
				sendConfig.Logger.Info("Message sent successfully", zap.String("response", message))
			}
		}
		return nil
	}

	req := &RetryableRequest{
		MaxRetries:    3,
		OperationName: sendConfig.Method,
		Context:       context.Background(),
		ReqExecutor:   reqExecutor,
		ErrChecker: func() error {
			if response.Error == nil {
				return nil
			}
			return errors.New(fmt.Sprintf("%v", response.Error))
		},
		ErrCallback: errCallback,
		DelayStrategy: func(index int) time.Duration {
			return time.Second * time.Duration(1)
		},
	}

	err := Retry(req)
	if err != nil {
		sendConfig.Logger.Error("Retry failed", zap.Error(err))
	} else {
		sendConfig.Logger.Info("Request succeeded", zap.Any("response", response))
	}
}
