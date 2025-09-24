package cloud

import (
	"errors"
	"sync"

	"github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common"
	profile2 "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common/profile"
	trtc "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/trtc/v20190722"
	"github.com/tencentyun/tls-sig-api-v2-golang/tencentyun"
)

type TCClientManager struct {
	AppIDTCClientMap map[uint64]*TCClient
	mu               sync.RWMutex
}

type TCRecordParams struct {
	RecordMode           uint64
	MaxIdleTime          uint64
	StreamType           uint64
	OutputFormat         uint64
	MaxMediaFileDuration uint64
	FillType             uint64
	AudioSampleRate      uint64
	AudioChannels        uint64
	AudioBitrate         uint64
	MixLayoutMode        uint64
	ResourceExpiredHour  uint64
	RoomType             uint64
	FileNamePrefix       []*string
	AudioSubscribeList   []*string
	AudioUnSubscribeList []*string
	VideoSubscribeList   []*string
	VideoUnSubscribeList []*string
}

type TCStorageConfig struct {
	Vendor    uint64
	Bucket    string
	AccessKey string
	SecretKey string
	Region    string
}

type TCStartRecordParams struct {
	RoomID       string
	UserID       string
	RecordParams *TCRecordParams
}

type TCStopRecordParams struct {
	TaskID string
}

type TCClient struct {
	AppID         uint64
	AppSecret     string
	SecretId      string
	SecretKey     string
	Region        string
	StorageConfig *TCStorageConfig
	Client        *trtc.Client
}

func (m *TCClientManager) GetTCClient(appID uint64) (*TCClient, error) {
	m.mu.RLock()
	defer m.mu.RUnlock()
	if tcClient, ok := m.AppIDTCClientMap[appID]; ok {
		return tcClient, nil
	}
	return nil, errors.New("not found")
}

func (m *TCClientManager) AddTCClient(tcClient *TCClient) error {
	m.mu.Lock()
	defer m.mu.Unlock()
	credential := common.NewCredential(tcClient.SecretId, tcClient.SecretKey)
	region := tcClient.Region
	profile := profile2.NewClientProfile()
	client, err := trtc.NewClient(credential, region, profile)
	if err != nil {
		return err
	}
	tcClient.Client = client
	m.AppIDTCClientMap[tcClient.AppID] = tcClient
	return nil
}

func (m *TCClientManager) RemoveTCClient(appID uint64) {
	m.mu.Lock()
	defer m.mu.Unlock()
	delete(m.AppIDTCClientMap, appID)
}

func (t *TCClient) TCStartRecord(params *TCStartRecordParams) (string, error) {
	createRequest := trtc.NewCreateCloudRecordingRequest()
	userSig, err := tencentyun.GenUserSig(int(t.AppID), t.AppSecret, params.UserID, 86400)
	if err != nil {
		return "", err
	}

	privateMapKey, err := tencentyun.GenPrivateMapKeyWithStringRoomID(int(t.AppID), t.AppSecret, params.UserID, 86400, params.RoomID, 63)
	if err != nil {
		return "", err
	}

	createRequest.SdkAppId = &t.AppID
	createRequest.RoomId = &params.RoomID
	createRequest.UserId = &params.UserID
	createRequest.UserSig = &userSig
	createRequest.PrivateMapKey = &privateMapKey
	createRequest.ResourceExpiredHour = &params.RecordParams.ResourceExpiredHour
	createRequest.RoomIdType = &params.RecordParams.RoomType // 字符串房间号

	subscribeStreamUsers := &trtc.SubscribeStreamUserIds{
		SubscribeAudioUserIds:   params.RecordParams.AudioSubscribeList,
		UnSubscribeAudioUserIds: params.RecordParams.AudioUnSubscribeList,
		SubscribeVideoUserIds:   params.RecordParams.VideoSubscribeList,
		UnSubscribeVideoUserIds: params.RecordParams.VideoUnSubscribeList,
	}

	createRequest.RecordParams = &trtc.RecordParams{
		RecordMode:             &params.RecordParams.RecordMode,
		MaxIdleTime:            &params.RecordParams.MaxIdleTime,
		StreamType:             &params.RecordParams.StreamType,
		OutputFormat:           &params.RecordParams.OutputFormat,
		MaxMediaFileDuration:   &params.RecordParams.MaxMediaFileDuration,
		FillType:               &params.RecordParams.FillType,
		SubscribeStreamUserIds: subscribeStreamUsers,
	}

	createRequest.StorageParams = &trtc.StorageParams{
		CloudStorage: &trtc.CloudStorage{
			Vendor:         &t.StorageConfig.Vendor,
			Region:         &t.StorageConfig.Region,
			Bucket:         &t.StorageConfig.Bucket,
			AccessKey:      &t.StorageConfig.AccessKey,
			SecretKey:      &t.StorageConfig.SecretKey,
			FileNamePrefix: params.RecordParams.FileNamePrefix,
		},
	}

	audioParams := trtc.AudioParams{
		SampleRate: &params.RecordParams.AudioSampleRate,
		Channel:    &params.RecordParams.AudioChannels,
		BitRate:    &params.RecordParams.AudioBitrate,
	}

	createRequest.MixTranscodeParams = &trtc.MixTranscodeParams{
		//VideoParams: &videoParams,
		AudioParams: &audioParams,
	}

	createRequest.MixLayoutParams = &trtc.MixLayoutParams{
		MixLayoutMode: &params.RecordParams.MixLayoutMode,
	}

	response, err := t.Client.CreateCloudRecording(createRequest)
	if err != nil {
		return "", err
	}

	return *response.Response.TaskId, nil
}

func (t *TCClient) TCStopRecord(params *TCStopRecordParams) (string, error) {
	deleteRequest := trtc.NewDeleteCloudRecordingRequest()
	deleteRequest.SdkAppId = &t.AppID
	deleteRequest.TaskId = &params.TaskID
	deleteResponse, deleteErr := t.Client.DeleteCloudRecording(deleteRequest)
	if deleteErr != nil {
		return "", deleteErr
	}
	return *deleteResponse.Response.TaskId, nil
}
