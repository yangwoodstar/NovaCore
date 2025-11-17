package cloud

import (
	"errors"
	"fmt"
	"sync"

	v20190722 "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/trtc/v20190722"
	"github.com/yangwoodstar/NovaCore/src/constString"
	"github.com/yangwoodstar/NovaCore/src/core/instanceAllocator"
	"github.com/yangwoodstar/NovaCore/src/tools"
	"go.uber.org/zap"
)

var tcClientManager *TCClientManager
var once sync.Once
var TcConfigMap = make(map[string]instanceAllocator.AppIDMapConfig)

func InitTCClientManager(tcConfigMap map[string]instanceAllocator.AppIDMapConfig) {
	for appIDStr, tcConfig := range tcConfigMap {
		TcConfigMap[appIDStr] = tcConfig
	}
}

func GetTCClientManager() *TCClientManager {
	once.Do(func() {
		tcClientManager = &TCClientManager{
			AppIDTCClientMap: make(map[uint64]*TCClient),
		}

		for appIDStr, tcConfig := range TcConfigMap {
			appID := tcConfig.TCConfig.AppID
			storageConfig := TCStorageConfig{
				Vendor:    0,
				Bucket:    tcConfig.TCConfig.Bucket,
				AccessKey: tcConfig.TCConfig.AK,
				SecretKey: tcConfig.TCConfig.SK,
				Region:    tcConfig.TCConfig.CosRegion,
			}

			tcClient := TCClient{
				AppID:         appID,
				AppSecret:     tcConfig.TCConfig.AppSecret,
				SecretId:      tcConfig.TCConfig.SecretId,
				SecretKey:     tcConfig.TCConfig.SecretKey,
				Region:        tcConfig.TCConfig.Region,
				StorageConfig: &storageConfig,
			}

			tools.GetLogger().Info("add tc client", zap.Any("tcClient", tcClient))
			err := tcClientManager.AddTCClient(&tcClient)
			if err != nil {
				tools.GetLogger().Error("GetTCClientManager add TCClient", zap.String("error", err.Error()), zap.Uint64("appID", appID), zap.String("appIDStr", appIDStr))
				continue
			}
			tools.GetLogger().Info("GetTCClientManager add TCClient", zap.Uint64("appID", appID))
		}
	})
	return tcClientManager
}

func GetTCClient(appID uint64) (*TCClient, error) {
	return tcClientManager.GetTCClient(appID)
}

func DoStartRecord(prefix, appID, roomID, taskID string, subscribeList, unSubscribeList []*string) (string, error) {
	tcConfig, ok := TcConfigMap[appID]
	if !ok {
		tools.GetLogger().Error("DoStartRecord", zap.String("error", "not found appID"), zap.String("appID", appID), zap.String("roomID", roomID), zap.String("taskID", taskID))
		return "", errors.New(fmt.Sprintf("not found appID %s %s ", appID, roomID))
	}

	origin := "origin"
	recordFirstPrefix := constString.RecordFirstPrefix
	recordSecondPrefix := constString.RecordSecondPrefix
	objType := prefix
	fileNamePrefix := []*string{&origin, &recordFirstPrefix, &recordSecondPrefix, &objType, &appID, &roomID, &taskID}

	recordParam := &TCRecordParams{
		RecordMode:           2,
		MaxIdleTime:          300,
		StreamType:           1,
		OutputFormat:         3,
		MaxMediaFileDuration: 360,
		FillType:             1,
		AudioSampleRate:      1,
		AudioChannels:        2,
		AudioBitrate:         64000,
		MixLayoutMode:        1,
		ResourceExpiredHour:  24,
		RoomType:             0,
		FileNamePrefix:       fileNamePrefix,
		AudioSubscribeList:   subscribeList,
		AudioUnSubscribeList: unSubscribeList,
		VideoSubscribeList:   subscribeList,
		VideoUnSubscribeList: unSubscribeList,
	}

	params := &TCStartRecordParams{
		RoomID:       roomID,
		UserID:       tools.JoinStrings(appID, "_", taskID),
		RecordParams: recordParam,
	}

	instance, err := GetTCClient(tcConfig.TCConfig.AppID)
	if err != nil {
		return "", err
	}
	retry, err := tools.RetryString(3, 1, func() (interface{}, error) {
		return instance.TCStartRecord(params)
	})
	if err != nil {
		return "", err
	}
	tools.GetLogger().Info("DoStartRecord", zap.String("taskID", retry.(string)), zap.String("appID", appID), zap.String("roomID", roomID), zap.String("taskID", taskID))
	return retry.(string), nil
}

func DoStopRecord(appID, tcTaskID string) (map[string]string, string, error) {
	tcConfig, ok := TcConfigMap[appID]
	if !ok {
		return nil, "", errors.New(fmt.Sprintf("not found appID %s %s ", appID, tcTaskID))
	}

	instance, err := GetTCClient(tcConfig.TCConfig.AppID)
	if err != nil {
		return nil, "", err
	}

	params := &TCStopRecordParams{
		TaskID: tcTaskID,
	}
	retry, err := tools.RetryString(3, 1, func() (interface{}, error) {
		return instance.TCStopRecord(params)
	})
	if err != nil {
		return nil, "", err
	}

	tools.GetLogger().Info("DoStopRecord", zap.String("taskID", retry.(string)), zap.String("appID", appID), zap.String("taskID", tcTaskID))
	return nil, tcTaskID, nil
}

func DoDescribeRecord(appID, roomID, taskID string) (*v20190722.DescribeCloudRecordingResponse, error) {
	tcConfig, ok := TcConfigMap[appID]
	if !ok {
		tools.GetLogger().Error("DoDescribeRecord", zap.String("error", "not found appID"), zap.String("appID", appID), zap.String("roomID", roomID), zap.String("taskID", taskID))
		return nil, errors.New(fmt.Sprintf("not found appID %s %s ", appID, roomID))
	}

	instance, err := GetTCClient(tcConfig.TCConfig.AppID)
	if err != nil {
		return nil, err
	}

	params := &TCDescribeRecordParams{
		TaskID: taskID,
	}
	retry, err := tools.RetryString(3, 1, func() (interface{}, error) {
		return instance.TCDescribeRecord(params)
	})
	if err != nil {
		return nil, err
	}

	if retry == nil {
		return nil, errors.New("describe record result is nil")
	}

	tcResponse := retry.(*v20190722.DescribeCloudRecordingResponse)

	if tcResponse.Response == nil || tcResponse.Response.StorageFileList == nil || len(tcResponse.Response.StorageFileList) == 0 {
		return nil, errors.New("describe record result is empty")
	}

	return tcResponse, nil
}
