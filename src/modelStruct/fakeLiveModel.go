package modelStruct

type FakeLiveTask struct {
	AppID           string       `json:"appID"`
	ParentRoomID    string       `json:"parentRoomID"`
	RoomID          string       `json:"roomID"`
	Name            string       `json:"name"`
	ClassType       int          `json:"classType"`
	LiveType        int          `json:"liveType"`
	TeacherID       string       `json:"teacherID"`
	TeacherName     string       `json:"teacherName"`
	StartTime       int64        `json:"startTime"`
	EndTime         int64        `json:"endTime"`
	MockChat        bool         `json:"mockChat"`
	IsSelfVideo     bool         `json:"isSelfVideo"`
	OssPath         OssPath      `json:"ossPath"`
	FakeLiveInfo    FakeLiveInfo `json:"fakeLiveInfo"`
	VideoSourceType int          `json:"videoSourceType"`
	VideoID         string       `json:"videoID"`
	Ts              int64        `json:"ts"`
	Kill            bool         `json:"kill"`
	ProcessType     int          `json:"processType"` //0: 停止 1: 开始 2: 刷新
	Width           int          `json:"width"`
	Height          int          `json:"height"`
}

type FakeLiveInfo struct {
	VideoSourceUrl string `json:"videoSourceUrl"` // 视频源地址
	SingleFile     string `json:"singleFile"`     // 单文件
}
