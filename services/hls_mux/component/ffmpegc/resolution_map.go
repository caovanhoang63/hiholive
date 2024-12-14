package ffmpegc

import (
	"fmt"
	"strconv"
)

type streamInfo struct {
	width        int
	height       int
	audioBitRate int
	videoBitRate int
	fps          int
}

// ToCmd generates a slice of FFmpeg command arguments based on stream properties and target height, resolution, and FPS.
func (s *streamInfo) ToCmd(height int, originResolution int, originFps int) []string {
	if s.height == originResolution && s.fps == originFps {
		return []string{
			"-map", "0:v:0",
			"-map", "0:a:0",
			fmt.Sprintf("-c:v:%d", height), "copy",
			fmt.Sprintf("-c:a:%d", height), "aac",
		}
	}

	return []string{
		"-map", "0:v:0",
		"-map", "0:a:0",
		fmt.Sprintf("-filter:v:%d", height),
		fmt.Sprintf("scale=w=%d:h=%d,fps=%d", s.width, s.height, s.fps),
		fmt.Sprintf("-x264-params:v:%d", height),
		fmt.Sprintf("keyint=%d:scenecut=0", 15),
		fmt.Sprintf("-b:v:%d", height),
		fmt.Sprintf("%dk", s.videoBitRate),
		fmt.Sprintf("-b:a:%d", height),
		fmt.Sprintf("%dk", s.audioBitRate),
		fmt.Sprintf("-c:v:%d", height), "libx264",
		fmt.Sprintf("-c:a:%d", height), "aac",
	}
}

func ResolutionCmd(resolution, fps int, supported []int, info1 map[string]ResolutionInfo) ([]string, error) {
	if fps%30 != 0 {
		return nil, fmt.Errorf("fps=%d is  not a multiple of 25 or 30", fps)
	}

	streamMap := ""
	cmd := make([]string, 0)
	index := 0

	fmt.Println(supported)
	for _, v := range supported {
		value := info1[strconv.Itoa(v)]
		for fKey, fValue := range value.Fps {
			fpsV, _ := strconv.Atoi(fKey)
			streamMap += fmt.Sprintf("v:%d,a:%d,name:%dp%d ", index, index, value.Height, fpsV)
			info := &streamInfo{
				width:        value.Width,
				height:       value.Height,
				fps:          fpsV,
				audioBitRate: fValue.ABitRate,
				videoBitRate: fValue.VBitRate,
			}
			cmd = append(cmd, info.ToCmd(index, resolution, fps)...)
			index++
		}

	}
	return append(cmd, "-var_stream_map", streamMap), nil
}
