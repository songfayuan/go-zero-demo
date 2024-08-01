package main

import (
	"bytes"
	"fmt"
	"log"
	"math"
	"strings"

	"github.com/Esword618/unioffice/color"
	"github.com/Esword618/unioffice/common"
	"github.com/Esword618/unioffice/document"
	"github.com/Esword618/unioffice/measurement"
	"github.com/Esword618/unioffice/schema/soo/wml"
	"github.com/fogleman/gg"
)

//教程：https://blog.csdn.net/u011019141/article/details/140788882

//与fill-word-template.go区别：图片不保存到磁盘，仅放在缓冲区

func main() {
	// 定义文档路径和图表文件路径
	docPath := "_examples/word-template/template.docx"
	updatedDocPath := "/Users/songfayuan/Downloads/updated_demo.docx"

	// 打开文档
	doc, err := openDocument(docPath)
	if err != nil {
		log.Fatalf("无法打开文档: %v", err)
	}

	// 填充模板中的变量
	fillTemplate(doc, map[string]string{
		"{{TASK_NAME}}": "任务名称示例",
		"{{DETAILS}}":   "详细信息示例",
	})

	// 在指定标签处插入表格
	if err := insertTableAt(doc, "{{biaoge}}"); err != nil {
		log.Fatalf("插入表格时出错: %v", err)
	}

	// 创建折线图并存储到缓存中
	lineChartBuffer, err := createLineChart()
	if err != nil {
		log.Fatalf("创建图表时出错: %v", err)
	}

	// 在指定标签处插入图表
	if err := insertImageAt(doc, lineChartBuffer, "{{tubiao}}"); err != nil {
		log.Fatalf("插入图表时出错: %v", err)
	}

	// 创建柱状图并存储到缓存中
	barChartBuffer, err := createBarChart()
	if err != nil {
		log.Fatalf("创建柱状图时出错: %v", err)
	}

	// 在指定标签处插入柱状图
	if err := insertImageAt(doc, barChartBuffer, "{{zhuzhuangtu}}"); err != nil {
		log.Fatalf("插入柱状图时出错: %v", err)
	}

	// 创建饼图并存储到缓存中
	pieChartBuffer, err := createPieChart()
	if err != nil {
		log.Fatalf("创建饼图时出错: %v", err)
	}

	// 在指定标签处插入饼图
	if err := insertImageAt(doc, pieChartBuffer, "{{bingtu}}"); err != nil {
		log.Fatalf("插入饼图时出错: %v", err)
	}

	// 删除{{a}}到{{b}}之间的段落
	if err := removeParagraphsBetweenTags(doc, "{{a}}", "{{b}}"); err != nil {
		log.Fatalf("删除段落时出错: %v", err)
	}

	// 删除指定标签
	if err := removeParagraphWithTag(doc, "{{shanchu}}"); err != nil {
		log.Fatalf("删除指定标签时出错: %v", err)
	}

	// 保存更新后的Word文档
	if err := doc.SaveToFile(updatedDocPath); err != nil {
		log.Fatalf("无法保存文档: %v", err)
	}
	fmt.Println("文档更新成功")
}

// 打开文档
func openDocument(path string) (*document.Document, error) {
	return document.Open(path) // 使用unioffice库打开指定路径的文档
}

// 填充模板中的变量
func fillTemplate(doc *document.Document, replacements map[string]string) {
	for _, para := range doc.Paragraphs() { // 遍历文档中的每个段落
		for _, run := range para.Runs() { // 遍历段落中的每个运行（文本片段）
			text := run.Text()
			for placeholder, replacement := range replacements { // 遍历需要替换的占位符
				if strings.Contains(text, placeholder) { // 如果文本包含占位符
					text = strings.ReplaceAll(text, placeholder, replacement) // 替换占位符
					run.Clear()                                               // 清除原有内容
					run.AddText(text)                                         // 添加替换后的文本
				}
			}
		}
	}
}

// 在指定标签处插入表格
func insertTableAt(doc *document.Document, tag string) error {
	paras := doc.Paragraphs() // 获取文档中的所有段落
	for _, para := range paras {
		if paraContainsTag(&para, tag) { // 如果段落包含指定标签
			// 创建并配置表格
			table := doc.InsertTableAfter(para)     // 在标签段落之后插入表格
			table.Properties().SetWidthPercent(100) // 设置表格宽度为100%
			borders := table.Properties().Borders()
			borders.SetAll(wml.ST_BorderSingle, color.Black, measurement.Dxa) // 设置所有边框为单线黑色

			for i := 0; i < 3; i++ { // 创建表格行和单元格
				row := table.AddRow()
				for j := 0; j < 3; j++ {
					cell := row.AddCell()
					cellPara := cell.AddParagraph()
					cellRun := cellPara.AddRun()
					cellRun.AddText(fmt.Sprintf("单元格 %d-%d", i+1, j+1))
				}
			}

			//分隔不同表格
			//doc.InsertParagraphAfter(para).AddRun().AddText("--------------")
			doc.InsertParagraphAfter(para).AddRun()

			// 创建并配置表格
			table = doc.InsertTableAfter(para)      // 在标签段落之后插入表格
			table.Properties().SetWidthPercent(100) // 设置表格宽度为100%
			borders = table.Properties().Borders()
			borders.SetAll(wml.ST_BorderSingle, color.Black, measurement.Dxa) // 设置所有边框为单线黑色

			for i := 0; i < 3; i++ { // 创建表格行和单元格
				row := table.AddRow()
				for j := 0; j < 3; j++ {
					cell := row.AddCell()
					cellPara := cell.AddParagraph()
					cellRun := cellPara.AddRun()
					cellRun.AddText(fmt.Sprintf("单元格 %d-%d", i+1, j+1))
				}
			}

			// 移除标签段落
			replaceParagraphWithTable(&para, tag) // 替换标签段落为表格
			// 删除段落
			doc.RemoveParagraph(para) // 从文档中删除标签段落
			return nil
		}
	}
	return fmt.Errorf("未找到标签 %s", tag) // 如果未找到标签段落，返回错误
}

