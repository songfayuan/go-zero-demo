package main

import (
	"bytes"
	"fmt"
	"image/color"
	"os"

	"github.com/fogleman/gg"
)

type Data struct {
	Label string
	Value float64
}

func main() {
	data := []Data{
		{"192.168.129.13", 60.1},
		{"192.168.129.73", 59.3},
		{"192.168.129.24", 38.2},
		{"192.168.129.19", 22.7},
		{"192.168.129.16", 16.7},
	}

	colors := []color.RGBA{
		{255, 99, 71, 255},   // 红色
		{255, 165, 0, 255},   // 橙色
		{30, 144, 255, 255},  // 蓝色
		{169, 169, 169, 255}, // 灰色
		{169, 169, 169, 255}, // 灰色
	}

	const (
		width     = 600
		height    = 200
		barHeight = 20
		padding   = 10
	)

	dc := gg.NewContext(width, height)
	dc.SetColor(color.White)
	dc.Clear()

	// Calculate total height of the bars
	totalBarHeight := float64(len(data))*(barHeight+padding) - padding

	// Calculate starting Y coordinate to center the bars vertically
	startY := (height - totalBarHeight) / 2

	// Draw bars and text
	for i, d := range data {
		y := startY + float64(i)*(barHeight+padding)

		// Draw circle with number
		dc.SetColor(colors[i])
		dc.DrawCircle(20, y+barHeight/2, 10)
		dc.Fill()
		dc.SetColor(color.White)
		dc.DrawStringAnchored(fmt.Sprintf("%d", i+1), 20, y+barHeight/2, 0.5, 0.5)

		// Draw label
		dc.SetColor(color.Black)
		dc.DrawStringAnchored(d.Label, 50, y+barHeight/2, 0, 0.5)

		// Draw bar
		barWidth := d.Value * 5 // Scale value to fit the width
		dc.SetColor(colors[i])
		dc.DrawRectangle(150, y, barWidth, barHeight)
		dc.Fill()

		// Draw value
		dc.SetColor(color.Black)
		dc.DrawStringAnchored(fmt.Sprintf("%.1f GB", d.Value), 150+barWidth+10, y+barHeight/2, 0, 0.5)
	}

	// Save to buffer
	var buf bytes.Buffer
	dc.EncodePNG(&buf)

	// Optionally save to file
	if err := os.WriteFile("bar_chart.png", buf.Bytes(), 0644); err != nil {
		fmt.Println("Error saving file:", err)
	}
}
