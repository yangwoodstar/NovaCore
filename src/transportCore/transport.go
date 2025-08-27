package transportCore

import (
	"sync"
)

var (
	transportInstance *UnifiedTransport
	transportOnce     sync.Once
)

// Transport 接口定义了传输层的基本操作
type Transport interface {
	Read() (UnificationMessage, error)
	Write(p []byte, topic, routerKey string, priority int) error
	Close()
}

// UnifiedTransport 统一的传输结构体
type UnifiedTransport struct {
	senders  map[string]Transport
	receives map[string]Transport
	msgChan  chan UnificationMessage
	mu       sync.Mutex
}

// GetUnifiedTransport NewUnifiedTransport 创建一个新的 UnifiedTransport 实例
func GetUnifiedTransport() *UnifiedTransport {
	return transportInstance
}

// NewUnifiedTransport 创建一个新的 UnifiedTransport 实例
func NewUnifiedTransport() *UnifiedTransport {
	transportOnce.Do(func() {
		transportInstance = &UnifiedTransport{
			senders:  make(map[string]Transport),
			receives: make(map[string]Transport),
			msgChan:  make(chan UnificationMessage, 1000),
		}
	})
	return transportInstance
}

func (u *UnifiedTransport) AddSender(appId, topic string, t Transport) {
	u.mu.Lock()
	defer u.mu.Unlock()
	u.senders[appId+topic] = t
}

func (u *UnifiedTransport) AddReceiver(appId, topic string, t Transport) {
	u.mu.Lock()
	defer u.mu.Unlock()
	u.receives[appId+topic] = t
	go func() {
		for {
			msg, readErr := t.Read()
			if readErr != nil {

				continue
			}
			u.msgChan <- msg
		}
	}()
}

// Read 从传输层读取数据
func (u *UnifiedTransport) Read() (UnificationMessage, error) {
	select {
	case msg := <-u.msgChan:
		//rt.logger.Info("Read message", zap.String("topic", responseMessage.Topic), zap.String("msg", string(responseMessage.Msg)))
		return msg, nil
	}
}

// Write 向传输层写入数据
func (u *UnifiedTransport) Write(p []byte, appID, topic, routerKey string, priority int) error {
	transport, exist := u.senders[appID+topic]
	if exist {
		err := transport.Write(p, topic, routerKey, priority)
		if err != nil {

		}
	}
	return nil
}

// Close 关闭传输层
func (u *UnifiedTransport) Close() {

}
