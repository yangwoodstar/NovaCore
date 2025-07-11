package liveByteInstance

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/volcengine/volc-sdk-golang/base"
	live "github.com/volcengine/volc-sdk-golang/service/live/v20230101"
	"sync"
	"time"
)

type ByteDanceInstance struct {
	Live   *live.Live
	Config *LiveConfig
}

type LiveConfig struct {
	AK string
	SK string
}

type PullToPushTaskConfig struct {
	Title            string
	StartTime        int32
	EndTime          int32
	CallbackUrl      string
	Type             int32
	CycleMode        int32
	PushUrl          string
	PullUrl          string
	PlayTimes        int32
	Domain           string
	App              string
	Stream           string
	PreDownload      int32
	ContinueStrategy int32
	StartOffset      float32
	EndOffset        float32
}

type RtmpSnapshotConfig struct {
	Domain       string
	App          string
	Stream       string
	Bucket       string
	StorageDir   string
	TimeInterval int32
}

var (
	defaultLiveInstance *ByteDanceInstance
	onceLive            sync.Once
)

func NewInstance(config *LiveConfig) *ByteDanceInstance {
	onceLive.Do(func() {
		liveByteService := live.NewInstance()
		liveByteService.SetCredential(base.Credentials{
			AccessKeyID:     config.AK,
			SecretAccessKey: config.SK,
		})
		defaultLiveInstance = &ByteDanceInstance{
			Live:   liveByteService,
			Config: config,
		}
	})
	return defaultLiveInstance
}

func (instance *ByteDanceInstance) CreatePullToPushTask(pushConfig *PullToPushTaskConfig) (*live.CreatePullToPushTaskRes, error) {
	body := &live.CreatePullToPushTaskBody{
		// 拉流转推任务的名称，默认为空表示不配置任务名称。支持由中文、大小写字母（A - Z、a - z）和数字（0 - 9）组成，长度为 1 到 20 各字符。
		Title: StringPtr(pushConfig.Title),
		// 任务的开始时间，Unix 时间戳，单位为秒。
		// note：
		// 拉流转推任务持续时间最长为 7 天。
		StartTime: pushConfig.StartTime,
		// 任务的结束时间，Unix 时间戳，单位为秒。
		// note：
		// 拉流转推任务持续时间最长为 7 天。
		EndTime: pushConfig.EndTime,
		// 接收拉流转推任务状态回调的地址，最大长度为 512 个字符，默认为空。
		CallbackURL: StringPtr(pushConfig.CallbackUrl),
		// 拉流来源类型，支持的取值及含义如下。
		// <li> 0：直播源； </li>
		// <li> 1：点播视频。 </li>
		Type: pushConfig.Type,
		// 点播视频文件循环播放模式，当拉流来源类型为点播视频时为必选参数，参数取值及含义如下所示。
		// <li> -1：无限次循环，至任务结束； </li>
		// <li> 0：有限次循环，循环次数以 PlayTimes 取值为准； </li>
		// <li> >0：有限次循环，循环次数以 CycleMode 取值为准。 </li>
		CycleMode: Int32Ptr(pushConfig.CycleMode),
		// 推流地址，即直播源或点播视频转推的目标地址。
		DstAddr: StringPtr(pushConfig.PushUrl),
		// 直播源的拉流地址，拉流来源类型为直播源时，为必传参数，最大长度为 1000 个字符。
		//SrcAddr: StringPtr(pushConfig.PullUrl),
		// 点播视频文件循环播放次数，当 CycleMode 取值为 0 时，PlayTimes 取值将作为循环播放次数。
		// note：
		// PlayTimes 为冗余参数，您可以将 PlayTimes 置 0 后直接使用 CycleMode 指定点播视频文件循环播放次数。
		PlayTimes: Int32Ptr(pushConfig.PlayTimes),
		// 推流域名，推流地址（DstAddr）为空时必传；反之，则该参数不生效。
		Domain: StringPtr(pushConfig.Domain),
		// 推流应用名称，推流地址（DstAddr）为空时必传；反之，则该参数不生效。
		App: StringPtr(pushConfig.App),
		// 推流的流名称，推流地址（DstAddr）为空时必传；反之，则该参数不生效。
		Stream: StringPtr(pushConfig.Stream),
		// 是否开启点播预热，开启点播预热后，系统会自动将点播视频文件缓存到 CDN 节点上，当用户请求直播时，可以直播从 CDN 节点获取视频，从而提高直播流畅度。拉流来源类型为点播视频时，参数生效。
		// <li> 0：不开启； </li>
		// <li> 1：开启（默认值）。 </li>
		PreDownload: Int32Ptr(pushConfig.PreDownload),
		// 续播策略，续播策略指转推点播视频进行直播时出现断流并恢复后，如何继续播放的策略，拉流来源类型为点播视频（Type 为 1）时参数生效，支持的取值及含义如下。
		// <li> 0：从断流处续播（默认值）； </li>
		// <li> 1：从断流处+自然流逝时长处续播。 </li>
		ContinueStrategy: Int32Ptr(pushConfig.ContinueStrategy),
		// 群组所属名称，您可以调用 [ListPullToPushGroup](https://www.volcengine.com/docs/6469/1327382) 获取可用的群组。
		// note：
		// <li> 使用主账号调用时，为非必填，默认加入 default 群组，default 群组不存在时会默认创建，并绑定 default 项目。 </li>
		// <li> 使用子账号调用时，为必填。 </li>
		GroupName: StringPtr("default"),
	}

	// 点播文件地址和开始播放、结束播放的时间设置。
	// note：
	// <li> 当 Type 为点播类型时配置生效。 </li>
	// <li> 与 SrcAddrS 和 OffsetS 字段不可同时填写。 </li>
	VodSrcAddrs1 := live.CreatePullToPushTaskBodyVodSrcAddrsItem{
		// 当前点播文件开始播放的时间偏移值，单位为秒。默认为空时表示开始播放时间不进行偏移。
		StartOffset: Float32Ptr(pushConfig.StartOffset),
		// 当前点播文件结束播放的时间偏移值，单位为秒，默认为空时表示结束播放时间不进行偏移。
		// 点播文件地址。
		SrcAddr: pushConfig.PullUrl,
	}

	body.VodSrcAddrs = append(body.VodSrcAddrs, &VodSrcAddrs1)

	bodyData, err := json.Marshal(body)
	if err != nil {
		return nil, err
	}

	fmt.Println("Request Body:", string(bodyData))

	resp, err := instance.Live.CreatePullToPushTask(context.Background(), body)

	if err != nil {
		return nil, err
	}

	return resp, nil
}

