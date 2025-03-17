package modelStruct

type Resolution struct {
	Width  int `json:"width"`
	Height int `json:"height"`
}

type VideoStream struct {
	CodecType string `json:"codec_type"`
	Width     int    `json:"width"`
	Height    int    `json:"height"`
}

type MediaFormat struct {
	Streams []VideoStream `json:"streams"`
}
