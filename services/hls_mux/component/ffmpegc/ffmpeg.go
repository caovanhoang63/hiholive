package ffmpegc

import (
	"fmt"
	"github.com/caovanhoang63/hiholive/shared/go/srvctx"
	"log"
	"os"
	"os/exec"
)

var (
	DefaultConfig = NewFfmpegConfig("rtmp://localhost:1935/stream/", "./hls_output")
)

type FfmpegConfig struct {
	rtmpUrl     string
	mountFolder string
}

func NewFfmpegConfig(rtmpUrl, mountFolder string) *FfmpegConfig {
	return &FfmpegConfig{
		rtmpUrl:     rtmpUrl,
		mountFolder: mountFolder,
	}
}

type Ffmpeg struct {
	serviceContext srvctx.ServiceContext
	ffmpegConfig   *FfmpegConfig
}

func NewFfmpeg(serviceContext srvctx.ServiceContext) *Ffmpeg {
	return &Ffmpeg{}
}

func (f *Ffmpeg) WithConfig(config *FfmpegConfig) *Ffmpeg {
	if config == nil {
		config = DefaultConfig
	}
	f.ffmpegConfig = config
	return f
}

func (f *Ffmpeg) NewStream(key string) {
	log.Println(key)

	// output folder for HLS file (.m3u8 and .ts)
	outputDir := f.ffmpegConfig.mountFolder + "/" + key
	outputFile := outputDir + "/index-%v.m3u8"
	url := f.ffmpegConfig.rtmpUrl + key
	// create folder if not existed
	err := os.MkdirAll(outputDir, os.ModePerm)
	if err != nil {
		log.Fatalf("Lỗi khi tạo thư mục: %v", err)
	}

	//ffmpeg -f flv -i "rtmp://server/live/livestream" \
	//-map 0:v:0 -map 0:a:0 -map 0:v:0 -map 0:a:0 -map 0:v:0 -map 0:a:0 \
	//-c:v libx264 -crf 22 -c:a aac -ar 44100 \
	//-filter:v:0 scale=w=480:h=360  -maxrate:v:0 600k -b:a:0 500k \
	//-filter:v:1 scale=w=640:h=480  -maxrate:v:1 1500k -b:a:1 1000k \
	//-filter:v:2 scale=w=1280:h=720 -maxrate:v:2 3000k -b:a:2 2000k \
	//-var_stream_map "v:0,a:0,name:360p v:1,a:1,name:480p v:2,a:2,name:720p" \
	//-preset fast -hls_list_size 10 -threads 0 -f hls \
	//-hls_time 3 -hls_flags independent_segments \
	//-master_pl_name "livestream.m3u8" \
	//-y "livestream-%v.m3u8"

	cmd := exec.Command("ffmpeg",
		"-i", url,
		"-map", "0:v:0",
		"-map", "0:a:0",
		"-map", "0:v:0",
		"-map", "0:a:0",
		"-map", "0:v:0",
		"-map", "0:a:0",
		"-c:v", "libx264",
		"-c:a", "aac",
		"-crf", "22",
		"-ar", "44100",

		"-filter:v:0", "scale=w=480:h=360",
		"-b:v:0", "600k",
		"-b:a:0", "500k",
		"-filter:v:1", "scale=w=640:h=480",
		"-b:v:1", "1500k",
		"-b:a:1", "1000k",
		"-filter:v:2", "scale=w=1280:h=720",
		"-b:v:2", "3000k",
		"-b:a:2", "2000k",
		"-var_stream_map", "v:0,a:0,name:360p v:1,a:1,name:480p v:2,a:2,name:720p",
		"-threads", "0",
		"-hls_time", "5",
		"-hls_list_size", "10",
		"-hls_flags", "independent_segments",
		"-f", "hls",
		"-hls_segment_filename", outputDir+"/stream_%v_%03d.ts",
		"-master_pl_name", "master.m3u8",
		outputFile)

	fmt.Println(cmd.String())
	// write log output of FFmpeg
	cmd.Stdout = log.Writer()
	cmd.Stderr = log.Writer()

	if err := cmd.Run(); err != nil {
		log.Fatalf("FFmpeg err : %v", err)
	}

}

func (f *Ffmpeg) Start() {}
