package main

import (
	"fmt"
	"log"

	"github.com/fogleman/gg"
)

func main() {
	// 定义折线图文件路径
	lineChartFile := "line_chart.png"

	// 创建折线图
	if err := createLineChart(lineChartFile); err != nil {
		log.Fatalf("创建折线图时出错: %v", err)
	}

	fmt.Println("折线图生成成功")
}

// 创建折线图并保存为图片
func createLineChart(filename string) error {
	const (
		width  = 900 // 画布宽度
		height = 600 // 画布高度
	)

	// 数据
	data := []struct {
		label string
		value float64
	}{
		{"身份证", 15},
		{"电话号码", 20},
		{"地址信息", 25},
		{"银行卡号", 30},
		{"财务数据", 35},
		{"基础信息", 40},
		{"户籍信息", 45},
	}

	// 创建画布
	dc := gg.NewContext(width, height)
	dc.SetRGB(1, 1, 1) // 背景色为白色
	dc.Clear()

	barWidth := float64(width-200) / float64(len(data)) // 调整每个数据点之间的间隔
	maxValue := 50.0                                    // 纵坐标最大值

	// 加载自定义字体
	if err := dc.LoadFontFace("_examples/word-template/msyh.ttf", 12); err != nil {
		return fmt.Errorf("无法加载字体: %v", err)
	}

	// 绘制折线图
	dc.SetRGB(0, 0, 0) // 线条颜色为黑色
	for i, d := range data {
		x := 50 + float64(i)*(barWidth+20)                       // x轴起始位置
		y := height - 50 - (d.value / maxValue * (height - 100)) // y轴起始位置减去数据点的高度

		if i == 0 {
			dc.MoveTo(x, y)
		} else {
			dc.LineTo(x, y)
		}
	}
	dc.Stroke()

	// 绘制坐标轴
	dc.SetRGB(0, 0, 0)                              // 黑色
	dc.DrawLine(50, height-50, width-50, height-50) // X轴
	dc.DrawLine(50, height-50, 50, 50)              // Y轴
	dc.Stroke()

	// 添加横坐标标签
	for i, d := range data {
		dc.DrawStringAnchored(d.label, 50+float64(i)*(barWidth+20), height-30, 0.5, 1)
	}

	// 添加纵坐标标签
	for i := 0; i <= int(maxValue); i += 5 {
		y := height - 50 - (float64(i) / maxValue * (height - 100))
		dc.DrawStringAnchored(fmt.Sprintf("%d", int(i)), 30, y, 1, 0.5)
	}

	// 保存图像为PNG文件
	return dc.SavePNG(filename)
}
