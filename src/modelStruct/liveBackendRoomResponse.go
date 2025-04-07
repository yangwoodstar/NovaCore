package modelStruct

import "time"

// 主响应结构体
type Response struct {
	Code int      `json:"code"`
	Data RoomData `json:"data"`
}

type RoomData struct {
	ID                uint32     `json:"id"`
	AppID             string     `json:"appId"`
	ParentRoomID      string     `json:"parentRoomId"`
	RoomID            string     `json:"roomId"`
	Name              string     `json:"name"`
	ClassType         int        `json:"classType"`
	LiveType          int        `json:"liveType"`
	LiveStatus        int        `json:"liveStatus"`
	TeacherID         string     `json:"teacherId"`
	StartTime         int64      `json:"startTime"`
	EndTime           int64      `json:"endTime"`
	PlayBackID        PlayBack   `json:"playBackId"`
	RoomConfig        RoomConfig `json:"roomConfig"`
	OSSPath           OSSPath    `json:"ossPath"`
	BjyInfo           BInfo      `json:"bjyInfo"`
	CourseChapterType int        `json:"courseChapterType"`
	PlayBackVideoID   *string    `json:"playBackVideoID"` // 使用指针以处理可能为 null 的情况
	CreateAt          time.Time  `json:"createAt"`
	UpdateAt          time.Time  `json:"updateAt"`
}

type PlayBack struct {
	ID                uint32    `json:"id"`
	ParentRoomID      string    `json:"parentRoomId"`
	AppID             string    `json:"appId"`
	RoomID            string    `json:"roomId"`
	Name              string    `json:"name"`
	EncodeStatus      int       `json:"encodeStatus"`
	EditedStatus      int       `json:"editedStatus"`
	SwapStatus        int       `json:"swapStatus"`
	Size              *int      `json:"size"`              // 使用指针以处理可能为 null 的情况
	TranscodeDuration *int      `json:"transcodeDuration"` // 使用指针以处理可能为 null 的情况
	Duration          *int      `json:"duration"`          // 使用指针以处理可能为 null 的情况
	FirstFrame        *string   `json:"firstFrame"`        // 使用指针以处理可能为 null 的情况
	PlayBackInfo      *string   `json:"playBackInfo"`      // 使用指针以处理可能为 null 的情况
	OriginalInfo      *string   `json:"originalInfo"`      // 使用指针以处理可能为 null 的情况
	CreateAt          time.Time `json:"createAt"`
	UpdateAt          time.Time `json:"updateAt"`
}

type RoomConfig struct {
	VideoBg      Background `json:"videoBg"`
	CoursewareBg Background `json:"coursewareBg"`
	RecordBg     Background `json:"recordBg"`
	Watermark    Watermark  `json:"watermark"`
	BaseInfo     BaseInfo   `json:"baseInfo"`
	HotWords     []string   `json:"hotWords"`
	ID           uint32     `json:"id"`
	RoomID       string     `json:"roomId"`
	CreateAt     time.Time  `json:"createAt"`
	UpdateAt     time.Time  `json:"updateAt"`
}

type Background struct {
	ID       uint32    `json:"id"`
	ImageURL string    `json:"imageUrl"`
	CreateAt time.Time `json:"createAt"`
	UpdateAt time.Time `json:"updateAt"`
}

type Watermark struct {
	ID       uint32    `json:"id"`
	Hidden   int       `json:"hidden"`
	ImageURL string    `json:"imageUrl"`
	Position string    `json:"position"`
	CreateAt time.Time `json:"createAt"`
	UpdateAt time.Time `json:"updateAt"`
}

type BaseInfo struct {
	ID              uint32    `json:"id"`
	SwitchMode      int       `json:"switchMode"`
	MaxParticipants int       `json:"maxParticipants"`
	PreTime         int       `json:"preTime"`
	LiveType        int       `json:"liveType"`
	LiveMode        int       `json:"liveMode"`
	Layout          int       `json:"layout"`
	IsConnect       bool      `json:"isConnect"`
	CreateAt        time.Time `json:"createAt"`
	UpdateAt        time.Time `json:"updateAt"`
}

type OSSPath struct {
	Year    string `json:"year"`
	Grade   string `json:"grade"`
	Subject string `json:"subject"`
	Term    string `json:"term"`
}

type BInfo struct {
	RoomID  string `json:"roomId"`
	RoomURL string `json:"roomURL"`
	Data    struct {
		User string `json:"user"`
	} `json:"data"`
}
