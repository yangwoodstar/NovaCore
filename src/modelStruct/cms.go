package modelStruct

import "encoding/json"

type Credentials struct {
	CurrentTime     string `json:"currentTime"`
	ExpiredTime     string `json:"expiredTime"`
	AccessKeyId     string `json:"accessKeyId"`
	SecretAccessKey string `json:"secretAccessKey"`
	Region          string `json:"region"`
	Endpoint        string `json:"endpoint"`
	InnerEndpoint   string `json:"innerEndpoint"`
	Bucket          string `json:"bucket"`
	SessionToken    string `json:"sessionToken"`
}

type CmsHttpResponse struct {
	Code    int             `json:"code"`
	Msg     string          `json:"msg"`
	Time    string          `json:"time"`
	Data    json.RawMessage `json:"data"` // 使用 RawMessage 延迟解析
	TraceID string          `json:"traceId"`
}
