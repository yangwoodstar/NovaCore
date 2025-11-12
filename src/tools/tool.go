package tools

import (
	"encoding/base64"
	"fmt"
	"github.com/yangwoodstar/NovaCore/src/constString"
	aimath "github.com/yangwoodstar/NovaCore/src/modelStruct/aiMath"
	"go.uber.org/zap"
	"strings"
	"time"
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
	FileName     string
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
	return fmt.Sprintf("%s/%s/%s/%s/%s/%s", config.FirstPrefix, config.SecondPrefix, config.EnvType, appIDProcess, roomIDProcess, fileName)
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
	if config.FileName != "" {
		fileName = config.FileName
	}
	return fmt.Sprintf("%s/%s/%s/%s/%s/%s/%s/%s/%s", "origin", config.FirstPrefix, config.SecondPrefix, config.EnvType, config.AppID, config.RoomID, config.TaskID, config.TcTaskID, fileName)
}

func Base64EncodeAndReplace(roomNumber string) string {
	// Base64编码
	encoded := base64.StdEncoding.EncodeToString([]byte(roomNumber))

	// 替换字符
	encoded = strings.ReplaceAll(encoded, "/", "-")
	encoded = strings.ReplaceAll(encoded, "=", ".")

	return encoded
}

// Retry 封装重试逻辑
func RetryString(attempts int, sleep time.Duration, fn func() (interface{}, error)) (interface{}, error) {
	var err error
	var res interface{}
	response := aimath.Response{}
	for i := 0; i < attempts; i++ {
		res, err = fn()
		if err == nil {
			GetLogger().Info("Retry", zap.Any("res", res))
			GetLogger().Info("Retry", zap.Any("response", response.Result))
			return res, nil
		} else {
			GetLogger().Error("Retry", zap.Error(err))
		}
		if i < 3 {
			time.Sleep(1 * time.Second)
		}
	}
	return res, err
}
