package modelStruct

import rtc_v20231101 "github.com/volcengine/volc-sdk-golang/service/rtc/v20231101"

type ByteDanceConfig struct {
	AK      string
	SK      string
	AppID   string
	AppKey  string
	Region  string
	Bucket  string
	Account string
}

type ByteDanceInstance struct {
	Rtc    *rtc_v20231101.Rtc
	Config ByteDanceConfig
}

type AppIDMapConfig struct {
	AppID  string `mapstructure:"appID"`
	AppKey string `mapstructure:"appKey"`
}

type RecordParams struct {
	AppID           string       `json:"appID"`
	RoomID          string       `json:"roomID"`
	TaskID          string       `json:"taskID"`
	FileType        string       `json:"fileType"`
	TargetUserList  []StreamInfo `json:"targetUserList"`
	ExcludeUserList []StreamInfo `json:"excludeUserList"`
	ResolutionLevel int          `json:"resolutionLevel"`
	RecordMod       int          `json:"recordMod"`
	FirstPrefix     string       `json:"firstPrefix"`
	SecondPrefix    string       `json:"secondPrefix"`
	EnvType         string       `json:"envType"`
}
