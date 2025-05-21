package modelStruct

type BaseRoomInfo struct {
	AppID  string `json:"appID"`
	RoomID string `json:"roomID"`
	UserID string `json:"userID"`
	Reply  string `json:"reply"`
	ID     int64  `json:"ID"`
	Method string `json:"method"`
	Role   int    `json:"role"`
}
