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

// level, title, time, details, advice, env, message string, owners []string
type WarningInfo struct {
	Level   string   `json:"level"`
	Title   string   `json:"title"`
	Time    string   `json:"time"`
	Details string   `json:"details"`
	Advice  string   `json:"advice"`
	Env     string   `json:"env"`
	Message string   `json:"message"`
	Owners  []string `json:"owners"`
}
