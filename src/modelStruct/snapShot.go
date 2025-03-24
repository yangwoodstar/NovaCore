package modelStruct

type SnapShotStart struct {
	AppID    string `json:"AppId"`    // 应用的唯一标志
	RoomID   string `json:"RoomId"`   // 房间 ID，是房间的唯一标志
	TaskID   string `json:"TaskId"`   // 截图任务 ID
	UserID   string `json:"UserId"`   // 用户 ID
	Interval int32  `json:"Interval"` // 截图间隔，单位为秒
}

type SnapShotStop struct {
	AppID  string `json:"AppId"`  // 应用的唯一标志
	RoomID string `json:"RoomId"` // 房间 ID，是房间的唯一标志
	TaskID string `json:"TaskId"` // 截图任务 ID
}
