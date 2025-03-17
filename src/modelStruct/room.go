package modelStruct

type LeaveParams struct {
}

type JoinParams struct {
	RoomID    string    `json:"roomID"`
	ForceJoin bool      `json:"forceJoin"`
	User      UserModel `json:"user"`
}

// V5LiveUserModel 表示用户模型
type UserModel struct {
	ID                 string `json:"id"`
	Role               int    `json:"role"`
	Avatar             string `json:"avatar"`
	Name               string `json:"name"`
	EndType            int    `json:"endType"`
	Status             int    `json:"status"`
	HasAudioPermission bool   `json:"hasAudioPermission"`
	HasVideoPermission bool   `json:"hasVideoPermission"`
}
