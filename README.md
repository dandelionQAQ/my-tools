# my-tools

个人工具集，使用 Go 开发，可作为命令行工具使用，也可作为 Go 包导入。

---

## screenshot

截图工具，支持多屏幕、区域截图、多格式输出。

### 安装

```bash
cd screenshot/cmd
go build -o screenshot.exe
```

### 命令行使用

```bash
# 全屏截图（默认屏幕0，保存为screenshot.png）
./screenshot.exe

# 指定屏幕和保存路径
./screenshot.exe --screen 0 --output D:/screenshots/test.png

# 区域截图（从坐标100,200开始，截取400x300区域）
./screenshot.exe --region 100,200,400,300 --output partial.jpg
```

### 参数说明

| 参数 | 默认值 | 说明 |
|------|--------|------|
| `--screen` | `0` | 屏幕序号 |
| `--region` | 空（全屏） | 截取区域，格式：`x,y,width,height` |
| `--output` | `screenshot.png` | 保存路径，支持 `.png` `.jpg` `.jpeg` |

### 作为 Go 包使用

```go
import "github.com/dandelionQAQ/my-tools/screenshot"

// 全屏截图
err := screenshot.Capture(screenshot.CaptureParams{
    Screen: 0,
    Output: "test.png",
})

// 区域截图
err := screenshot.Capture(screenshot.CaptureParams{
    Screen: 0,
    Region: []int{100, 200, 400, 300},
    Output: "partial.jpg",
})
```

### 项目结构

```
screenshot/
├── screenshot.go    // 核心逻辑
├── cmd/
│   └── main.go      // 命令行入口
├── go.mod
└── go.sum
```
