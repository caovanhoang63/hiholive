package ffmpegc

import (
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

	// Thư mục đầu ra cho các file HLS (.m3u8 và .ts)
	outputDir := f.ffmpegConfig.mountFolder + "/" + key
	outputFile := outputDir + "/index.m3u8"
	url := f.ffmpegConfig.rtmpUrl + key
	// Tạo thư mục đầu ra nếu chưa tồn tại
	err := os.MkdirAll(outputDir, os.ModePerm)
	if err != nil {
		log.Fatalf("Lỗi khi tạo thư mục: %v", err)
	}

	// Câu lệnh FFmpeg để chuyển đổi RTMP stream thành HLS với chế độ ghi liên tục
	cmd := exec.Command("ffmpeg", "-i", url,
		"-c:v", "libx264", // Sử dụng codec H.264 cho video
		"-c:a", "aac", // Sử dụng codec AAC cho audio
		"-f", "hls", // Định dạng output là HLS
		"-hls_time", "10", // Độ dài mỗi segment HLS (10 giây)
		"-hls_list_size", "10", // Giữ tối đa 10 segment trong playlist
		"-hls_flags", "delete_segments", // Xóa các segment cũ khi không còn trong playlist
		"-hls_segment_filename", outputDir+"/segment_%03d.ts", // Đặt tên cho các file segment
		outputFile)

	// Ghi log output của FFmpeg để theo dõi
	cmd.Stdout = log.Writer()
	cmd.Stderr = log.Writer()

	// Chạy câu lệnh FFmpeg
	if err := cmd.Run(); err != nil {
		log.Fatalf("Lỗi khi chạy FFmpeg: %v", err)
	}

}

func (f *Ffmpeg) Start() {}
