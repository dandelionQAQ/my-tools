package main

import (
	"flag"
	"fmt"
	"strconv"
	"strings"

	"github.com/dandelionQAQ/my-tools/screenshot"
)

func main() {
	screen := flag.Int("screen", 0, "屏幕序号")
	region := flag.String("region", "", "截取区域，格式：x,y,width,height")
	output := flag.String("output", "screenshot.png", "保存路径")
	flag.Parse()

	var regionArr []int
	if *region != "" {
		parts := strings.Split(*region, ",")
		if len(parts) != 4 {
			fmt.Println("region格式错误，应为：x,y,width,height")
			return
		}
		for _, p := range parts {
			n, err := strconv.Atoi(strings.TrimSpace(p))
			if err != nil {
				fmt.Println("region包含非数字：", p)
				return
			}
			regionArr = append(regionArr, n)
		}
	}

	err := screenshot.Capture(screenshot.CaptureParams{
		Screen: *screen,
		Region: regionArr,
		Output: *output,
	})
	if err != nil {
		fmt.Println("出错了：", err)
	}
}
