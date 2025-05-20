package core

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"github.com/volcengine/volc-sdk-golang/base"
	rtc_v20231101 "github.com/volcengine/volc-sdk-golang/service/rtc/v20231101"
	"github.com/yangwoodstar/NovaCore/src/constString"
	"github.com/yangwoodstar/NovaCore/src/modelStruct"
	"github.com/yangwoodstar/NovaCore/src/tools"
	"go.uber.org/zap"
	"strings"
	"sync"
)

type InstanceManager struct {
	AppIDMap  map[string]modelStruct.AppIDMapConfig
	Config    modelStruct.ByteDanceConfig
	instances sync.Map // key -> *ByteDanceInstance
	mu        sync.RWMutex
}

var (
	defaultManager *InstanceManager
	once           sync.Once
)

func InitInstanceManager(appIDMap map[string]modelStruct.AppIDMapConfig, config modelStruct.ByteDanceConfig) {
	defaultManager = &InstanceManager{
		AppIDMap:  appIDMap,
		Config:    config,
		instances: sync.Map{},
	}
}

// GetInstanceManager 获取实例管理器的单例
func GetInstanceManager() *InstanceManager {
	return defaultManager
}

// getConfigForKey 根据 key 获取配置
func (m *InstanceManager) getConfigForKey(liveAppID string) (modelStruct.ByteDanceConfig, error) {
	// 从配置中获取对应的凭证
	// 这里需要根据你的实际配置管理方式来实现
	appConfig, ok := m.AppIDMap[liveAppID]
	if ok == false {
		tools.Logger.Error("getConfigForKey error", zap.String("liveAppID", liveAppID))
		return modelStruct.ByteDanceConfig{}, errors.New("not exist")
	}

	return modelStruct.ByteDanceConfig{
		AK:     m.Config.AK,
		SK:     m.Config.SK,
		AppID:  appConfig.AppID,
		AppKey: appConfig.AppKey,
		Region: m.Config.Region, // 可以根据需要从配置中获取
	}, nil
}

// GetInstance 根据 key 获取或创建实例
func (m *InstanceManager) GetInstance(key string) (*modelStruct.ByteDanceInstance, error) {
	if instance, ok := m.instances.Load(key); ok {
		return instance.(*modelStruct.ByteDanceInstance), nil
	}

	appConfig, err := m.getConfigForKey(key)
	if err != nil {
		return nil, fmt.Errorf("failed to get config for key %s: %w", key, err)
	}
	// 创建新实例
	m.mu.Lock()
	defer m.mu.Unlock()

	// 双重检查，避免并发创建
	if instance, ok := m.instances.Load(key); ok {
		return instance.(*modelStruct.ByteDanceInstance), nil
	}

	// 创建新实例
	instance := &modelStruct.ByteDanceInstance{
		Rtc:    rtc_v20231101.NewInstance(),
		Config: appConfig,
	}

	// 设置凭证
	instance.Rtc.SetCredential(base.Credentials{
		AccessKeyID:     appConfig.AK,
		SecretAccessKey: appConfig.SK,
		Region:          appConfig.Region,
	})

	// 存储实例
	m.instances.Store(key, instance)
	return instance, nil
}

