package modelStruct

type PageChangeData struct {
	DocLoadTimeout bool   `json:"docLoadTimeout"`
	LoadTime       int    `json:"loadTime"`
	Url            string `json:"url"`
	Uuid           string `json:"uuid"`
}

type PageChangeInfo struct {
	Data PageChangeData `json:"data"`
	Type int            `json:"type"`
}
