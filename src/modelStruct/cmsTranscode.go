package modelStruct

type FileTranscode struct {
	FileID           int64    `json:"fileId"`
	FileName         string   `json:"fileName"`
	MD5              string   `json:"md5"`
	FileType         string   `json:"fileType"`
	Size             int64    `json:"size"`
	TimeStamp        int64    `json:"timeStamp"`
	TransferPriority int      `json:"transferPriority"`
	UserInfo         UserInfo `json:"userInfo"`
	Cost             int64    `json:"cost"` // 处理时间
	Biz              int      `json:"biz"`  // 业务类型
}

type UserInfo struct {
	Application string       `json:"application"`
	ID          string       `json:"id"`
	Name        string       `json:"name"`
	Mobile      string       `json:"mobile"`
	Username    string       `json:"username"`
	Description *string      `json:"description"` // 使用指针处理可能的null值
	Roles       []string     `json:"roles"`
	RoleDetails []RoleDetail `json:"roleDetails"`
}

type RoleDetail struct {
	ID          string `json:"id"`
	Domain      string `json:"domain"`
	Code        string `json:"code"`
	Name        string `json:"name"`
	Description string `json:"description"`
}

type QueryFilePath struct {
	FileKey  string `json:"fileKey"`
	ParentId int64  `json:"parentId"`
}

type MediaData struct {
	FileName string    `json:"fileName"`
	FileSize int       `json:"fileSize"`
	Cover    string    `json:"cover"`
	Dash     string    `json:"dash"`
	Duration int       `json:"duration"`
	PlayList []FLVItem `json:"playList"`
}

type TransferCodeMeta struct {
	FileType string    `json:"fileType"`
	Data     MediaData `json:"data"`
}

type SaveFileInfo struct {
	FileKey          string           `json:"fileKey"`
	OriginFileId     int64            `json:"originFileId"`
	Percent          int              `json:"percent"`
	TransferCodeMeta TransferCodeMeta `json:"transferCodeMeta"`
}

type MqSaveFileInfo struct {
	FileId           int64            `json:"fileId"`
	TransferStatus   int64            `json:"transferStatus"`
	Percent          int64            `json:"percent"`
	TransferCodeMeta TransferCodeMeta `json:"transferCodeMeta"`
}

type UploadFileInfo struct {
	DirId            int64  `json:"dirId"`
	FileName         string `json:"fileName"`
	FileSize         int    `json:"fileSize"`
	ContentSign      string `json:"contentSign"`
	TransferPriority int    `json:"transferPriority"`
}

type QueryFileInfo struct {
	FileName     string   `json:"fileName"`
	DirPaths     []string `json:"dirPaths"`
	OriginFileId int64    `json:"originFileId"`
	Type         int      `json:"type"`
}
