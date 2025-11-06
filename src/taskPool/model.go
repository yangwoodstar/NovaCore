package taskPool

import "github.com/yangwoodstar/NovaCore/src/transportCore"

type Task struct {
	RoomID string
	AppID  string
	Data   transportCore.UnificationMessage
	Method string
	Reply  string
	Topic  string
	Index  int
}
