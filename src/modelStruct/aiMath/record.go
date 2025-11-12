package aimath

type RecordQuery struct {
	AppID      string `json:"appID"`
	RoomID     string `json:"roomID"`
	TaskID     string `json:"taskID"`
	RecordType int    `json:"recordType"`
}

type ResponseMetadata struct {
	RequestId string `json:"RequestId"`
	Action    string `json:"Action"`
	Version   string `json:"Version"`
	Service   string `json:"Service"`
	Region    string `json:"Region"`
}

type Response struct {
	ResponseMetadata ResponseMetadata `json:"ResponseMetadata"`
	Result           string           `json:"Result"`
}
