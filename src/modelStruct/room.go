package modelStruct

type EmptyParams struct {
}

type ToggleParams struct {
	Status int `json:"status"`
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

type CustomDataParams struct {
	Key   string          `json:"key"`
	Value CustomDataValue `json:"value"`
}

type CustomDataValue struct {
	Type   int            `json:"type"`
	User   CustomDataUser `json:"user"`
	UserID string         `json:"userID"`
}

type CustomDataUser struct {
	ID            string `json:"id"`
	Role          int    `json:"role"`
	Name          string `json:"name"`
	Avatar        string `json:"avatar"`
	RoomID        string `json:"roomID"`
	Status        int    `json:"status"`
	VideoOn       bool   `json:"videoOn"`
	AudioOn       bool   `json:"audioOn"`
	ScreenVideoOn bool   `json:"screenVideoOn"`
	ScreenAudioOn bool   `json:"screenAudioOn"`
	OnStage       bool   `json:"onStage"`
}
