package modelStruct

import "encoding/json"

type BigDataModel struct {
	ParentRoomID string `json:"parentRoomID"`
}

// User 结构体定义
type User struct {
	AppID              string `json:"appID"`
	RoomID             string `json:"roomID"`
	Avatar             string `json:"avatar"`
	EndType            int    `json:"endType"`
	HasAudioPermission bool   `json:"hasAudioPermission"`
	HasVideoPermission bool   `json:"hasVideoPermission"`
	ID                 string `json:"id"`
	Name               string `json:"name"`
	Role               int    `json:"role"`
	SessionID          string `json:"sessionID"`
	Status             int    `json:"status"`
}

// Params 结构体定义
type Params struct {
	Override  bool   `json:"override"`
	RoomID    string `json:"roomID"`
	SessionID string `json:"sessionID"`
	User      User   `json:"user"`
}

type ParamsRoomIDModel struct {
	RoomID string `json:"roomID"`
}

// From 定义发送者信息结构体
type FromUser struct {
	Avatar string `json:"avatar"` // 用户头像 URL
	ID     string `json:"id"`     // 用户 ID
	Name   string `json:"name"`   // 用户名称
	Role   int    `json:"role"`   // 用户角色
	RoomID string `json:"roomID"` // 房间 ID
}

type RPCBroadcastModel struct {
	JSONRPC       string          `json:"jsonrpc"`
	Method        string          `json:"method"`
	AppID         string          `json:"appID"`
	RoomID        string          `json:"roomID"`
	Req           string          `json:"req"`
	ParentRoomID  string          `json:"parentRoomID"`
	CurrentRoomID string          `json:"currentRoomID"`
	Params        json.RawMessage `json:"params"`
	RoomIDs       []string        `json:"roomIDs"`
	UserIDs       []string        `json:"userIDs"`
	Reversal      bool            `json:"reversal"`
	Reply         string          `json:"reply"`
	Index         int             `json:"index"`
	BigData       BigDataModel    `json:"bigData"`
	From          *FromUser       `json:"from,omitempty"`
}
