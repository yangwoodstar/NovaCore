package modelStruct

import (
	"time"
)

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
	AppID              string `json:"appID"`
	RoomID             string `json:"roomID"`
	Cursor             int    `json:"cursor"`
	IP                 string `json:"ip"`
	DeviceInfo         string `json:"deviceInfo"`
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

type ActiveUserParams struct {
	ID         string `json:"id"`
	UserRoomID string `json:"userRoomID"`
}

type DisableRaiseHandParams struct {
	Disable bool `json:"disable"`
}

type Room struct {
	ID                int       `json:"id"`
	AppID             string    `json:"app_id"`
	ParentRoomID      string    `json:"parent_room_id"`
	RoomID            string    `json:"room_id"`
	Name              string    `json:"name"`
	ClassType         int8      `json:"class_type"`
	LiveType          int8      `json:"live_type"`
	LiveStatus        int8      `json:"live_status"`
	TeacherID         string    `json:"teacher_id"`
	StartTime         int64     `json:"start_time"`
	EndTime           int64     `json:"end_time"`
	PlayBackID        int64     `json:"play_back_id"`
	ConfigID          uint      `json:"config_id"`
	OssPath           []byte    `json:"oss_path"`
	BjyInfo           []byte    `json:"bjy_info"`
	CourseChapterType int8      `json:"course_chapter_type"`
	PlayBackVideoID   string    `json:"play_back_video_id"`
	MockChat          bool      `json:"mock_chat"`
	CreateAt          time.Time `json:"create_at"`
	UpdateAt          time.Time `json:"update_at"`
	IsSelfVideo       bool      `json:"is_self_video"`
	DisasterRecovery  bool      `json:"disaster_recovery"`
	TeacherName       string    `json:"teacher_name"`
	VideoID           string    `json:"video_id"`
	VideoSourceType   int8      `json:"video_source_type"`
}
