package tools

import (
	"encoding/json"
	"github.com/yangwoodstar/NovaCore/src/constString"
	"github.com/yangwoodstar/NovaCore/src/modelStruct"
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