func (instance *ByteDanceInstance) UpdatePullToPushTask(pushConfig *PullToPushTaskConfig) (*live.UpdatePullToPushTaskRes, error) {
	body := &live.UpdatePullToPushTaskBody{
		// 拉流转推任务的名称，默认为空表示不配置任务名称。支持由中文、大小写字母（A - Z、a - z）和数字（0 - 9）组成，长度为 1 到 20 各字符。
		Title: StringPtr(pushConfig.Title),
		// 任务的开始时间，Unix 时间戳，单位为秒。
		// note：
		// 拉流转推任务持续时间最长为 7 天。
		StartTime: pushConfig.StartTime,
		// 任务的结束时间，Unix 时间戳，单位为秒。
		// note：
		// 拉流转推任务持续时间最长为 7 天。
		EndTime: pushConfig.EndTime,
		// 接收拉流转推任务状态回调的地址，最大长度为 512 个字符，默认为空。
		CallbackURL: StringPtr(pushConfig.CallbackUrl),
		// 拉流来源类型，支持的取值及含义如下。
		// <li> 0：直播源； </li>
		// <li> 1：点播视频。 </li>
		Type: pushConfig.Type,
		// 点播视频文件循环播放模式，当拉流来源类型为点播视频时为必选参数，参数取值及含义如下所示。
		// <li> -1：无限次循环，至任务结束； </li>
		// <li> 0：有限次循环，循环次数以 PlayTimes 取值为准； </li>
		// <li> >0：有限次循环，循环次数以 CycleMode 取值为准。 </li>
		CycleMode: Int32Ptr(pushConfig.CycleMode),
		// 推流地址，即直播源或点播视频转推的目标地址。
		DstAddr: StringPtr(pushConfig.PushUrl),
		// 直播源的拉流地址，拉流来源类型为直播源时，为必传参数，最大长度为 1000 个字符。
		SrcAddr: StringPtr(pushConfig.PullUrl),
		// 点播视频文件循环播放次数，当 CycleMode 取值为 0 时，PlayTimes 取值将作为循环播放次数。
		// note：
		// PlayTimes 为冗余参数，您可以将 PlayTimes 置 0 后直接使用 CycleMode 指定点播视频文件循环播放次数。
		PlayTimes: Int32Ptr(pushConfig.PlayTimes),
		// 推流域名，推流地址（DstAddr）为空时必传；反之，则该参数不生效。
		Domain: StringPtr(pushConfig.Domain),
		// 推流应用名称，推流地址（DstAddr）为空时必传；反之，则该参数不生效。
		App: StringPtr(pushConfig.App),
		// 推流的流名称，推流地址（DstAddr）为空时必传；反之，则该参数不生效。
		Stream: StringPtr(pushConfig.Stream),
		// 是否开启点播预热，开启点播预热后，系统会自动将点播视频文件缓存到 CDN 节点上，当用户请求直播时，可以直播从 CDN 节点获取视频，从而提高直播流畅度。拉流来源类型为点播视频时，参数生效。
		// <li> 0：不开启； </li>
		// <li> 1：开启（默认值）。 </li>
		PreDownload: Int32Ptr(pushConfig.PreDownload),
		// 续播策略，续播策略指转推点播视频进行直播时出现断流并恢复后，如何继续播放的策略，拉流来源类型为点播视频（Type 为 1）时参数生效，支持的取值及含义如下。
		// <li> 0：从断流处续播（默认值）； </li>
		// <li> 1：从断流处+自然流逝时长处续播。 </li>
		ContinueStrategy: Int32Ptr(pushConfig.ContinueStrategy),
		// 群组所属名称，您可以调用 [ListPullToPushGroup](https://www.volcengine.com/docs/6469/1327382) 获取可用的群组。
		// note：
		// <li> 使用主账号调用时，为非必填，默认加入 default 群组，default 群组不存在时会默认创建，并绑定 default 项目。 </li>
		// <li> 使用子账号调用时，为必填。 </li>
		GroupName: StringPtr("default"),
	}

	// 为拉流转推视频添加的水印配置信息。
	Watermark := live.UpdatePullToPushTaskBodyWatermark{
		// 水印图片字符串，图片最大 2MB，最小 100Bytes，最大分辨率为 1080×1080。图片 Data URL 格式为：data:image/<mediatype>;base64,<data>。
		// <li> mediatype：图片类型，支持 png、jpg、jpeg 格式； </li>
		// <li> data：base64 编码的图片字符串。 </li>
		// 例如，data:image/png;base64,iVBORw0KGg****mCC
		Picture: "",
		// 水平偏移，表示水印左侧边与转码流画面左侧边之间的距离，使用相对比率，取值范围为 [0,1)。
		RelativePosX: 0.1,
		// 垂直偏移，表示水印顶部边与转码流画面顶部边之间的距离，使用相对比率，取值范围为 [0,1)。
		RelativePosY: 0.1,
		// 水印宽度占直播原始画面宽度百分比，支持精度为小数点后两位。
		Ratio: 0.1,
	}

	body.Watermark = &Watermark

	// 点播文件地址和开始播放、结束播放的时间设置。
	// note：
	// <li> 当 Type 为点播类型时配置生效。 </li>
	// <li> 与 SrcAddrS 和 OffsetS 字段不可同时填写。 </li>
	VodSrcAddrs1 := live.UpdatePullToPushTaskBodyVodSrcAddrsItem{
		// 当前点播文件开始播放的时间偏移值，单位为秒。默认为空时表示开始播放时间不进行偏移。
		StartOffset: Float32Ptr(pushConfig.StartOffset),
		// 当前点播文件结束播放的时间偏移值，单位为秒，默认为空时表示结束播放时间不进行偏移。
		EndOffset: Float32Ptr(pushConfig.EndOffset),
		// 点播文件地址。
		SrcAddr: pushConfig.PullUrl,
	}

	body.VodSrcAddrs = append(body.VodSrcAddrs, &VodSrcAddrs1)

	resp, err := instance.Live.UpdatePullToPushTask(context.Background(), body)

	if err != nil {
		return nil, err
	}

	return resp, nil
}