func StartNormalRecordStream(instance *modelStruct.ByteDanceInstance, recordParams modelStruct.RecordParams) (string, error) {
	//taskId := uuid.New()
	streamList := make([]*rtc_v20231101.StartRecordBodyTargetStreamsStreamListItem, 0)
	for _, value := range recordParams.TargetUserList {
		streamInfo := rtc_v20231101.StartRecordBodyTargetStreamsStreamListItem{}
		streamInfo.UserID = value.UserId
		streamInfo.StreamType = &value.StreamType
		streamInfo.Index = &value.Index
		streamList = append(streamList, &streamInfo)
	}

	excludeStreamList := make([]*rtc_v20231101.StartRecordBodyExcludeStreamsStreamListItem, 0)
	for _, value := range recordParams.ExcludeUserList {
		streamInfo := rtc_v20231101.StartRecordBodyExcludeStreamsStreamListItem{}
		streamInfo.UserID = value.UserId
		streamInfo.StreamType = &value.StreamType
		streamInfo.Index = &value.Index
		excludeStreamList = append(excludeStreamList, &streamInfo)
	}

	excludeStreams := rtc_v20231101.StartRecordBodyExcludeStreams{
		StreamList: excludeStreamList,
	}

	targetStreams := rtc_v20231101.StartRecordBodyTargetStreams{
		StreamList: streamList,
	}

	tosConfig := rtc_v20231101.StartRecordBodyStorageConfigTosConfig{
		AccountID: instance.Config.Account,
		Bucket:    instance.Config.Bucket,
	}

	var storageType int32 = 0
	storageConfig := rtc_v20231101.StartRecordBodyStorageConfig{
		Type:      &storageType,
		TosConfig: &tosConfig,
	}

	appIDProcess := strings.ReplaceAll(recordParams.AppID, "-", "")
	appIDProcess = strings.ReplaceAll(appIDProcess, "_", "")
	roomIDProcess := strings.ReplaceAll(recordParams.RoomID, "-", "")
	roomIDProcess = strings.ReplaceAll(roomIDProcess, "_", "")

	prefix := make([]*string, 0)
	firstPrefix := recordParams.FirstPrefix
	secondPrefix := recordParams.SecondPrefix
	thirdPrefix := recordParams.EnvType
	//dateString := util.GetDateString()
	prefix = append(prefix, &firstPrefix)
	prefix = append(prefix, &secondPrefix)
	prefix = append(prefix, &thirdPrefix)
	prefix = append(prefix, &appIDProcess)
	prefix = append(prefix, &roomIDProcess)
	//prefix = append(prefix, &dateString)
	pattern := recordParams.AppID + "_" + recordParams.RoomID + "_" + recordParams.TaskID
	//pattern := roomID + "_" + index + "_" + util.GenerateUUID()
	fileNameConfig := rtc_v20231101.StartRecordBodyFileNameConfig{
		Prefix:  prefix,
		Pattern: &pattern,
	}
	defaultMedia := tools.GetDefaultMedia()
	width, height, canvasWidth, canvasHeight, bitrate, fps := tools.GetMediaParameters(defaultMedia, recordParams.ResolutionLevel)

	encode := rtc_v20231101.StartRecordBodyEncode{
		VideoFps:     &fps,
		VideoWidth:   &width,
		VideoHeight:  &height,
		VideoBitrate: &bitrate,
	}

	//fileType := "FLV"
	fileFormatArray := []*string{&recordParams.FileType}
	fileFormat := rtc_v20231101.StartRecordBodyFileFormatConfig{
		FileFormat: fileFormatArray,
	}
	background := "#FFFFFF"
	recordLayout := rtc_v20231101.StartRecordBodyLayout{
		CustomLayout: &rtc_v20231101.StartRecordBodyLayoutCustomLayout{
			Canvas: &rtc_v20231101.StartRecordBodyLayoutCustomLayoutCanvas{
				Width:      &canvasWidth,
				Height:     &canvasHeight,
				Background: &background,
			},
		},
	}

	var maxRecordTime int32 = constString.MaxRecordTime
	recordControl := rtc_v20231101.StartRecordBodyControl{
		MaxRecordTime: &maxRecordTime,
	}

	var record = int32(recordParams.RecordMod) // 0
	param := &rtc_v20231101.StartRecordBody{
		RecordMode:       &record,
		TargetStreams:    &targetStreams,
		StorageConfig:    storageConfig,
		Control:          &recordControl,
		FileNameConfig:   &fileNameConfig,
		Encode:           &encode,
		FileFormatConfig: &fileFormat,
		Layout:           &recordLayout,
		ExcludeStreams:   &excludeStreams,
	}

	param.AppID = instance.Config.AppID
	param.RoomID = recordParams.RoomID
	param.TaskID = recordParams.TaskID
	jsonStr, err := json.Marshal(param)
	if err != nil {
		tools.Logger.Error("json marshal failed", zap.Error(err))
	}
	tools.Logger.Info("StartRecordStream", zap.String("param", string(jsonStr)))
	startTime := tools.GetTimeStamp()
	resp, statusCode, err := instance.Rtc.StartRecord(context.Background(), param)
	var response string
	if err != nil {
		if resp != nil && resp.ResponseMetadata.Error != nil {
			errStr, _ := json.Marshal(resp.ResponseMetadata.Error)
			tools.Logger.Error("error", zap.String("error", string(errStr)), zap.Int("statusCode", statusCode))
			// 网关返回的错误
			if resp.ResponseMetadata.Error.CodeN != nil && *resp.ResponseMetadata.Error.CodeN != 0 {
				switch *resp.ResponseMetadata.Error.CodeN {
				// InvalidAccessKey
				case 100009:
					tools.Logger.Error("请求的AK不合法")
					return "", err
				// SignatureDoesNotMatch
				case 100010:
					tools.Logger.Error("签名结果不正确")
					return "", err
				}
			} else {
				// 服务端返回的错误
				switch resp.ResponseMetadata.Error.Code {
				case "InvalidParameter":
					tools.Logger.Error("请求的参数错误, 请根据具体Error中的Message提示调整参数")
					return "", err
				}
			}
		} else {
			tools.Logger.Error("error", zap.Error(err), zap.Int("statusCode", statusCode))
			return "", err
		}
	} else {
		b, _ := json.Marshal(resp)
		response = string(b)
		tools.Logger.Info("success", zap.String("resp", string(b)))
	}
	endTime := tools.GetTimeStamp()
	if endTime-startTime > 200 {
		tools.Logger.Info("StopPushStreamToCdn", zap.Int64("timeDiff", endTime-startTime), zap.Any("param", param), zap.String("res", string(response)), zap.Int("resCode", statusCode))
	}
	tools.Logger.Info("StartRecordStream", zap.Int64("timeDiff", endTime-startTime))
	return response, nil
}

func StopNormalRecordStream(ctx context.Context, instance *modelStruct.ByteDanceInstance, roomID, taskID string) (string, error) {
	param := &rtc_v20231101.StopRecordBody{}
	param.AppID = instance.Config.AppID
	param.TaskID = taskID
	param.RoomID = roomID
	jsonStr, err := json.Marshal(param)
	if err != nil {
		return "", err
	}

	tools.Logger.Info("StopNormalRecordStream", zap.String("params", string(jsonStr)))

	resp, statusCode, err := instance.Rtc.StopRecord(ctx, param)
	var response string
	if err != nil {
		if resp != nil && resp.ResponseMetadata.Error != nil {
			errStr, _ := json.Marshal(resp.ResponseMetadata.Error)
			// 网关返回的错误
			if resp.ResponseMetadata.Error.CodeN != nil && *resp.ResponseMetadata.Error.CodeN != 0 {
				switch *resp.ResponseMetadata.Error.CodeN {
				// InvalidAccessKey
				case 100009:
					return "", err
				// SignatureDoesNotMatch
				case 100010:
					return "", err
				}
			} else {
				// 服务端返回的错误
				switch resp.ResponseMetadata.Error.Code {
				case "InvalidParameter":
					return "", err
				}
			}
			tools.Logger.Error("error", zap.String("error", string(errStr)), zap.Int("statusCode", statusCode))
		} else {
			return "", err
		}
	} else {
		b, _ := json.Marshal(resp)
		response = string(b)
	}

	return response, nil
}
