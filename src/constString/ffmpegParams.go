package constString

const (
	FFMpeg  = "ffmpeg"
	FFProbe = "ffprobe"
)

const (
	Resolution      = "-vf"
	Scale           = "scale="
	FrameRate       = "-r"
	Profile         = "-profile:v"
	VBitrate        = "-b:v"
	ABitrate        = "-b:a"
	Input           = "-i"
	VideoCodec      = "-c:v"
	AudioCodec      = "-c:a"
	ForceKeyFrame   = "-force_key_frames"
	HlsTime         = "-hls_time"
	HlsListSize     = "-hls_list_size"
	HlsSegmentName  = "-hls_segment_filename"
	Hls             = "hls"
	OutFile         = "-f"
	CopyCodec       = "copy"
	LibX265         = "libx265"
	LibX264         = "libx264"
	CrfDefaultValue = "25"
	Crf             = "-crf"

	AAC                = "aac"
	SampleRate         = "-ar"
	RealTimeLive       = "-movflags"
	RealTimeLiveHeader = "+faststart"
	Preset             = "-preset"
	PresetValue        = "veryfast"
	Ultrafast          = "ultrafast"
	Complex            = "-filter_complex"
	Map                = "-map"
	OutV               = "[outv]"
	OutA               = "[outa]"
	SeekTime           = "-ss"
	VideoFrame         = "-vframes"
	Concat             = "concat"
	Safe               = "-safe"
	Zero               = "0"
	C                  = "-c"
	Copy               = "copy"
	FFlags             = "-fflags"
	GenPts             = "+genpts"
)

const (
	Error               = "error"
	ShowFormat          = "-show_format"
	ShowStreams         = "-show_streams"
	OutputFormat        = "-of"
	JsonFormat          = "json"
	V                   = "-v"
	A                   = "a"
	SelectStreams       = "-select_streams"
	VideoStream         = "v:0"
	VideoStreamPrefix   = "0:v"
	AudioStreamPrefix   = "0:a?"
	Dash                = "dash"
	DashSegment         = "-dash_segment_type"
	DashSegDuration     = "-seg_duration"
	DashFragDuration    = "-frag_duration"
	DashUseTemplate     = "-use_template"
	DashUseTimeline     = "-use_timeline"
	DashSegmentFormat   = "mp4"
	DashSegmentTime     = "-segment_time"
	Duration            = "format=duration"
	SetVideoResolution  = "-s:v"
	ShowCodec           = "stream=codec_type"
	ShowWidth           = "stream=width"
	ShowHeight          = "stream=height"
	Video               = "video"
	ShowEntries         = "-show_entries"
	AdaptationSet       = "-adaptation_sets"
	AdaptationSetParams = "id=0,streams=v id=1,streams=a"
	VideoPix            = "-pix_fmt"
	VideoPixValue       = "yuv420p"
	SetFps              = "-r"
	FpsValue            = "15"
	SetVF               = "-vf"
	VFValue             = "format=yuv420p"
	SetColorPri         = "-color_primaries"
	ColorPriValue       = "bt709"
	SetColorTrc         = "-color_trc"
	ColorTrcValue       = "bt709"
	SetColorSpace       = "-colorspace"
	ColorSpaceValue     = "bt709"
	ShowDuration        = "format=duration"

	AddAudioParams = "-f lavfi -i anullsrc=channel_layout=stereo:sample_rate=44100"
	AudioDuration  = "-shortest"
	AddAudioStream = "1:a"
)

// probe
const (
	OF      = "-of"
	Default = "default=noprint_wrappers=1:nokey=1"
)

const (
	MediaMp4 = ".mp4"
)

const (
	//标清 type = 0
	Resolution320Type        = 360
	Resolution320Name        = "流畅"
	Resolution320pWidth      = 896
	Resolution320pHeight     = 378
	Resolution320BitRateType = "1500k"
	//高清 type = 1
	Resolution540Type        = 540
	Resolution540Name        = "标清"
	Resolution540pWidth      = 1280
	Resolution540pHeight     = 540
	Resolution540BitRateType = "2000k"
	//超清 type = 2
	Resolution720Type        = 720
	Resolution720Name        = "高清"
	Resolution720pWidth      = 1664
	Resolution720pHeight     = 702
	Resolution720BitRateType = "2500k"
	//超清 type = 3
	Resolution1080pBitRateType = "3000k"
	Resolution1080Type         = 1080
	Resolution1080Name         = "超清"
	Resolution1080pWidth       = 1920
	Resolution1080pHeight      = 810
)
