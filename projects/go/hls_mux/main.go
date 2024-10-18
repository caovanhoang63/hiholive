package main

import (
	"context"
	"github.com/joho/godotenv"
	"hiholive/shared/go/uploadprovider"
	"io/ioutil"
	"log"
	"os"
	"os/exec"
	"path/filepath"
)

func main() {
	failList := make([]int, 0)
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatal("Error loading .env file")
	}
	// Load S3 configuration from environment variables
	s3BucketName := os.Getenv("S3_BUCKET_NAME")
	s3Region := os.Getenv("S3_REGION")
	s3ApiKey := os.Getenv("S3_API_KEY")
	s3Secret := os.Getenv("S3_SECRET")
	s3Domain := os.Getenv("S3_DOMAIN")

	// Initialize S3 provider
	s3Provider := uploadprovider.NewS3Provider(
		s3BucketName,
		s3Region,
		s3ApiKey,
		s3Secret,
		s3Domain,
	)

	// RTMP stream URL
	rtmpURL := "rtmp://localhost:1935/stream/test"

	// Local output directory for HLS (.m3u8 and .ts files)
	outputDir := "./hls_output"
	playlistFile := "pipe:1" // The playlist will be piped to stdout
	segmentFileTemplate := outputDir + "/segment_%05d.ts"

	// Create output directory if it doesn't exist
	err = os.MkdirAll(outputDir, os.ModePerm)
	if err != nil {
		log.Fatalf("Error creating directory: %v", err)
	}

	// FFmpeg command to convert RTMP to HLS
	cmd := exec.Command("ffmpeg", "-i", rtmpURL,
		"-c:v", "libx264", // Use H.264 codec for video
		"-c:a", "aac", // Use AAC codec for audio
		"-f", "hls", // HLS output format
		"-hls_time", "10", // Segment duration in seconds
		"-hls_list_size", "10", // Keep last 10 segments in playlist
		"-hls_flags", "delete_segments", // Delete old segments no longer in playlist
		"-hls_segment_filename", segmentFileTemplate, // Store segments locally
		playlistFile) // Pipe the playlist to stdout

	// Create pipes for stdout and stderr
	stdoutPipe, err := cmd.StdoutPipe()
	if err != nil {
		log.Fatalf("Error creating stdout pipe: %v", err)
	}

	// Start FFmpeg
	if err := cmd.Start(); err != nil {
		log.Fatalf("Error starting FFmpeg: %v", err)
	}

	// Concurrently monitor for new segment files and upload them to S3
	go func() {
		for {
			files, err := ioutil.ReadDir(outputDir)
			if err != nil {
				log.Fatalf("Error reading output directory: %v", err)
			}
			cur := len(files) - 1

			if cur == 0 {
				continue
			}
			// just push the last file to S3
			file := files[cur]
			filePath := filepath.Join(outputDir, file.Name())
			fileData, err := ioutil.ReadFile(filePath)
			if err != nil {
				log.Printf("Error reading file %s: %v", filePath, err)
				continue
			}
			// Upload the file to S3
			// TODO: Implement stream key and store it to redis
			dst := "hls/" + file.Name() // Destination path in S3 (e.g., hls/segment_00001.ts)
			ctx := context.Background()
			_, err = s3Provider.SaveFileUploaded(ctx, fileData, dst)
			if err != nil {
				log.Printf("Error uploading file %s to S3: %v", file.Name(), err)
			} else {
				log.Printf("Uploaded %s to S3 successfully.", file.Name())
				failList = append(failList, cur)
			}
		}
	}()

	// Upload the playlist file to S3 as it's being generated
	go func() {
		buf := make([]byte, 4096)
		for {
			n, err := stdoutPipe.Read(buf)
			if err != nil {
				log.Fatalf("Error reading from stdout pipe: %v", err)
			}
			if n == 0 {
				break
			}
			// Upload the playlist (.m3u8) to S3
			ctx := context.Background()
			s3Key := "hls/index.m3u8" // S3 path for playlist file
			_, uploadErr := s3Provider.SaveFileUploaded(ctx, buf[:n], s3Key)
			if uploadErr != nil {
				log.Fatalf("Error uploading playlist to S3: %v", uploadErr)
			}
		}
	}()

	// Wait for FFmpeg to finish
	if err := cmd.Wait(); err != nil {
		log.Fatalf("Error waiting for FFmpeg: %v", err)
	}

	log.Println("Streaming and uploading to S3 completed.")
}
