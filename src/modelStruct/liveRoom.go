package modelStruct

type Data struct {
	OssPath string `json:"ossPath"`
}
type QueryOssPath struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
	Data    Data   `json:"data"`
}
