package modelStruct

// "year":"","grade":"","subject":"","term":""
type OssPath struct {
	Year    string `json:"year"`
	Grade   string `json:"grade"`
	Subject string `json:"subject"`
	Term    string `json:"term"`
}

type TranscodeInfo struct {
	ParentRoomID     string   `json:"parentRoomID"`
	AppID            string   `json:"appID"`
	RoomID           string   `json:"roomID"`
	ID               string   `json:"id"`
	TS               int64    `json:"ts"`
	Force            bool     `json:"force"`
	FileList         []string `json:"fileList"`
	ProcessType      int      `json:"processType"`
	DisasterRecovery bool     `json:"disasterRecovery"`
	Priority         int      `json:"priority"`
	OssPath          OssPath  `json:"ossPath"`
}
