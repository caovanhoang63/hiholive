package ffmpegc

var resolutionSupport = []int{180, 240, 360, 480, 720, 1080, 1440}

var resolutionFpsSupport = map[int][]int{
	1440: {60, 30},
	1080: {60, 30},
	720:  {60, 30},
	480:  {30},
	360:  {30},
	240:  {30},
	180:  {30},
}

var aBitRateMap = map[int]int{
	1440 + 60: 4150, // 2k
	1440 + 30: 3800, // 2k
	1080 + 60: 3000,
	1080 + 30: 2800,
	720 + 60:  2000,
	720 + 30:  1500,
	480 + 30:  1000,
	360 + 30:  750,
	240 + 30:  500,
	180 + 30:  200,
}

var heightMap = map[int]int{
	1440: 2560,
	1080: 1920,
	720:  1280,
	480:  854,
	360:  640,
	240:  426,
	180:  320,
}

var vBitRateMap = map[int]int{
	1440: 192,
	1080: 192,
	720:  128,
	480:  96,
	360:  48,
	240:  48,
	180:  48,
}
