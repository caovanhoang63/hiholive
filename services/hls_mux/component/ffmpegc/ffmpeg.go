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
	DefaultConfig = NewFfmpegConfig("./hls_output")
)

type FfmpegConfig struct {
	mountFolder string
}

func NewFfmpegConfig(mountFolder string) *FfmpegConfig {
	return &FfmpegConfig{
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

func (f *Ffmpeg) NewStream(serverUrl string, key string) {
	log.Println(serverUrl)

	// output folder for HLS file (.m3u8 and .ts)
	outputDir := f.ffmpegConfig.mountFolder + "/" + key
	outputFile := outputDir + "/index-%v.m3u8"
	url := serverUrl + "/" + key
	// create folder if not existed
	err := os.MkdirAll(outputDir, os.ModePerm)
	if err != nil {
		log.Fatalf("Lỗi khi tạo thư mục: %v", err)
	}

	sigChan := make(chan os.Signal, 1)
	signal.Notify(sigChan, syscall.SIGINT, syscall.SIGTERM)

	param, err := ResolutionCmd(1080, 60)
	if err != nil {
		return
	}
	args := []string{
		"-i", url,
		"-async", "1",
		"-crf", "28",
		"-ar", "44100",
		"-sws_flags", "bilinear",
		"-preset", "ultrafast",
		"-tune", "zerolatency",
	}
	args = append(args, param...)
	args = append(args,
		"-threads", "0",
		"-hls_time", "2",
		"-hls_list_size", "6",
		"-hls_flags", "independent_segments",
		"-http_persistent", "0",
		"-f", "hls",
		"-hls_playlist_type", "event",
		"-hls_segment_type", "fmp4", // use fmp4 instead of ts
		"-hls_segment_filename", outputDir+"/%v_%03d.m4s",
		"-master_pl_name", "master.m3u8",
		outputFile,
	)

	cmd := exec.Command("ffmpeg", args...)

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
