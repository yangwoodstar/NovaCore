package tools

type Media struct {
	RecordFrameRate int32
	RecordWidth     int32
	RecordHeight    int32
	RecordBitrate   int32
	Resolutions     map[int]Resolution
}

type Resolution struct {
	Width     int32
	Height    int32
	Bitrate   int32
	FrameRate int32
}

func GetDefaultMedia() Media {
	return Media{
		RecordFrameRate: 15,
		RecordWidth:     1920,
		RecordHeight:    810,
		RecordBitrate:   2000,
		Resolutions: map[int]Resolution{
			180:  {320, 180, 300, 15},
			360:  {640, 360, 600, 15},
			540:  {960, 540, 1000, 15},
			720:  {1280, 720, 1500, 15},
			1080: {1920, 1080, 2500, 15},
		},
	}
}

func GetMediaParameters(media Media, resolutionLevel int) (int32, int32, int32, int32, int32, int32) {
	width := media.RecordWidth
	height := media.RecordHeight
	bitrate := media.RecordBitrate
	canvasWidth := media.RecordWidth
	canvasHeight := media.RecordHeight
	fps := media.RecordFrameRate

	if res, exists := media.Resolutions[resolutionLevel]; exists {
		width = res.Width
		height = res.Height
		canvasWidth = res.Width
		canvasHeight = res.Height
		bitrate = res.Bitrate
		fps = res.FrameRate
	}

	return width, height, canvasWidth, canvasHeight, bitrate, fps
}
