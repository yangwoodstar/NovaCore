package tools

import (
	"fmt"
	"github.com/yangwoodstar/NovaCore/src/constString"
	"strings"
)

type RecordConfig struct {
	AppID        string
	RoomID       string
	TaskID       string
	TcTaskID     string
	RecordType   int
	FirstPrefix  string
	SecondPrefix string
	EnvType      string
	ObjFileName  string
}

func CleanString(origin string) string {
	origin = strings.ReplaceAll(origin, "-", "")
	origin = strings.ReplaceAll(origin, "_", "")
	return origin
}

func GetRecordFileUrl(domain, path string) string {
	return fmt.Sprintf("%s%s", domain, path)
}

func GetRecordFilePath(config *RecordConfig) string {
	var fileType string
	if config.RecordType == constString.AVType {
		fileType = constString.MP4Suffix
	} else if config.RecordType == constString.AudioType {
		fileType = constString.MP3Suffix
	} else if config.RecordType == constString.VideoType {
		fileType = constString.MP4Suffix
	}
	appIDProcess := CleanString(config.AppID)
	roomIDProcess := CleanString(config.RoomID)
	fileName := fmt.Sprintf("%s_%s_%s.%s", config.AppID, config.RoomID, config.TaskID, fileType)
	return fmt.Sprintf("/%s/%s/%s/%s/%s/%s", config.FirstPrefix, config.SecondPrefix, config.EnvType, appIDProcess, roomIDProcess, fileName)
}

func GetTcRecordFilePath(config *RecordConfig) string {
	var fileType string
	if config.RecordType == constString.AVType {
		fileType = constString.MP4Suffix
	} else if config.RecordType == constString.AudioType {
		fileType = constString.MP3Suffix
	} else if config.RecordType == constString.VideoType {
		fileType = constString.MP4Suffix
	}
	withoutSuffix := strings.TrimSuffix(config.ObjFileName, ".m3u8")
	fileName := fmt.Sprintf("%s.%s", withoutSuffix, fileType)
	return fmt.Sprintf("/%s/%s/%s/%s/%s/%s/%s/%s/%s", "origin", config.FirstPrefix, config.SecondPrefix, config.EnvType, config.AppID, config.RoomID, config.TaskID, config.TcTaskID, fileName)
}
