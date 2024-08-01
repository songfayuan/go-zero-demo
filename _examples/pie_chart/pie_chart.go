package main

import (
	"fmt"
	"github.com/fogleman/gg"
	"math"
)

func main() {
	// 定义饼图图片文件路径
	pieChartFile := "pie_chart.png"

	// 创建饼图并保存为图片
	if err := createPieChart(pieChartFile); err != nil {
		fmt.Printf("创建饼图时出错: %v\n", err)
	} else {
		fmt.Println("饼图创建成功")
	}
}

// 创建饼图并保存为图片
func createPieChart(filename string) error {
	const (
		width  = 640
		height = 600
		radius = 200 // 半径
	)

	dc := gg.NewContext(width, height)
	dc.SetRGB(1, 1, 1) // 背景色为白色
	dc.Clear()

	data := []struct {
		label string
		value float64
	}{
		{"身份证", 15}, {"电话号码", 20}, {"地址信息", 25}, {"银行卡号", 30}, {"财务数据", 35}, {"基础信息", 40}, {"户籍信息", 45},
	}

	colors := []struct{ R, G, B float64 }{
		{0.9, 0.3, 0.3}, // 红色
		{0.3, 0.9, 0.3}, // 绿色
		{0.3, 0.3, 0.9}, // 蓝色
		{0.9, 0.9, 0.3}, // 黄色
		{0.9, 0.3, 0.9}, // 紫色
		{0.3, 0.9, 0.9}, // 青色
		{0.9, 0.6, 0.3}, // 橙色
	}

	// 设置字体（选择支持中文的字体）
	if err := dc.LoadFontFace("_examples/word-template/msyh.ttf", 12); err != nil {
		return fmt.Errorf("无法加载字体: %v", err)
	}

	var total float64
	for _, d := range data {
		total += d.value
	}

	startAngle := -math.Pi / 2 // 从顶部开始绘制
	for i, d := range data {
		percentage := d.value / total
		angle := percentage * 2 * math.Pi

		// 设置扇形颜色
		color := colors[i%len(colors)]
		dc.SetRGB(color.R, color.G, color.B)
		dc.DrawArc(width/2, height/2, radius, startAngle, startAngle+angle)
		dc.LineTo(width/2, height/2)
		dc.Fill()

		// 计算标签位置
		midAngle := startAngle + angle/2
		labelX := width/2 + (radius+20)*math.Cos(midAngle)
		labelY := height/2 + (radius+20)*math.Sin(midAngle)

		// 绘制标签和数据
		labelText := fmt.Sprintf("%s: %.0f", d.label, d.value)
		dc.SetRGB(0, 0, 0) // 文字颜色（黑色）
		dc.DrawStringAnchored(labelText, labelX, labelY, 0.5, 0.5)

		startAngle += angle
	}

	return dc.SavePNG(filename) // 保存图像为PNG文件
}