func (instance *ByteDanceInstance) DeletePullToPushTask(taskID string) (*live.DeletePullToPushTaskRes, error) {
	body := &live.DeletePullToPushTaskBody{
		// 任务 ID，任务的唯一标识，您可以通过[获取拉流转推任务列表](https://www.volcengine.com/docs/6469/1126896)接口获取。
		TaskID: taskID,
		// 任务所属的群组名称，您可以通过[获取拉流转推任务列表](https://www.volcengine.com/docs/6469/1126896)接口获取。
		// note：
		// <li> 使用主账号调用时，为非必填。 </li>
		// <li> 使用子账号调用时，为必填。 </li>
		//GroupName: StringPtr("default"),
	}

	bodyData, err := json.Marshal(body)
	if err != nil {
		fmt.Println("bodyData", bodyData)
		return nil, err
	}

	fmt.Println("Request Body:", string(bodyData))

	resp, err := instance.Live.DeletePullToPushTask(context.Background(), body)

	if err != nil {
		return nil, err
	}

	return resp, nil
}

func (instance *ByteDanceInstance) RestartPullToPushTask(taskID string) (*live.RestartPullToPushTaskRes, error) {
	body := &live.RestartPullToPushTaskBody{
		// 任务 ID，任务的唯一标识，您可以通过[获取拉流转推任务列表](https://www.volcengine.com/docs/6469/1126896)接口获取状态为停用的任务 ID。
		TaskID: taskID,
		// 任务所属的群组名称，您可以通过[获取拉流转推任务列表](https://www.volcengine.com/docs/6469/1126896)接口获取。
		// note：
		// <li> 使用主账号调用时，为非必填。 </li>
		// <li> 使用子账号调用时，为必填。 </li>
		GroupName: StringPtr("default"),
	}

	resp, err := instance.Live.RestartPullToPushTask(context.Background(), body)

	if err != nil {
		return nil, err
	}

	return resp, nil
}

