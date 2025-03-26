package modelStruct

type SnapShotStart struct {
	AppID       string `json:"AppId"`       // 应用的唯一标志
	RoomID      string `json:"RoomId"`      // 房间 ID，是房间的唯一标志
	TaskID      string `json:"TaskId"`      // 截图任务 ID
	UserID      string `json:"UserId"`      // 用户 ID
	Interval    int32  `json:"Interval"`    // 截图间隔，单位为秒
	CallbackUrl string `json:"CallbackUrl"` // 截图回调地址
}

type SnapShotStop struct {
	AppID  string `json:"AppId"`  // 应用的唯一标志
	RoomID string `json:"RoomId"` // 房间 ID，是房间的唯一标志
	TaskID string `json:"TaskId"` // 截图任务 ID
}

type SnapShotCallback struct {
	AppID    string `json:"AppId"`    // 应用的唯一标志
	RoomID   string `json:"RoomId"`   // 房间 ID，是房间的唯一标志
	TaskID   string `json:"TaskId"`   // 截图任务 ID
	ImageUrl string `json:"ImageUrl"` // 截图地址
	Width    int32  `json:"Width"`    // 截图宽度
	Height   int32  `json:"Height"`   // 截图高度
	Size     int64  `json:"Size"`     // 文件大小
}
