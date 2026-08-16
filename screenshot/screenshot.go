package screenshot

import (
	"fmt"
	"image"
	"image/jpeg"
	"image/png"
	"os"
	"path/filepath"

	sc "github.com/kbinani/screenshot"
)

type CaptureParams struct {
	Screen int    `json:"screen"` // 屏幕序号
	Region []int  `json:"region"` // [x, y, width, height]，空表示全屏
	Output string `json:"output"` // 保存路径
}

func Capture(params CaptureParams) error {
	if n := sc.NumActiveDisplays(); params.Screen >= n {
		return fmt.Errorf("屏幕序号 %d 不存在，共 %d 个屏幕", params.Screen, n)
	}

	if params.Output == "" {
		return fmt.Errorf("output不能为空")
	}

	ext := filepath.Ext(params.Output)
	switch ext {
	case ".png", ".jpg", ".jpeg":
	default:
		return fmt.Errorf("不支持的格式：%s，支持png/jpg", ext)
	}

	bounds := sc.GetDisplayBounds(params.Screen)

	var captureRect image.Rectangle
	if len(params.Region) == 4 {
		if params.Region[0]+params.Region[2] > bounds.Dx() ||
			params.Region[1]+params.Region[3] > bounds.Dy() {
			return fmt.Errorf("截屏区域超出屏幕范围")
		}
		if params.Region[0] < 0 || params.Region[1] < 0 {
			return fmt.Errorf("x和y不能为负数")
		}
		x := bounds.Min.X + params.Region[0]
		y := bounds.Min.Y + params.Region[1]
		captureRect = image.Rect(x, y, x+params.Region[2], y+params.Region[3])
	} else {
		captureRect = bounds
	}

	img, err := sc.CaptureRect(captureRect)
	if err != nil {
		return fmt.Errorf("截图失败: %v", err)
	}

	os.MkdirAll(filepath.Dir(params.Output), 0755)
	file, err := os.Create(params.Output)
	if err != nil {
		return fmt.Errorf("创建文件失败：%v", err)
	}
	defer file.Close()

	switch ext {
	case ".png":
		png.Encode(file, img)
	case ".jpg", ".jpeg":
		jpeg.Encode(file, img, nil)
	}

	fmt.Printf("截图成功：%s（%dx%d）\n", params.Output, captureRect.Dx(), captureRect.Dy())
	return nil
}