func (instance *ByteDanceInstance) StopPullToPushTask(taskID string) (*live.StopPullToPushTaskRes, error) {
	body := &live.StopPullToPushTaskBody{
		// 任务 ID，任务的唯一标识，您可以通过[获取拉流转推任务列表](https://www.volcengine.com/docs/6469/1126896)接口获取状态为未开始或生效中的任务 ID。
		TaskID: taskID,
		// 任务所属的群组名称，您可以通过[获取拉流转推任务列表](https://www.volcengine.com/docs/6469/1126896)接口获取。
		// note：
		// <li> 使用主账号调用时，为非必填。 </li>
		// <li> 使用子账号调用时，为必填。 </li>
		GroupName: StringPtr("default"),
	}

	resp, err := instance.Live.StopPullToPushTask(context.Background(), body)

	if err != nil {
		return nil, err
	}

	return resp, nil
}

func (instance *ByteDanceInstance) ListPullToPushTask(page, size int32, contains string) (*live.ListPullToPushTaskV2Res, error) {
	body := &live.ListPullToPushTaskV2Body{
		// 查询数据的页码，默认为 1，表示查询第一页的数据。
		Page: Int32Ptr(page),
		// 每页显示的数据条数，默认为 20，最大值为 500。
		Size: Int32Ptr(size),
		// 拉流转推任务的名称，不区分大小写，支持模糊查询。 例如，title 取值为 doc 时，则返回任务名称为 docspace、docs、DOC 等 title 中包含 doc 关键词的所有任务列表。
		Title: StringPtr(contains),
	}

	// 群组名称列表，默认为空表示查询所有群组的任务信息。
	GroupNames := []string{"default"}

	body.GroupNames = StringPtrs(GroupNames)

	resp, err := instance.Live.ListPullToPushTaskV2(context.Background(), body)

	if err != nil {
		return nil, err
	}

	return resp, nil
}

