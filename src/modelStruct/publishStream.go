package modelStruct

type PublishStreamParams struct {
	Stream PublishStream `json:"stream"`
}

type Position struct {
	X int `json:"x"`
	Y int `json:"y"`
}

type TrimProfile struct {
	Name     string   `json:"name"`
	Position Position `json:"position"`
	Width    int      `json:"width"`
	Height   int      `json:"height"`
}

type PublishStreamExt struct {
	TrimProfile []TrimProfile `json:"trimProfile"`
}

type PublishStream struct {
	StreamID  string            `json:"streamID"`
	MediaType int               `json:"mediaType"`
	Video     bool              `json:"video"`
	Audio     bool              `json:"audio"`
	Width     int               `json:"width"`
	Height    int               `json:"height"`
	Ext       *PublishStreamExt `json:"ext"` // 可选字段，可能为nil
}

type PublishStreamResponse struct {
	SessionID string `json:"sessionID"`
	Rtmp      string `json:"rtmp"`
}

type Hidden struct {
	Hidden bool `json:"hidden"`
}
