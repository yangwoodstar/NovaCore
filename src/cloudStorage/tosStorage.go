package cloudStorage

import (
	"context"
	"errors"
	"github.com/volcengine/ve-tos-golang-sdk/v2/tos"
	"github.com/yangwoodstar/NovaCore/src/modelStruct"
	"github.com/yangwoodstar/NovaCore/src/tools"
	"go.uber.org/zap"
)

type TosInfo struct {
	EndPoint    string
	AccessKey   string
	SecretKey   string
	Region      string
	Bucket      string
	StsToken    string
	ExpiredTime int64
}

type TosClient struct {
	Client     *tos.ClientV2
	TosContext context.Context
	TosInfo    *TosInfo
	Logger     *zap.Logger
}

func NewInternalTosClient(cms *modelStruct.Credentials, private bool) (*TosClient, error) {
	timeMilli, err := tools.SwitchToTime(cms.ExpiredTime)
	if err != nil {
		return nil, err
	}

	if timeMilli < tools.GetTimeStamp()+600000 {
		return nil, errors.New("expired")
	}

	tosClient := &TosClient{}
	var ctx = context.Background()
	if private == false {
		tosClient.TosInfo.Region = cms.Region
	} else {
		tosClient.TosInfo.Region = cms.InnerEndpoint
	}

	tosClient.TosInfo.Bucket = cms.Bucket
	tosClient.TosInfo.AccessKey = cms.AccessKeyId
	tosClient.TosInfo.SecretKey = cms.SecretAccessKey
	tosClient.TosInfo.ExpiredTime = timeMilli
	tosClient.TosContext = ctx
	tosClient.TosInfo.StsToken = cms.SessionToken

	credential := tos.NewStaticCredentials(tosClient.TosInfo.AccessKey, tosClient.TosInfo.SecretKey)
	// 初始化 Client 时使用 STS Token
	credential.WithSecurityToken(tosClient.TosInfo.StsToken)
	client, err := tos.NewClientV2(tosClient.TosInfo.EndPoint, tos.WithRegion(tosClient.TosInfo.Bucket), tos.WithCredentials(credential), tos.WithMaxRetryCount(3))
	if err != nil {
		return nil, err
	}

	tosClient.Client = client
	return tosClient, nil
}

func NewTosClient(tosInfo *TosInfo, logger *zap.Logger) (*TosClient, error) {
	var ctx = context.Background()
	client, err := tos.NewClientV2(tosInfo.EndPoint, tos.WithRegion(tosInfo.Region), tos.WithCredentials(tos.NewStaticCredentials(tosInfo.AccessKey, tosInfo.SecretKey)), tos.WithMaxRetryCount(3))
	if err != nil {
		logger.Error("Error creating new client", zap.String("error", err.Error()))
		return nil, err
	}

	tosClient := &TosClient{
		Client:     client,
		TosContext: ctx,
		TosInfo:    tosInfo,
		Logger:     logger,
	}
	return tosClient, nil
}

func (tosClient *TosClient) ResumeUploadFile(filePath, objectKey string) error {
	tosClient.Logger.Debug("ResumeUploadFile", zap.String("filePath", filePath), zap.String("objectKey", objectKey))
	// 直接使用文件路径上传文件
	output, err := tosClient.Client.UploadFile(tosClient.TosContext, &tos.UploadFileInput{
		CreateMultipartUploadV2Input: tos.CreateMultipartUploadV2Input{
			Bucket: tosClient.TosInfo.Bucket,
			Key:    objectKey,
		},
		// 上传的文件路径
		FilePath: filePath,
		// 上传时指定分片大小
		PartSize: tos.DefaultPartSize,
		// 分片上传任务并发数量
		TaskNum: 5,
		// 开启断点续传
		EnableCheckpoint: true,
	})
	if err != nil {
		tosClient.Logger.Error("Error uploading file", zap.String("error", err.Error()))
	} else {
		tosClient.Logger.Debug("UploadFile Request ID:", zap.String("RequestID", output.RequestID))
	}
	return err
}

func (tosClient *TosClient) ResumeDownloadFile(filePath, objectKey string) error {
	tosClient.Logger.Debug("ResumeDownloadFile", zap.String("filePath", filePath), zap.String("objectKey", objectKey))
	output, err := tosClient.Client.DownloadFile(tosClient.TosContext, &tos.DownloadFileInput{
		HeadObjectV2Input: tos.HeadObjectV2Input{
			Bucket: tosClient.TosInfo.Bucket,
			Key:    objectKey,
		},
		// 下载的文件路径
		FilePath: filePath,
		// 下载时指定分片大小
		PartSize: tos.DefaultPartSize,
		// 分片下载任务并发数量
		TaskNum: 5,
		// 开启断点续传
		EnableCheckpoint: true,
		// 指定断点续传临时文件路径
		CheckpointFile: "./checkpoint",
		// 数据传输回调
		//DataTransferListener: &listener{},
	})

	if err != nil {
		tosClient.Logger.Error("Error downloading file", zap.String("error", err.Error()))
	} else {
		tosClient.Logger.Info("DownloadFile Request ID:", zap.String("RequestID", output.RequestID))
	}
	return err
}
