package modelStruct

import "encoding/json"

type ResponseModel struct {
	ID            uint64          `json:"id"`
	JSONRPC       string          `json:"jsonrpc"`
	Meta          json.RawMessage `json:"meta"` // 使用 map 来表示动态键值对
	RequestMethod string          `json:"requestMethod"`
	Result        json.RawMessage `json:"result"` // 使用 interface{} 以支持不同类型的结果
	Sender        string          `json:"sender"`
	Timestamp     int64           `json:"timestamp"` // 使用 int64 以存储时间戳
}