func (instance *ByteDanceInstance) CreatePullToPushGroup(projectName, groupName, key, value, category string) (*live.CreatePullToPushGroupRes, error) {
	body := &live.CreatePullToPushGroupBody{
		// 群组名称，支持有中文、大小写字母和数字组成，最大长度为 20 个字符。
		Name: groupName,
		// 为任务群组设置所属项目，您可以在[访问控制-项目列表](https://console.volcengine.com/iam/resourcemanage/project)查看已有项目并对项目进行管理。
		// 项目是火山引擎提供的一种资源管理方式，您可以对不同业务或项目使用的云资源进行分组管理，以实现基于项目的账单查看、子账号授权等功能。
		ProjectName: projectName,
	}

	// 为任务群组设置资源标签。您可以通过资源标签从不同维度对云资源进行分类和聚合管理，如资源分账等场景。
	// note：
	// 如需使用标签进行资源分账，可以在可以在[账单管理-费用标签](https://console.volcengine.com/finance/bill/tag/)处管理启用标签，将对应标签运用到账单明细等数据中。
	Tags1 := live.CreatePullToPushGroupBodyTagsItem{
		// 标签 Key 值。
		Key: key,
		// 标签 Value 值。
		Value: value,
		// 标签类型，支持以下取值。
		// <li> System：系统内置标签； </li>
		// <li> Custom：自定义标签。 </li>
		Category: category,
	}

	body.Tags = append(body.Tags, &Tags1)

	resp, err := instance.Live.CreatePullToPushGroup(context.Background(), body)

	if err != nil {
		return nil, err
	}

	return resp, nil
}

func (instance *ByteDanceInstance) CreateSnapshotPreset(rtmpSnapshotConfig *RtmpSnapshotConfig) {
	body := &live.CreateSnapshotPresetV2Body{
		// 域名空间，即直播流地址的域名所属的域名空间。您可以调用 [ListDomainDetail](https://www.volcengine.com/docs/6469/1126815) 接口或在视频直播控制台的[域名管理](https://console.volcengine.com/live/main/domain/list)页面，查看直播流使用的域名所属的域名空间。
		Vhost: rtmpSnapshotConfig.Domain,
		// 应用名称，取值与直播流地址中 AppName 字段取值相同。支持由大小写字母（A - Z、a - z）、数字（0 - 9）、下划线（_）、短横线（-）和句点（.）组成，长度为 1 到 30 个字符。
		App: rtmpSnapshotConfig.App,
		// 截图配置的生效状态，取值及含义如下所示。
		// <li> 1：（默认值）生效； </li>
		// <li> 0：不生效。 </li>
		Status: Int32Ptr(1),
	}

	// 截图配置的详细参数配置。
	SnapshotPresetConfig := live.CreateSnapshotPresetV2BodySnapshotPresetConfig{
		// 截图间隔时间，单位为秒，默认值为 10，取值范围为正整数。
		Interval: Int32Ptr(rtmpSnapshotConfig.TimeInterval),
	}

	// 图片格式为 JPEG 时的截图参数，开启 JPEG 截图时设置。
	// note：
	// JPEG 截图和 JPG 截图必须开启且只能开启一个。
	JpegParam := live.CreateSnapshotPresetV2BodySnapshotPresetConfigJPEGParam{
		// 当前格式的截图是否开启，取值及含义如下所示。
		// <li> false：（默认值）不开启； </li>
		// <li> true：开启。 </li>
		Enable: BoolPtr(true),
	}

	// 截图存储到 TOS 时的配置。
	// note：
	// TOSParam 和 ImageXParam 配置且配置其中一个。
	TOSParam := live.CreateSnapshotPresetV2BodySnapshotPresetConfigJPEGParamTOSParam{
		// 截图是否使用 TOS 存储，取值及含义如下所示。
		// <li> false：（默认值）不使用； </li>
		// <li> true：使用。 </li>
		Enable: BoolPtr(true),
		// TOS 存储对应的 Bucket。
		// 例如，存储路径为 live-test-tos-example/live/liveapp 时，Bucket 取值为 live-test-tos-example。
		// note：
		// 使用 TOS 存储时 Bucket 为必填项。
		Bucket: StringPtr(rtmpSnapshotConfig.Bucket),
		// ToS 存储对应的 Bucket 下的存储目录，默认为空。
		// 例如，存储位置为 live-test-tos-example/live/liveapp 时，StorageDir 取值为 live/liveapp。
		StorageDir: StringPtr(rtmpSnapshotConfig.StorageDir),
		// 存储方式为实时截图时的存储规则，支持以 {Domain}/{App}/{Stream}/{UnixTimestamp} 样式设置存储规则，支持输入字母、数字、-、!、_、.、* 及占位符。
		// note：
		// 参数 ExactObject 和 OverwriteObject 传且仅传一个。
		//ExactObject: StringPtr("{Domain}/{App}/{Stream}/{UnixTimestamp}"),
		// 存储方式为覆盖截图时的存储规则，支持以 {Domain}/{App}/{Stream} 样式设置存储规则，支持输入字母、数字、-、!、_、.、* 及占位符。
		// note：
		// 参数 ExactObject 和 OverwriteObject 传且仅传一个。
		OverwriteObject: StringPtr("{Domain}/{App}/{Stream}"),
	}

	JpegParam.TOSParam = &TOSParam

	// 截图存储到 veImageX 时的配置。
	// note：
	// TOSParam 和 ImageXParam 配置且配置其中一个。

	SnapshotPresetConfig.JPEGParam = &JpegParam

	// 截图格式为 JPG 时的截图参数，开启 JPG 截图时设置。
	// note：
	// JPEG 截图和 JPG 截图必须开启且只能开启一个。
	//JpgParam := live.CreateSnapshotPresetV2BodySnapshotPresetConfigJpgParam{}

	//SnapshotPresetConfig.JpgParam = &JpgParam

	body.SnapshotPresetConfig = SnapshotPresetConfig

	resp, err := instance.Live.CreateSnapshotPresetV2(context.Background(), body)

	if err != nil {
		fmt.Printf("error %v", err)
	} else {
		fmt.Printf("success %+v", resp)
	}
}

