package modelStruct

type AIMathCallbackParams struct {
	ClipTaskId     string `json:"clipTaskId"`
	ClipDuration   int64  `json:"clipDuration"`
	ClipFileSize   int64  `json:"clipFileSize"`
	ClipLink       string `json:"clipLink"`
	RtcType        int    `json:"rtcType"` // 1: TRTC, 2: ZEGOCLOUD
	ClipObjectName string `json:"clipObjectName"`
}
