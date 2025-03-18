package modelStruct

import "encoding/json"

type RequestModel struct {
	JsonRPC string          `json:"jsonrpc"`
	Method  string          `json:"method"`
	Reply   string          `json:"reply"`
	Forward string          `json:"forward"`
	Params  json.RawMessage `json:"params"` // 使用指针以便处理 JsonNode
	ID      int64           `json:"id"`
	Token   string          `json:"token"`
	To      string          `json:"to"`
	BaseRequest
}
