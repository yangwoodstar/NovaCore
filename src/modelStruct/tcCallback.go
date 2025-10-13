package modelStruct

// FileMessage 结构体
type FileMessage struct {
	FileName       string `json:"FileName"`
	UserId         string `json:"UserId"`
	TrackType      string `json:"TrackType"`
	MediaId        string `json:"MediaId"`
	StartTimeStamp int64  `json:"StartTimeStamp"`
	EndTimeStamp   int64  `json:"EndTimeStamp"`
}

// Payload 结构体
type Payload struct {
	Status      int           `json:"Status"`
	FileList    []string      `json:"FileList"`
	FileMessage []FileMessage `json:"FileMessage"`
}

type EventInfo struct {
	RoomId    string  `json:"RoomId"`
	EventTs   int64   `json:"EventTs"`
	EventMsTs int64   `json:"EventMsTs"`
	UserId    string  `json:"UserId"`
	TaskId    string  `json:"TaskId"`
	Payload   Payload `json:"Payload"`
}

type EventCallback struct {
	EventGroupId int       `json:"EventGroupId"`
	EventType    int       `json:"EventType"`
	CallbackTs   int64     `json:"CallbackTs"`
	EventInfo    EventInfo `json:"EventInfo"`
}
