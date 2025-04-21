package modelStruct

type PlaybackCreate struct {
	AppID        string `json:"appID"`
	ParentRoomID string `json:"parentRoomID"`
	RoomID       string `json:"roomID"`
	VideoID      string `json:"videoID"`
	Name         string `json:"name"`
}

type RecordInfoPartModel struct {
	Start     int64 `json:"start"`
	End       int64 `json:"end"`
	Increment int64 `json:"increment"`
}
type RecordReportInfoModel struct {
	Parts        []*RecordInfoPartModel `json:"parts"`
	ClassStartTS int64                  `json:"classStartTS"`
}

// 公共基础结构体（包含完全相同的字段）
type BasePlayBackReq struct {
	ParentVideoID    *uint32                `json:"parentVideoID,omitempty"`
	ParentRoomID     string                 `json:"parentRoomID"`
	RoomID           string                 `json:"roomID"`
	AppID            string                 `json:"appID"`
	Size             *int64                 `json:"size,omitempty"`
	TranscodeDur     *int64                 `json:"transcodeDur,omitempty"`
	Duration         *int64                 `json:"duration,omitempty"`
	FirstFrame       *string                `json:"firstFrame,omitempty"`
	PlayBackInfo     *PlayBackInfo          `json:"playBackInfo,omitempty"`
	OriginalInfo     *PlayBackInfo          `json:"originalInfo,omitempty"`
	Width            int                    `json:"width"`
	Height           int                    `json:"height"`
	TaskID           string                 `json:"taskID"`
	Name             string                 `json:"name"`
	EncodeStatus     int                    `json:"encodeStatus"`
	CoursewareWidth  uint32                 `json:"coursewareWidth"`
	CoursewareHeight uint32                 `json:"coursewareHeight"`
	TeacherWidth     uint32                 `json:"teacherWidth"`
	TeacherHeight    uint32                 `json:"teacherHeight"`
	HRatioCourseware uint32                 `json:"hRatioCourseware"`
	HRatioTeacher    uint32                 `json:"hRatioTeacher"`
	Parts            []*RecordInfoPartModel `json:"parts"`
}

// PlayBackReq 结构体（创建请求）
type PlayBackReq struct {
	BasePlayBackReq // 嵌入公共字段
}

// UpdatePlayBackReq 结构体（更新请求）
type UpdatePlayBackReq struct {
	BasePlayBackReq         // 嵌入公共字段
	ID               uint32 `json:"ID"`
	DisasterRecovery bool   `json:"disasterRecovery,omitempty"`
}
