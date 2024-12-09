package ffmpegc

import "fmt"

type streamInfo struct {
	width        int
	height       int
	audioBitRate int
	videoBitRate int
	fps          int
	mapFps       int
}

func (s *streamInfo) ToCmd(index int, originResolution int, originFps int) []string {
	if s.height == originResolution {
		return []string{
			"-map", "0:v:0",
			"-map", "0:a:0",
			fmt.Sprintf("-c:v:%d", index), "copy",
			fmt.Sprintf("-c:a:%d", index), "aac",
		}
	}

	return []string{
		"-map", "0:v:0",
		"-map", "0:a:0",
		fmt.Sprintf("-filter:v:%d", index),
		fmt.Sprintf("scale=w=%d:h=%d,fps=%d", s.width, s.height, s.fps),
		fmt.Sprintf("-x264-params:v:%d", index),
		fmt.Sprintf("keyint=%d:scenecut=0", 15),
		fmt.Sprintf("-b:v:%d", index),
		fmt.Sprintf("%dk", vBitRateMap[s.height+s.fps]),
		fmt.Sprintf("-b:a:%d", index),
		fmt.Sprintf("%dk", aBitRateMap[s.height]),
		fmt.Sprintf("-c:v:%d", index), "libx264",
		fmt.Sprintf("-c:a:%d", index), "aac",
	}
}

func ResolutionCmd(resolution, fps int) ([]string, error) {
	if fps%30 != 0 {
		return nil, fmt.Errorf("fps=%d is not a multiple of 25 or 30", fps)
	}

	resolutionIndex := -1
	for i, r := range resolutionSupport {
		if resolution == r {
			resolutionIndex = i
		}
	}

	if resolution > resolutionSupport[len(resolutionSupport)-1] {
		resolutionIndex = len(resolutionSupport) - 1
	}

	if resolutionIndex == -1 {
		return nil, fmt.Errorf("invalid resolution %d", resolution)
	}

	streamMap := ""
	cmd := make([]string, 0)
	index := 0

	for i := 0; i <= resolutionIndex; i++ {
		height := resolutionSupport[i]

		for _, f := range resolutionFpsSupport[height] {

			fpsMap := f
			streamMap += fmt.Sprintf("v:%d,a:%d,name:%dp%d ", index, index, height, f)

			info := &streamInfo{
				width:  widthMap[height],
				height: height,
				fps:    f,
				mapFps: fpsMap,
			}
			cmd = append(cmd, info.ToCmd(index, resolution, fps)...)
			index++
		}
	}
	return append(cmd, "-var_stream_map", streamMap), nil
}
