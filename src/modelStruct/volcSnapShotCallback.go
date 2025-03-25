package modelStruct

type StreamInfo struct {
	Index      uint32 `json:"Index"`      // 流索引
	UserId     string `json:"UserId"`     // 用户 ID
	StreamType uint32 `json:"StreamType"` // 流类型
}

type ImageData struct {
	Stream      StreamInfo `json:"Stream"`                // 流信息
	ObjectKey   string     `json:"ObjectKey,omitempty"`   // 对象键名称（TOS/S3 平台生效）
	VeImageXURI string     `json:"VeImageXUri,omitempty"` // VeImageX 资源标识符（VeImageX 平台生效）
	Format      uint32     `json:"Format"`                // 截图格式
	Width       uint32     `json:"Width"`                 // 截图宽度（像素）
	Height      uint32     `json:"Height"`                // 截图高度（像素）
	Size        uint64     `json:"Size"`                  // 文件大小（字节）
	Timestamp   uint64     `json:"TimeStamp"`             // UNIX 时间戳
}

type ScreenshotConfig struct {
	AppID             string      `json:"AppId"`                       // 应用的唯一标志
	BusinessID        string      `json:"BusinessId"`                  // 业务标识
	RoomID            string      `json:"RoomId"`                      // 房间 ID，是房间的唯一标志
	TaskID            string      `json:"TaskId"`                      // 截图任务 ID
	Bucket            string      `json:"Bucket,omitempty"`            // 存储截图的桶名称（仅 TOS/S3 平台生效）
	VeImageXServiceID string      `json:"VeImageXServiceId,omitempty"` // VeImageX 的服务 ID（仅 VeImageX 平台生效）
	Data              []ImageData `json:"Data"`                        // 截图数据集合
	Domain            string      `json:"Domain"`                      // 存储截图的域名（仅 TOS/S3 平台生效）
}
