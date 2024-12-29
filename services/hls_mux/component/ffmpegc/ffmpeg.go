package ffmpegc

import (
	"encoding/json"
	"fmt"
	"github.com/caovanhoang63/hiholive/shared/go/srvctx"
	"github.com/redis/go-redis/v9"
	"golang.org/x/net/context"
	"log"
	"os"
	"os/exec"
	"os/signal"
	"syscall"
)

var (
	DefaultConfig = NewFfmpegConfig("./hls_output", nil)
)

type FfmpegConfig struct {
	mountFolder string
	rd          *redis.Client
}

func NewFfmpegConfig(mountFolder string, rd *redis.Client) *FfmpegConfig {
	return &FfmpegConfig{
		mountFolder: mountFolder,
		rd:          rd,
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

type ResolutionInfo struct {
	Width  int                   `json:"width"`
	Height int                   `json:"height"`
	Fps    map[string]FpsBitRate `json:"fps"`
}

type FpsBitRate struct {
	ABitRate int `json:"aBitRate"`
	VBitRate int `json:"vBitRate"`
}

// NewStream initializes and runs a Ffmpeg process to create an HLS stream with the specified parameters.
func (f *Ffmpeg) NewStream(streamId, serverUrl, streamKey string, fps, resolution int) (closeFunc func() error) {
	results, err := f.ffmpegConfig.rd.MGet(context.Background(),
		"system_setting:STREAM_RESOLUTION_SUPPORT",
		"system_setting:STREAM_RESOLUTION_INFO").Result()
	if err != nil {
		return
	}
	result1, ok1 := results[0].(string) // Casting the first result
	result2, ok2 := results[1].(string) // Casting the second result
	if !ok1 || !ok2 {
		return
	}

	fmt.Println(streamId, serverUrl, streamKey, fps, resolution)

	supportedMap := map[string][]int{}
	resolutionInfo := map[string]ResolutionInfo{}

	err = json.Unmarshal([]byte(result1), &supportedMap)
	err = json.Unmarshal([]byte(result2), &resolutionInfo)
	if err != nil {
		return
	}

	supported := supportedMap["supported"]

	// output folder for HLS file (.m3u8 and .m4s)
	outputDir := f.ffmpegConfig.mountFolder + "/" + streamId
	outputFile := outputDir + "/index-%v.m3u8"
	url := serverUrl + "/" + streamKey
	// create folder if not existed
	err = os.MkdirAll(outputDir, os.ModePerm)
	if err != nil {
		log.Fatalf("Lỗi khi tạo thư mục: %v", err)
	}

	sigChan := make(chan os.Signal, 1)
	signal.Notify(sigChan, syscall.SIGINT, syscall.SIGTERM)
	param, err := ResolutionCmd(resolution, fps, supported, resolutionInfo)
	if err != nil {
		return
	}
	args := []string{
		"-i", url,

		// Flag to enable audio-video sync. Value "1" allows synchronization
		"-async", "1",

		// CRF (Constant Rate Factor) flag to control video quality. Lower values mean higher quality (0-51). 28 is a mid-to-low quality setting, used to reduce file size.
		"-crf", "28",

		// Audio sample rate (in Hz). This is set to 32000 Hz, typically used for speech or broadcast-quality audio with lower fidelity.
		"-ar", "32000",

		// Flag specifying the interpolation method for resizing images. "bilinear" is a basic interpolation method, providing fast processing but not very sharp results.
		// other option

		//            Speed            Quality
		// bilinear	  Fastest	       Lowest (smooth but not sharp)	        Real-time applications where speed matters more than quality
		// bicubic	  Moderate	       Sharper than bilinear, better quality	General-purpose use when balance between speed and quality is needed
		// lanczos	  Slowest	       Highest (preserves fine details)	    High-quality video encoding where quality is the priority
		// neighbor	  Very Fast	       Most pixelated/blocky	                When speed is a priority and quality can be sacrificed
		// gauss	  Moderate	       Good (slightly soft)	                Balance between speed and quality
		// spline	  Moderate	       High (better than bicubic)	            High-quality applications, where speed is not the main concern
		// doc: https://www.ffmpeg.org/ffmpeg-scaler.html#scaler_005foptions
		"-sws_flags", "bilinear",

		// Specifies the encoding preset. "ultrafast" is the fastest preset, but may result in lower video quality and larger file size.
		"-preset", "ultrafast",

		// Flag to tune encoding for zero-latency, optimizing for real-time streaming and minimal delay.
		"-tune", "zerolatency",
	}

	args = append(args, param...)
	args = append(args,
		"-threads", "2", // Set the number of threads for encoding/decoding (2 threads in this case)
		"-hls_time", "2", // Set the duration (in seconds) of each HLS segment (2 seconds per segment)
		"-hls_list_size", "0", // Set the number of playlist entries (0 means unlimited, keeps all segments in the playlist)
		"-hls_flags", "independent_segments", // Use independent segments, allowing segments to be played individually without depending on previous ones
		"-http_persistent", "0", // Disable HTTP persistent connections (each segment is requested in a new connection)
		"-f", "hls", // Set the output format to HLS (HTTP Live Streaming)
		"-hls_segment_type", "fmp4", // Use fragmented MP4 (fMP4) instead of MPEG-TS for HLS segments
		"-hls_segment_filename", outputDir+"/%v_%03d.m4s", // Define the output filename pattern for HLS segments (e.g., segment1_001.m4s)
		"-master_pl_name", "master.m3u8", // Set the name for the master playlist file
		outputFile, // Output file name (the playlist and segments will be written to this location)
	)

	ctx, cancel := context.WithCancel(context.Background())
	cmd := exec.CommandContext(ctx, "ffmpeg", args...)

	// write log output of FFmpeg
	cmd.Stdout = log.Writer()
	cmd.Stderr = log.Writer()

	go func() {
		<-sigChan
		if err = cmd.Process.Kill(); err != nil {
			log.Fatalf("Failed to kill process: %v", err)
		}
	}()

	go func() {
		if err = cmd.Run(); err != nil {
			log.Println("FFmpeg err : %v", err)
		}
	}()

	return func() error {
		cancel()
		return nil
	}
}

func (f *Ffmpeg) Start() {}