func (instance *ByteDanceInstance) DescribeLiveBatchPushStreamMetrics(vhost, app, startTime, endTime string, aggregation int32) (*live.DescribeLiveBatchPushStreamMetricsRes, error) {
	body := &live.DescribeLiveBatchPushStreamMetricsBody{
		// 推流域名，您可以调用 [ListDomainDetail](https://www.volcengine.com/docs/6469/1126815) 接口或在视频直播控制台的[域名管理](https://console.volcengine.com/live/main/domain/list)页面，查看直播流使用的推流域名。
		Domain: vhost,
		// 应用名称，取值与直播流地址中的 AppName 字段取值相同，查询流粒度数据时必传，且需同时传入 Stream。支持由大小写字母（A - Z、a - z）、数字（0 - 9）、下划线（_）、短横线（-）和句点（.）组成，长度为 1 到 30 个字符。
		// note：
		// 查询流粒度的监控数据时，需同时指定 App 和 Stream 来指定直播流。
		App: StringPtr(app),
		// 流名称，预置与直播流地址中的 StreamName 字段取值相同，查询流粒度数据时必传，且需同时传入 Stream。支持由大小写字母（A - Z、a - z）、数字（0 - 9）、下划线（_）、短横线（-）和句点（.）组成，长度为 1 到 100 个字符。
		// note：
		// 查询流粒度的监控数据时，需同时指定 App 和 Stream 来指定直播流。
		//Stream: StringPtr("example_stream"),
		// 查询的开始时间，RFC3339 格式的时间戳，精度为秒。
		// note：
		// 单次查询最大时间跨度为 1 天，历史查询最大时间范围为 366 天。
		StartTime: startTime,
		// 查询的结束时间，RFC3339 格式的时间戳，精度为秒。
		EndTime: endTime,
		// 数据聚合的时间粒度，单位为秒，支持的时间粒度如下所示。
		// <li> 5：5 秒； </li>
		// <li> 30：30 秒； </li>
		// <li> 60：（默认值）1 分钟。 </li>
		Aggregation: Int32Ptr(aggregation),
		// 数据聚合时间粒度内，动态指标的聚合算法，取值及含义如下所示。
		// <li> max：（默认值）计算聚合时间粒度内的最大值； </li>
		// <li> avg：计算聚合时间粒度内的平均值。 </li>
		AggType: StringPtr("avg"),
	}

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel() // 即使提前返回也确保释放资源

	resp, err := instance.Live.DescribeLiveBatchPushStreamMetrics(ctx, body)

	if err != nil {
		return nil, err
	} else {
		return resp, nil
	}
}
