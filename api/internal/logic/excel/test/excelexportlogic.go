package test

import (
	"context"
	"fmt"
	"github.com/xuri/excelize/v2"
	"net/http"

	"go-zero-demo/api/internal/svc"
	"go-zero-demo/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type ExcelExportLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
	writer http.ResponseWriter
}

type cellValue struct {
	sheet string
	cell  string
	value string
}

func NewExcelExportLogic(ctx context.Context, svcCtx *svc.ServiceContext, writer http.ResponseWriter) *ExcelExportLogic {
	return &ExcelExportLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
		writer: writer,
	}
}

func (l *ExcelExportLogic) ExcelExport(req *types.ExcelExportlReq) (resp *types.DefaultResponse, err error) {
	//这里仅演示Excel导出逻辑，真实数据自己增加对应的查询逻辑。
	excelFile := excelize.NewFile()
	//insert title
	cellValues := make([]*cellValue, 0)
	cellValues = append(cellValues, &cellValue{
		sheet: "sheet1",
		cell:  "A1",
		value: "序号",
	}, &cellValue{
		sheet: "sheet1",
		cell:  "B1",
		value: "IP地址",
	}, &cellValue{
		sheet: "sheet1",
		cell:  "C1",
		value: "账号",
	}, &cellValue{
		sheet: "sheet1",
		cell:  "D1",
		value: "姓名",
	}, &cellValue{
		sheet: "sheet1",
		cell:  "E1",
		value: "最近访问时间",
	}, &cellValue{
		sheet: "sheet1",
		cell:  "F1",
		value: "设备状态",
	}, &cellValue{
		sheet: "sheet1",
		cell:  "G1",
		value: "访问分级",
	})
	// 创建一个工作表
	index, _ := excelFile.NewSheet("Sheet1")
	// 设置工作簿的默认工作表
	excelFile.SetActiveSheet(index)
	//插入表格头
	for _, cellValue := range cellValues {
		excelFile.SetCellValue(cellValue.sheet, cellValue.cell, cellValue.value)
	}
	//设置表格头字体样式
	styleId, err := excelFile.NewStyle(&excelize.Style{
		Font: &excelize.Font{
			Bold:   true,  //黑体
			Italic: false, //倾斜
			Family: "宋体",
			Size:   14,
			//Color:  "微软雅黑",
		},
	})
	if err != nil {
		fmt.Println(err)
	}
	for _, data := range cellValues {
		excelFile.SetCellStyle(data.sheet, data.cell, data.cell, styleId)
	}
	excelFile.SetColWidth("sheet1", "B", "G", 20)

	cnt := 1
	for i := 0; i <= 6; i++ {
		cnt = cnt + 1
		for k1, v1 := range cellValues {
			switch k1 {
			case 0:
				v1.cell = fmt.Sprintf("A%d", cnt)
				v1.value = fmt.Sprintf("%d", i+1)
			case 1:
				v1.cell = fmt.Sprintf("B%d", cnt)
				v1.value = "1"
			case 2:
				v1.cell = fmt.Sprintf("C%d", cnt)
				v1.value = "2"
			case 3:
				v1.cell = fmt.Sprintf("D%d", cnt)
				v1.value = "3"
			case 4:
				v1.cell = fmt.Sprintf("E%d", cnt)
				v1.value = "4"
			case 5:
				v1.cell = fmt.Sprintf("F%d", cnt)
				v1.value = "5"
			case 6:
				v1.cell = fmt.Sprintf("G%d", cnt)
				v1.value = "6"
			}
		}
		for _, vc := range cellValues {
			excelFile.SetCellValue(vc.sheet, vc.cell, vc.value)
		}
	}
	fileName := "ABCD.xlsx"

	//如果是下载，则需要在Header中设置这两个参数
	l.writer.Header().Add("Content-Type", "application/octet-stream")
	l.writer.Header().Add("Content-Disposition", "attachment; filename= "+fileName)
	//l.writer.Header().Add("Content-Transfer-Encoding", "binary")
	excelFile.Write(l.writer)
	return
}
