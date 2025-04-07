package modelStruct

type At struct {
	AtMobiles []string `json:"atMobiles"`
	IsAtAll   bool     `json:"isAtAll"`
}

type Markdown struct {
	Title string `json:"title"`
	Text  string `json:"text"`
}

type WebhookMessage struct {
	Msgtype  string   `json:"msgtype"`
	Markdown Markdown `json:"markdown"`
	At       At       `json:"at"`
}
