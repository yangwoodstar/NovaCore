package modelStruct

type Task struct {
	RoomID  string
	AppID   string
	RoomIDs []string
	Data    []byte
	Method  string
	Reply   string
	Topic   string
	Index   int
	Role    int
}
