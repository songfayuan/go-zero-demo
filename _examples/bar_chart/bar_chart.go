package main

import (
	"fmt"
	"log"

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
		width  = 900 // 画布宽度
		height = 700 // 画布高度
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

	barWidth := float64(width-200) / float64(len(data)) // 调整柱子的宽度
	barSpacing := 10.0                                  // 增加柱子之间的间隔
	maxValue := 50.0                                    // 纵坐标最大值

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
		x := 50 + float64(i)*(barWidth+barSpacing)               // x轴起始位置
		y := height - 50 - (d.value / maxValue * (height - 100)) // y轴起始位置减去柱子的高度
		color := colors[i%len(colors)]                           // 循环使用颜色
		dc.SetRGB(color.R, color.G, color.B)
		dc.DrawRectangle(x, y, barWidth, (d.value / maxValue * (height - 100)))
		dc.Fill()
	}

	// 绘制坐标轴
	dc.SetRGB(0, 0, 0)                              // 黑色
	dc.DrawLine(50, height-50, width-50, height-50) // X轴
	dc.DrawLine(50, height-50, 50, 50)              // Y轴
	dc.Stroke()

	// 添加横坐标标签
	for i, d := range data {
		dc.DrawStringAnchored(d.label, 50+float64(i)*(barWidth+barSpacing)+barWidth/2, height-30, 0.5, 1)
	}

	// 添加纵坐标标签
	for i := 0; i <= int(maxValue); i += 5 {
		y := height - 50 - (float64(i) / maxValue * (height - 100))
		dc.DrawStringAnchored(fmt.Sprintf("%d", int(i)), 30, y, 1, 0.5)
	}

	// 保存图像为PNG文件
	return dc.SavePNG(filename)
}
