package transportCore

// 定义统一消息接口
type UnificationMessage interface {
	GetBody() []byte
	Ack() error
	GetTopic() string
}
