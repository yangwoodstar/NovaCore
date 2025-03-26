package constString

const (
	NoPriority         = 0
	SmallClassPriority = 5
	LargeClassPriority = 6
	DefaultPriority    = 10
)

const (
	RabbitMQKindFanout               = "fanout"
	RabbitMQKindConsistentHash       = "x-consistent-hash"
	RabbitMQKindDirect               = "direct"
	RabbitMQConsistentHashBindingKey = "1"
	RabbitMQDefaultBindingKey        = ""
)

const (
	TranscodeNormalType   = 0
	TranscodeFakeLiveType = 1
)
