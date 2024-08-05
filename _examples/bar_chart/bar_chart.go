package main

import (
	"fmt"
	"log"
	"math"

	"github.com/fogleman/gg"
)

func main() {
	// 定义柱状图文件路径
	barChartFile := "bar_chart.png"

	// 创建柱状图
	if err := createBarChart(barChartFile); err != nil {
		log.Fatalf("创建柱状图时出错: %v", err)
	}

	fmt.Println("柱状图生成成功")
}

// 创建柱状图并保存为图片
func createBarChart(filename string) error {
	const (
		width      = 900  // 画布宽度
		height     = 700  // 画布高度
		barWidth   = 60   // 柱子的固定宽度
		barSpacing = 60.0 // 柱子之间的间隔
		margin     = 50   // 边距
	)

	// 数据
	data := []struct {
		label string
		value float64
	}{
		{"身份证", 15},
		{"电话号码", 2000},
		{"地址信息", 2500},
		{"银行卡号", 3000},
		{"财务数据", 3500},
		{"基础信息", 4000},
		{"户籍信息", 4500},
	}

	// 计算数据中的最大值
	var maxValue float64
	for _, d := range data {
		if d.value > maxValue {
			maxValue = d.value
		}
	}
	// 设置纵坐标最大值，稍微高于数据中的最大值，保证柱子不贴顶
	maxValue *= 1.1

	// 创建画布
	dc := gg.NewContext(width, height)
	dc.SetRGB(1, 1, 1) // 背景色为白色
	dc.Clear()

	// 定义颜色
	colors := []struct{ R, G, B float64 }{
		{0.8, 0.2, 0.2}, // 红色
		{0.2, 0.8, 0.2}, // 绿色
		{0.2, 0.2, 0.8}, // 蓝色
		{0.8, 0.8, 0.2}, // 黄色
		{0.8, 0.2, 0.8}, // 紫色
		{0.2, 0.8, 0.8}, // 青色
		{0.8, 0.8, 0.8}, // 灰色
	}

	// 加载自定义字体
	if err := dc.LoadFontFace("_examples/word-template/msyh.ttf", 12); err != nil {
		return fmt.Errorf("无法加载字体: %v", err)
	}

	// 绘制柱状图
	for i, d := range data {
		x := margin + float64(i)*(barWidth+barSpacing)                    // x轴起始位置
		y := height - margin - (d.value / maxValue * (height - 2*margin)) // y轴起始位置减去柱子的高度
		color := colors[i%len(colors)]                                    // 循环使用颜色
		dc.SetRGB(color.R, color.G, color.B)
		dc.DrawRectangle(x, y, barWidth, (d.value / maxValue * (height - 2*margin)))
		dc.Fill()
	}

	// 绘制坐标轴
	dc.SetRGB(0, 0, 0)                                              // 黑色
	dc.DrawLine(margin, height-margin, width-margin, height-margin) // X轴
	dc.DrawLine(margin, height-margin, margin, margin)              // Y轴
	dc.Stroke()

	// 计算纵坐标标签间隔
	interval := calculateInterval(maxValue, 10)

	// 添加横坐标标签
	for i, d := range data {
		dc.DrawStringAnchored(d.label, margin+float64(i)*(barWidth+barSpacing)+barWidth/2, height-margin+20, 0.5, 1)
	}

	// 添加纵坐标标签
	for i := 0.0; i <= maxValue; i += interval {
		y := height - margin - (i / maxValue * (height - 2*margin))
		dc.DrawStringAnchored(fmt.Sprintf("%.0f", i), margin-10, y, 1, 0.5)
	}

	// 保存图像为PNG文件
	return dc.SavePNG(filename)
}

// 计算适当的标签间隔
func calculateInterval(maxValue float64, maxLabels int) float64 {
	interval := maxValue / float64(maxLabels)
	// 向上取整到最近的10的倍数
	magnitude := math.Pow(10, math.Floor(math.Log10(interval)))
	normalized := interval / magnitude
	if normalized > 5 {
		interval = 10 * magnitude
	} else if normalized > 2 {
		interval = 5 * magnitude
	} else {
		interval = 2 * magnitude
	}
	return interval
}
