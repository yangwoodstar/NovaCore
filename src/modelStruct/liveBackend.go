package modelStruct

type RoomListRequest struct {
	AppID     string `json:"appId"`
	StartTime int64  `json:"startTime"`
	EndTime   int64  `json:"endTime"`
}

/*
*code	number  非必须
data	object [] 非必须

parentRoomID	string
必须
roomID	string
必须
appID	string
必须
teacherId	string
必须
liveStatus	number
必须
startTime	number
必须
endTime	number
必须
message	string 非必须
trace_id	string 非必须
*/

type RoomListDataResponse struct {
	ParentRoomID string `json:"parentRoomId"`
	RoomID       string `json:"roomId"`
	AppID        string `json:"appId"`
	TeacherID    string `json:"teacherId"`
	LiveStatus   int    `json:"liveStatus"`
	StartTime    int64  `json:"startTime"`
	EndTime      int64  `json:"endTime"`
}

type RoomListResponse struct {
	Code    int                    `json:"code"`
	Data    []RoomListDataResponse `json:"data"`
	Message string                 `json:"message"`
	TraceID string                 `json:"trace_id"`
}

type RoomInfoRequest struct {
	AppID        string `json:"appId"`
	RoomID       string `json:"roomId"`
	ParentRoomID string `json:"parentRoomId"`
	RoomName     string `json:"roomName"`
}
