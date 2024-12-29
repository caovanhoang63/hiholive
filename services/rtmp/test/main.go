package main

import (
	"fmt"
	"net"
	"os"
	"os/exec"
	"time"
)

func main() {
	// Giả lập việc truyền tải video
	cmd := exec.Command("ffmpeg", "-re", "-i", "test.flv", "-c", "copy", "-f", "flv", "rtmp://localhost/stream/03b6e5a2-c5f2-11ef-910f-02420a000162")
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr

	// Chạy ffmpeg
	go func() {
		if err := cmd.Run(); err != nil {
			fmt.Printf("ffmpeg error: %v\n", err)
		}
	}()

	// Đợi 5 giây rồi ngắt kết nối (Giả lập lỗi ngắt kết nối)
	time.Sleep(5 * time.Second)

	// Giả lập ngắt kết nối bằng cách đóng kết nối mạng (client)
	conn, err := net.Dial("tcp", "localhost:1935") // Kết nối đến server RTMP
	if err != nil {
		fmt.Println("Error connecting to RTMP server:", err)
		return
	}
	defer conn.Close()

	// Đóng kết nối đột ngột
	conn.Close()
	fmt.Println("Simulated client disconnect")
}
