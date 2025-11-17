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

type RecordQueryResponse struct {
	Result Result `json:"Result"`
}

type Result struct {
	RecordTask RecordTask `json:"RecordTask"`
}

type RecordTask struct {
	StartTime int64 `json:"StartTime"` // 录制开始时间戳 (毫秒)
	EndTime   int64 `json:"EndTime"`   // 录制结束时间戳 (毫秒)
	Status    int   `json:"Status"`    // 任务状态
}
