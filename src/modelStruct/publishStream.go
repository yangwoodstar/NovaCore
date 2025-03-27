package modelStruct

type PublishStreamParams struct {
	Stream PublishStream `json:"stream"`
}

type PublishStream struct {
	StreamID  string `json:"streamID"`
	MediaType int    `json:"mediaType"`
	Video     bool   `json:"video"`
	Audio     bool   `json:"audio"`
	Width     int    `json:"width"`
	Height    int    `json:"height"`
}

type PublishStreamResponse struct {
	SessionID string `json:"sessionID"`
	Rtmp      string `json:"rtmp"`
}
