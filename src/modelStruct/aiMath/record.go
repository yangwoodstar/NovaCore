package aimath

type RecordQuery struct {
	AppID      string `json:"appID"`
	RoomID     string `json:"roomID"`
	TaskID     string `json:"taskID"`
	RecordType int    `json:"recordType"`
}
