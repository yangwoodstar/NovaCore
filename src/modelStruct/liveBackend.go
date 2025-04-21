package modelStruct

type RoomListRequest struct {
	AppID     string `json:"appId"`
	StartTime int64  `json:"startTime"`
	EndTime   int64  `json:"endTime"`
}

/*
*code	number  非必须
data	object [] 非必须

parentRoomID	string
必须
roomID	string
必须
appID	string
必须
teacherId	string
必须
liveStatus	number
必须
startTime	number
必须
endTime	number
必须
message	string 非必须
trace_id	string 非必须
*/

type RoomListDataResponse struct {
	ParentRoomID string `json:"parentRoomId"`
	RoomID       string `json:"roomId"`
	AppID        string `json:"appId"`
	TeacherID    string `json:"teacherId"`
	LiveStatus   int    `json:"liveStatus"`
	StartTime    int64  `json:"startTime"`
	EndTime      int64  `json:"endTime"`
}

type RoomListResponse struct {
	Code    int                    `json:"code"`
	Data    []RoomListDataResponse `json:"data"`
	Message string                 `json:"message"`
	TraceID string                 `json:"trace_id"`
}

type RoomInfoRequest struct {
	AppID        string `json:"appId"`
	RoomID       string `json:"roomId"`
	ParentRoomID string `json:"parentRoomId"`
	RoomName     string `json:"roomName"`
}

type FLVItem struct {
	Name string `json:"name"` // 必须字段，名称
	URL  string `json:"url"`  // 必须字段，URL 地址
	Size int64  `json:"size"`
	Type int32  `json:"type"`
}

type Dash struct {
	Dash string `json:"dash"` // 必须字段，Dash 地址
}

type PlayBackInfo struct {
	SignalFile     string    `json:"signalFile"` // 必须字段
	SignalFileSize int64     `json:"signalFileSize"`
	Dash           Dash      `json:"dash"` // 必须字段，具体结构根据业务补充
	Flv            []FLVItem `json:"flv"`  // 必须字段，具体结构根据业务补充
}

type ScreenWatermark struct {
	Hidden bool `json:"hidden"` // 必须字段，水印地址
}

type QueryPlaybackInfo struct {
	ParentRoomID     string          `json:"parentRoomId"`
	RoomID           string          `json:"roomId"`
	AppID            string          `json:"appId"`
	ParentVideoID    int             `json:"parentVideoId"`
	Size             int             `json:"size,omitempty"`
	TranscodeDur     int             `json:"transcodeDur,omitempty"`
	Duration         int             `json:"duration,omitempty"`
	FirstFrame       string          `json:"firstFrame,omitempty"`
	ID               int             `json:"id"`
	Name             string          `json:"name"`
	EncodeStatus     int             `json:"encodeStatus"`
	CoursewareWidth  int             `json:"coursewareWidth"`
	CoursewareHeight int             `json:"coursewareHeight"`
	TeacherWidth     int             `json:"teacherWidth"`
	TeacherHeight    int             `json:"teacherHeight"`
	HRatioCourseware int             `json:"hRatioCourseware"`
	HRatioTeacher    int             `json:"hRatioTeacher"`
	Width            int             `json:"width"`
	Height           int             `json:"height"`
	CreateAt         string          `json:"createAt"`
	UpdateAt         string          `json:"updateAt"`
	PlayBackInfo     PlayBackInfo    `json:"playBackInfo"`
	ScreenWatermark  ScreenWatermark `json:"screenWatermark"` // 根据实际情况使用具体结构体或类型
}

type QueryPlaybackInfoResponse struct {
	Code    int               `json:"code"`
	Message string            `json:"message"`
	Data    QueryPlaybackInfo `json:"data"`
}
