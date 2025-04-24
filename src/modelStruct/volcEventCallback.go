package modelStruct

import "encoding/json"

type MainEventModel struct {
	EventType  string `json:"EventType"`
	EventData  string `json:"EventData"` // 需要二次解析的嵌套JSON字符串
	EventTime  string `json:"EventTime"`
	EventId    string `json:"EventId"`
	AppId      string `json:"AppId"`
	Version    string `json:"Version"`
	Signature  string `json:"Signature"`
	Noce       string `json:"Noce"`
	Nonce      string `json:"Nonce"`
	BusinessId string `json:"BusinessId,omitempty"`
}

// EventData 字段对应的结构体（需二次解析）
type StopRecordEventDataModel struct {
	AppId          string                    `json:"AppId"`
	BusinessId     string                    `json:"BusinessId,omitempty"`
	RoomId         string                    `json:"RoomId"`
	TaskId         string                    `json:"TaskId"`
	Code           int                       `json:"Code"`
	RecordFileList []StopRecordCallbackModel `json:"RecordFileList"`
}

// 录音文件信息结构体
type StopRecordCallbackModel struct {
	Vid            string `json:"Vid,omitempty"`
	ObjectKey      string `json:"ObjectKey"`
	Duration       int    `json:"Duration"`  // 单位：毫秒
	Size           int    `json:"Size"`      // 单位：字节
	StartTime      int64  `json:"StartTime"` // 毫秒级时间戳
	AudioCodec     string `json:"AudioCodec"`
	EncryptionType int    `json:"EncryptionType"`
}

// 事件数据明细结构体（需二次解析）
type UploadDoneRecordEventDataModel struct {
	AppId        string `json:"AppId"`
	BusinessId   string `json:"BusinessId,omitempty"`
	RoomId       string `json:"RoomId"`
	TaskId       string `json:"TaskId"`
	Code         int    `json:"Code"` // 状态码（0表示成功）
	ErrorMessage string `json:"ErrorMessage"`
}

type ScreenStreamDataModel struct {
	RoomId      string          `json:"RoomId"`
	UserId      string          `json:"UserId"`
	DeviceType  string          `json:"DeviceType"`
	Timestamp   int64           `json:"Timestamp"`
	ExtraInfo   json.RawMessage `json:"ExtraInfo"`
	StreamIndex int             `json:"StreamIndex"`
	Reason      string          `json:"Reason"`
}

type TargetStreamModel struct {
	Index      int32  `json:"Index"`
	UserId     string `json:"UserId"`     //用户 ID，表示这个流所属的用户
	StreamType int32  `json:"StreamType"` // 0 普通音视频流 1 屏幕流
}