// 在指定标签处插入图表
func insertImageAt(doc *document.Document, imageBuffer *bytes.Buffer, tag string) error {
	paras := doc.Paragraphs() // 获取文档中的所有段落
	for _, para := range paras {
		if paraContainsTag(&para, tag) { // 如果段落包含指定标签
			img, err := common.ImageFromBytes(imageBuffer.Bytes()) // 从内存缓冲区中加载图片
			if err != nil {
				return fmt.Errorf("无法从内存缓冲区中加载图片: %v", err)
			}

			// 创建图片引用
			iref, err := doc.AddImage(img) // 将图片添加到文档中
			if err != nil {
				return fmt.Errorf("无法将图片添加到文档: %v", err)
			}

			// 创建新的段落和运行以插入图表
			newPara := doc.InsertParagraphAfter(para) // 在标签段落之后插入新段落
			run := newPara.AddRun()

			// 插入图片到文档
			imgInl, err := run.AddDrawingInline(iref) // 在运行中添加图片
			if err != nil {
				return fmt.Errorf("插入图片时出错: %v", err)
			}
			imgInl.SetSize(6*measurement.Inch, 4*measurement.Inch) // 设置图片尺寸为6x4英寸

			// 移除标签段落
			replaceParagraphWithTable(&para, tag) // 替换标签段落为图表
			// 删除段落
			doc.RemoveParagraph(para) // 从文档中删除标签段落
			return nil
		}
	}
	return fmt.Errorf("未找到标签 %s", tag) // 如果未找到标签段落，返回错误
}

// 判断段落是否包含指定标签
func paraContainsTag(para *document.Paragraph, tag string) bool {
	for _, run := range para.Runs() { // 遍历段落中的每个运行
		if strings.Contains(run.Text(), tag) { // 如果运行文本包含标签
			return true
		}
	}
	return false
}

// 移除标签段落
func replaceParagraphWithTable(para *document.Paragraph, tag string) {
	// 找到标签的 Run
	for _, run := range para.Runs() {
		log.Printf("替换标签：tag =  %v", tag)
		if strings.Contains(run.Text(), tag) {
			para.InsertRunAfter(para.AddRun())
			run.Clear()         // 清除原有内容
			para.RemoveRun(run) // 移除运行
			break
		}
	}
}

// 删除两个标签之间的段落
func removeParagraphsBetweenTags(doc *document.Document, startTag, endTag string) error {
	paras := doc.Paragraphs()
	startIndex, endIndex := -1, -1

	// 找到包含startTag和endTag的段落索引
	for i, para := range paras {
		if paraContainsTag(&para, startTag) {
			startIndex = i
		}
		if paraContainsTag(&para, endTag) {
			endIndex = i
			break
		}
	}

	if startIndex == -1 {
		return fmt.Errorf("未找到标签 %s", startTag)
	}
	if endIndex == -1 {
		return fmt.Errorf("未找到标签 %s", endTag)
	}
	if startIndex >= endIndex {
		return fmt.Errorf("标签 %s 和 %s 之间的顺序不正确", startTag, endTag)
	}

	// 删除startTag和endTag之间的段落
	for i := startIndex; i <= endIndex; i++ {
		doc.RemoveParagraph(paras[i])
	}

	return nil
}

// 删除指定标签段落
func removeParagraphWithTag(doc *document.Document, tag string) error {
	paras := doc.Paragraphs()
	for _, para := range paras {
		if paraContainsTag(&para, tag) {
			doc.RemoveParagraph(para)
			return nil
		}
	}
	return fmt.Errorf("未找到标签 %s", tag)
}

// 创建折线图并将其存储到缓存中
func createLineChart() (*bytes.Buffer, error) {
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
		return nil, nil
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

	// 将图表保存到缓冲区
	var buf bytes.Buffer
	if err := dc.EncodePNG(&buf); err != nil {
		return nil, err
	}
	return &buf, nil
}

// 创建柱状图并将其存储到缓存中
func createBarChart() (*bytes.Buffer, error) {
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
		return nil, nil
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

	// 将图表保存到缓冲区
	var buf bytes.Buffer
	if err := dc.EncodePNG(&buf); err != nil {
		return nil, err
	}
	return &buf, nil
}

// 创建饼图并将其存储到缓存中
func createPieChart() (*bytes.Buffer, error) {
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
		return nil, nil
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

	// 将图表保存到缓冲区
	var buf bytes.Buffer
	if err := dc.EncodePNG(&buf); err != nil {
		return nil, err
	}
	return &buf, nil
}
