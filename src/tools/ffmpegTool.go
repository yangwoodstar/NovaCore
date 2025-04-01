package tools

import (
	"encoding/json"
	"fmt"
	"github.com/yangwoodstar/NovaCore/src/constString"
	"github.com/yangwoodstar/NovaCore/src/modelStruct"
	"strconv"
	"strings"
)

func GetResolution(filePath string) (*modelStruct.Resolution, string, string, error) {
	var args []string
	args = append(args, constString.V)
	args = append(args, constString.Error)
	args = append(args, constString.SelectStreams)
	args = append(args, constString.VideoStream)
	args = append(args, constString.ShowEntries)
	args = append(args, constString.ShowCodec)
	args = append(args, constString.ShowEntries)
	args = append(args, constString.ShowWidth)
	args = append(args, constString.ShowEntries)
	args = append(args, constString.ShowHeight)
	args = append(args, constString.OutputFormat)
	args = append(args, constString.JsonFormat)
	args = append(args, filePath)
	exec := NewCommandExecutor()
	err := exec.Run(constString.FFProbe, args...)
	if err != nil {
		return nil, "", "", err
	}
	out := exec.Output()
	errOut := exec.StderrOutput()
	mediaFormat := modelStruct.MediaFormat{}
	resolution := modelStruct.Resolution{
		Width:  0,
		Height: 0,
	}

	err = json.Unmarshal([]byte(out), &mediaFormat)
	if err != nil {
		return nil, out, errOut, err
	}

	for _, stream := range mediaFormat.Streams {
		if stream.CodecType == constString.Video {
			resolution.Width = stream.Width
			resolution.Height = stream.Height
		}
	}
	return &resolution, out, errOut, nil
}

func GetVideoDurationArgs(filePath string) []string {
	var args []string
	args = append(args, constString.V)
	args = append(args, constString.Error)
	args = append(args, constString.ShowEntries)
	args = append(args, constString.Duration)
	args = append(args, constString.OF)
	args = append(args, constString.Default)
	args = append(args, filePath)
	return args
}

// GetVideoDuration 获取视频时长
func GetVideoDuration(filePath string) (int64, string, error) {
	args := GetVideoDurationArgs(filePath)
	exec := NewCommandExecutor()
	err := exec.Run(constString.FFProbe, args...)
	if err != nil {
		return 0, "", err
	}

	out := exec.Output()
	errOut := exec.StderrOutput()
	// 转换为浮点数
	durationStr := strings.TrimSpace(out)
	duration, err := strconv.ParseFloat(durationStr, 64)
	if err != nil {
		fmt.Println("Error parsing duration:", err)
		return 0, "", err
	}
	return int64(duration), errOut, nil
}

func HasAudio(filePath string) (bool, string, string, error) {
	// 构造FFprobe参数
	var args []string
	args = append(args, constString.V)
	args = append(args, constString.Error)
	args = append(args, constString.SelectStreams)
	args = append(args, constString.A)
	args = append(args, constString.ShowEntries)
	args = append(args, constString.ShowCodec)
	args = append(args, constString.OF)
	args = append(args, constString.JsonFormat)
	args = append(args, filePath)

	// 创建命令执行器
	exec := NewCommandExecutor()
	err := exec.Run(constString.FFProbe, args...)
	if err != nil {
		return false, "", "", err
	}

	// 获取输出
	out := exec.Output()
	errOut := exec.StderrOutput()

	// 解析JSON响应

	var result modelStruct.MediaFormat
	if err := json.Unmarshal([]byte(out), &result); err != nil {
		return false, out, errOut, err
	}

	// 判断是否有音频流
	hasAudio := len(result.Streams) > 0
	return hasAudio, out, errOut, nil
}
