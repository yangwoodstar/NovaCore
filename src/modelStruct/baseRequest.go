package modelStruct

type BaseRequest struct {
	AppID         string `json:"appID"`
	ParentRoomID  string `json:"parentRoomID"`
	RoomID        string `json:"roomID"`
	CurrentRoomID string `json:"currentRoomID"`
	Group         string `json:"group"`
	UserID        string `json:"userID"`
	SessionID     string `json:"sessionID"`
	AgentID       string `json:"agentID"`
	UserRole      int    `json:"userRole"`
}
