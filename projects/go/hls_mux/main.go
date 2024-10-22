package main

import (
	"log"
	"os"
	"os/exec"
)

func main() {

	//Địa chỉ RTMP stream
	rtmpURL := "rtmpc://localhost:1935/stream/test"

	// Thư mục đầu ra cho các file HLS (.m3u8 và .ts)
	outputDir := "./hls_output"
	outputFile := outputDir + "/index.m3u8"

	// Tạo thư mục đầu ra nếu chưa tồn tại
	err := os.MkdirAll(outputDir, os.ModePerm)
	if err != nil {
		log.Fatalf("Lỗi khi tạo thư mục: %v", err)
	}

	// Câu lệnh FFmpeg để chuyển đổi RTMP stream thành HLS với chế độ ghi liên tục
	cmd := exec.Command("ffmpeg", "-i", rtmpURL,
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

	log.Println("Stream RTMP đang được chuyển thành HLS và lưu liên tục tại", outputFile)
}
