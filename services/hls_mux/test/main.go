package main

import (
	"fmt"
	"github.com/caovanhoang63/hiholive/services/hls_mux/component/ffmpegc"
)

func main() {
	fmt.Println(ffmpegc.ResolutionCmd(1080, 60))
}
