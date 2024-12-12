package ffmpegc

var resolutionSupport = []int{720}

var resolutionFpsSupport = map[int][]int{
	1080: {60},
	720:  {60},
	480:  {30},
	360:  {30},
	240:  {30},
	180:  {30},
}

var vBitRateMap = map[int]int{
	1080 + 60: 3000,
	720 + 60:  2000,
	480 + 30:  1000,
	360 + 30:  750,
	240 + 30:  500,
	180 + 30:  200,
}

var widthMap = map[int]int{
	1080: 1920,
	720:  1280,
	480:  854,
	360:  640,
	240:  426,
	180:  320,
}

var aBitRateMap = map[int]int{
	1080: 192,
	720:  128,
	480:  96,
	360:  48,
	240:  48,
	180:  48,
}
