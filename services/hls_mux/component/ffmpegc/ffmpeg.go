package ffmpegc

import (
	"github.com/caovanhoang63/hiholive/shared/go/srvctx"
	"log"
	"os"
	"os/exec"
	"os/signal"
	"syscall"
)

var (
	DefaultConfig = NewFfmpegConfig("rtmp://34.126.85.171:1935/stream/", "./hls_output")
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

	sigChan := make(chan os.Signal, 1)
	signal.Notify(sigChan, syscall.SIGINT, syscall.SIGTERM)

	cmd := exec.Command("ffmpeg",
		"-i", url,

		"-map", "0:v:0",
		"-map", "0:a:0",

		"-map", "0:v:0",
		"-map", "0:a:0",

		"-map", "0:v:0",
		"-map", "0:a:0",

		"-map", "0:v:0",
		"-map", "0:a:0",

		"-map", "0:v:0",
		"-map", "0:a:0",

		"-map", "0:v:0",
		"-map", "0:a:0",

		"-async", "1",
		// audio compression
		"-crf", "22",

		// audio frequency
		"-ar", "44100",

		// scale algorithm
		"-sws_flags", "bilinear",

		// Improve RAM efficiency by increasing compression ratio
		"-preset", "veryfast",

		// 180p30
		"-filter:v:0", "scale=w=320:h=180,fps=30",
		"-x264-params:v:0", "keyint=60:scenecut=0",
		"-b:v:0", "250k",
		"-b:a:0", "48k",
		"-c:v:0", "libx264",
		"-c:a:0", "aac",
		"-map", "0:v:0",
		"-map", "0:a:0",
		// 240p30
		"-filter:v:1", "scale=w=426:h=240,fps=30",
		"-x264-params:v:1", "keyint=60:scenecut=0",
		"-b:v:1", "300k",
		"-b:a:1", "48k",
		"-c:v:1", "libx264",
		"-c:a:1", "aac",

		// 360p30
		"-filter:v:2", "scale=w=640:h=360,fps=30",
		"-x264-params:v:2", "keyint=60:scenecut=0",
		"-b:v:2", "600k",
		"-b:a:2", "48k",
		"-c:v:2", "libx264",
		"-c:a:2", "aac",

		// 480p30
		"-filter:v:3", "scale=w=854:h=480,fps=30",
		"-x264-params:v:3", "keyint=60:scenecut=0",
		"-b:v:3", "1500k",
		"-b:a:3", "48k",
		"-c:v:3", "libx264",
		"-c:a:3", "aac",

		// 720p30
		"-filter:v:4", "scale=w=1280:h=720,fps=30",
		"-x264-params:v:4", "keyint=60:scenecut=0",
		"-b:v:4", "2000k",
		"-b:a:4", "64k",
		"-c:v:4", "libx264",
		"-c:a:4", "aac",

		// 720p60
		"-filter:v:5", "scale=w=1280:h=720,fps=60",
		"-x264-params:v:5", "keyint=120:scenecut=0",
		"-b:v:5", "3000k",
		"-b:a:5", "192k",
		"-c:v:5", "libx264",
		"-c:a:5", "aac",

		// origin resolution
		"-c:v:6", "copy",
		"-c:a:6", "aac",

		"-var_stream_map",
		"v:0,a:0,name:180p30 v:1,a:1,name:240p30 v:2,a:2,name:360p30 v:3,a:3,name:480p60"+
			" v:4,a:4,name:720p30 v:5,a:5,name:720p60 v:6,a:6,name:1080p60",

		"-threads", "2",
		"-hls_time", "2",
		"-hls_list_size", "3",
		"-hls_flags", "independent_segments",
		"-f", "hls",

		"-hls_segment_filename", outputDir+"/%v_%03d.ts",
		"-master_pl_name", "master.m3u8",
		outputFile)

	// write log output of FFmpeg
	cmd.Stdout = log.Writer()
	cmd.Stderr = log.Writer()

	go func() {
		<-sigChan
		if err := cmd.Process.Kill(); err != nil {
			log.Fatalf("Failed to kill process: %v", err)
		}
	}()

	if err := cmd.Run(); err != nil {
		log.Fatalf("FFmpeg err : %v", err)
	}

}

func (f *Ffmpeg) Start() {}
