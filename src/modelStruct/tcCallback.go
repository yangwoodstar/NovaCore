package modelStruct

type Payload struct {
	Status int `json:"Status"`
}
type EventInfo struct {
	RoomId    string `json:"RoomId"`
	EventTs   string `json:"EventTs"`
	EventMsTs int64  `json:"EventMsTs"`
	UserId    string `json:"UserId"`
	TaskId    string `json:"TaskId"`
}

type EventCallback struct {
	EventGroupId int       `json:"EventGroupId"`
	EventType    int       `json:"EventType"`
	CallbackTs   int64     `json:"CallbackTs"`
	EventInfo    EventInfo `json:"EventInfo"`
}
