package main

import (
	"fmt"
	ffmpeg_go "github.com/u2takey/ffmpeg-go"
)

func main() {
	split := ffmpeg_go.Input("038f494c-47f3-4fe7-a3c4-a12a9ce9664e.jpg").VFlip().Split()
	split0, split1 := split.Get("0"), split.Get("1")
	overlayFile := ffmpeg_go.Input("1.jpg").Crop(10, 10, 158, 112)
	err := ffmpeg_go.Concat([]*ffmpeg_go.Stream{
		split0.Trim(ffmpeg_go.KwArgs{"start_frame": 10, "end_frame": 20}),
		split1.Trim(ffmpeg_go.KwArgs{"start_frame": 30, "end_frame": 40})}).
		Overlay(overlayFile.HFlip(), "").
		DrawBox(50, 50, 120, 120, "red", 5).
		Output("oke.jpg").
		OverWriteOutput().
		Run()

	if err != nil {
		fmt.Println(err)
	}
}
