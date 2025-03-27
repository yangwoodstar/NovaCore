package modelStruct

type PlaybackCreate struct {
	AppID        string `json:"appID"`
	ParentRoomID string `json:"parentRoomID"`
	RoomID       string `json:"roomID"`
	VideoID      string `json:"videoID"`
	Name         string `json:"name"`
}
