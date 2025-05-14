package constString

const (
	LiveBackendRoomInfo           = "/api/room/find"
	LiveBackendSeniorJunior       = "/api/room/sc-list"
	LiveBackendPlaybackID         = "/api/play-back/find"
	LiveBackendCreatePlaybackPath = "/api/play-back/create"
	LiveBackendUpdatePlaybackPath = "/api/video-source/update"
	LiveReportTranscodePath       = "/api/report/playbackConverter"
	LiveBackendGetRoomList        = "/api/room/status-list"
	LiveRecordBackendGetRoomList  = "/api/record/list"
	LiveRecordBackendRoomInfo     = "/api/record/info"
)

const (
	DocConverter      = "/api/report/docConverter"
	MediaConverter    = "/api/report/mediaConverter"
	PlaybackConverter = "/api/report/playbackConverter"
	DisasterEndRecord = "/api/report/disasterEndRecord"
	RtcToken          = "/api/record/token"

	VolcCallback = "/api/record/callback"

	RecordStart  = "/api/record/start"
	RecordStop   = "/api/record/stop"
	RecordFlush  = "/api/record/flush"
	RecordStatus = "/api/record/status"
	RoomInfo     = "/api/record/info"
	RoomList     = "/api/record/list"

	AiSnapShotStart  = "/api/ai/snapshot/start"
	AiSnapShotStop   = "/api/ai/snapshot/stop"
	AiSnapShotUpdate = "/api/ai/snapshot/update"
	AiEnglishPath    = "/ai/detection/face"
)

const (
	AiEnglishFinishPath = "/athena/record/finish"
)

const (
	AiMathPath                 = "/callback/live/upload/result"
	AiMathAutoRecordStopPath   = "/callback/live/recordStop"
	AiMathAutoRecordUploadPath = "/callback/live/recordUploadDone"
)

const (
	DocTable          = "z_doc_converter"
	MediaTable        = "z_media_converter"
	PlaybackTable     = "z_playback_converter"
	RecordStatusTable = "z_record_status"
)

const (
	// 公共响应码
	Success         = 0   // 操作成功
	SuccessInt      = 200 // 操作成功
	FailedInt       = 400 // 操作失败
	UnauthorizedInt = 401 // token 授权验证失败
	ExceptionInt    = 500 // 系统异常
	NotFoundInt     = 404
)

const (
	MaxIdleTime   = 86400
	MaxRecordTime = 21600 // 6 h
	MaxRecordHour = 6     // 6 h
)

const (
	ByteDance = 1
	Agora     = 2
)

const (
	RoomName   = "recordRoom"
	UserRole   = 0
	UserName   = "record"
	UserAvatar = ""
)

const (
	EventUserScreenVideoStreamStart = "UserScreenVideoStreamStart"
	EventUserScreenAudioStreamStart = "UserScreenAudioStreamStart"
	EventUserScreenVideoStreamStop  = "UserScreenVideoStreamStop"
	EventUserScreenAudioStreamStop  = "UserScreenAudioStreamStop"
	EventRecordUploadDone           = "RecordUploadDone"
	EventSnapshotRealTimeData       = "SnapshotRealTimeData"
	EventRecordStopped              = "RecordStopped"
	EventRecordStarted              = "RecordStarted"
)

const (
	Region = "cn-beijing"
)

const (
	PlaybackInfo = "/api/play-back/pb-info"
)
