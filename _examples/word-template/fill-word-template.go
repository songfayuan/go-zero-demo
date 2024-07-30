package main

import (
	"fmt"
	"log"
	"strings"

	"github.com/Esword618/unioffice/color"
	"github.com/Esword618/unioffice/common"
	"github.com/Esword618/unioffice/document"
	"github.com/Esword618/unioffice/measurement"
	"github.com/Esword618/unioffice/schema/soo/wml"
	"github.com/fogleman/gg"
)

//教程：https://blog.csdn.net/u011019141/article/details/140788882

func main() {
	// 定义文档路径和图表文件路径
	docPath := "/Users/songfayuan/Downloads/template.docx"
	chartFile := "/Users/songfayuan/Downloads/123456.PNG"
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

	// 创建折线图并保存为图片
	if err := createLineChart(chartFile); err != nil {
		log.Fatalf("创建图表时出错: %v", err)
	}

	// 在指定标签处插入图表
	if err := insertImageAt(doc, chartFile, "{{tubiao}}"); err != nil {
		log.Fatalf("插入图表时出错: %v", err)
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
func insertImageAt(doc *document.Document, imagePath string, tag string) error {
	paras := doc.Paragraphs() // 获取文档中的所有段落
	for _, para := range paras {
		if paraContainsTag(&para, tag) { // 如果段落包含指定标签
			img, err := common.ImageFromFile(imagePath) // 从文件中加载图片
			if err != nil {
				return fmt.Errorf("无法从文件中加载图片: %v", err)
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

// 创建折线图并保存为图片
func createLineChart(filename string) error {
	const (
		width  = 640
		height = 480
	)

	dc := gg.NewContext(width, height)
	dc.SetRGB(1, 1, 1) // 背景色为白色
	dc.Clear()

	dc.SetRGB(0, 0, 1) // 线条颜色为蓝色
	dc.SetLineWidth(2)

	data := []struct {
		x, y float64
	}{
		{1, 5}, {2, 7}, {3, 6}, {4, 8}, {5, 9}, // 折线图数据点
	}

	if len(data) > 0 {
		dc.MoveTo(data[0].x*100, height-data[0].y*40) // 移动到第一个数据点
		for _, pt := range data[1:] {
			dc.LineTo(pt.x*100, height-pt.y*40) // 绘制线条到下一个数据点
		}
		dc.Stroke() // 结束绘制
	}

	return dc.SavePNG(filename) // 保存图像为PNG文件
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
