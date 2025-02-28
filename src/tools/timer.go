package tools

import "time"

func GetTimeStamp() int64 {
	return time.Now().UnixNano() / int64(time.Millisecond)
}
