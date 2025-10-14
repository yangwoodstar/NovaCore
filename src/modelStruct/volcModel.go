package modelStruct

type QueryStream struct {
	UserID     string `json:"UserId"`
	StreamType int    `json:"StreamType"`
}

// RecordFile represents a record file within a record task
type RecordFile struct {
	Vid        string        `json:"Vid"`
	Duration   int64         `json:"Duration"`
	Size       int64         `json:"Size"`
	StartTime  int64         `json:"StartTime"`
	StreamList []QueryStream `json:"StreamList"`
}

// RecordTask represents the record task details
type RecordTask struct {
	StartTime      int64        `json:"StartTime"`
	EndTime        int64        `json:"EndTime"`
	Status         int          `json:"Status"`
	StopReason     string       `json:"StopReason"`
	RecordFileList []RecordFile `json:"RecordFileList"`
}

// Result represents the result containing the record task
type Result struct {
	RecordTask RecordTask `json:"RecordTask"`
}

// ResponseMetadata represents the metadata for the response
type QueryStreamResponseMetadata struct {
	RequestID string `json:"RequestId"`
	Action    string `json:"Action"`
	Version   string `json:"Version"`
	Service   string `json:"Service"`
	Region    string `json:"Region"`
}

type ResponseMetadata struct {
	RequestId string `json:"RequestId"`
	Action    string `json:"Action"`
	Version   string `json:"Version"`
	Service   string `json:"Service"`
	Region    string `json:"Region"`
}

// Response represents the entire response structure
type QueryStreamResponse struct {
	Result           Result           `json:"Result"`
	ResponseMetadata ResponseMetadata `json:"ResponseMetadata"`
}
