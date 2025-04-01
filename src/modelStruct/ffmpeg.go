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

type Stream struct {
	CodecType string `json:"codec_type"`
}

type MediaFormat struct {
	Streams []Stream `json:"streams"`
}
