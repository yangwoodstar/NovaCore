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
