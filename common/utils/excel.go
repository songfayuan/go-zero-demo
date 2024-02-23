package utils

import (
	"errors"
	"github.com/xuri/excelize/v2"
	"github.com/zeromicro/go-zero/core/mapping"
	"net/http"
	"reflect"
	"strconv"
	"strings"
)

type ExcelCallback func(options map[string]interface{}) error

type ExcelOption struct {
	Sheet    string // excel页
	StartRow int    // 开始行：默认1
}

var (
	localExcelOption = ExcelOption{
		Sheet:    "Sheet1",
		StartRow: 1,
	}
)

func __demo(r *http.Request) {
	// 保存结构
	type saveT struct {
		ColA int    `json:"ColA,optional" excel:"col=1"` // 第1列：a
		ColB string `json:"ColB,optional" excel:"col=2"` // 第2列：b
		ColC string `json:"ColC,optional" excel:"col=5"` // 第5列：c
	}

	// 文件源
	file, _, _ := r.FormFile("file")
	// 打开文件
	f, err := excelize.OpenReader(file)
	if err != nil {
		return
	}

	// 保存逻辑
	all := make([]saveT, 0)
	cbHandler := func(data map[string]interface{}) error {
		temp := saveT{}
		err := mapping.UnmarshalJsonMap(data, &temp)
		if err != nil {
			return err
		}
		all = append(all, temp)
		return nil
	}

	// 解析配置（可选）
	parseConfig := ExcelOption{Sheet: "Sheet1", StartRow: 1}
	// 解析逻辑
	if err := ParseExcel(f, saveT{}, cbHandler, parseConfig); err != nil {
		// error
	}
}

// ParseExcel 解析excel
// 使用参考上方 __demo
func ParseExcel(f *excelize.File, save interface{}, cbHandler ExcelCallback, opts ...ExcelOption) (err error) {
	if reflect.ValueOf(save).Kind() != reflect.Struct {
		return errors.New("保存位置必须传入结构体实体（非指针）")
	}
	if reflect.TypeOf(save).Kind() != reflect.Struct {
		return errors.New("保存位置必须传入结构体")
	}
	if cbHandler == nil {
		return errors.New("回调函数无效")
	}

	// 解析参数配置
	localExcelOption.parse(opts...)

	// 计算结excel的列数对应的构体字段
	fieldsMap := fieldsOrderMap(save)

	// 获取所有行
	rows, err := f.GetRows(localExcelOption.Sheet)
	if err != nil {
		return
	}

	for i, row := range rows {
		rowLen := len(row)
		// 从第 {localExcelOption.StartRow} 行开始
		if rowLen == 0 || i < localExcelOption.StartRow-1 {
			continue
		}

		rowValue := map[string]interface{}{}
		for colId, colVal := range row {
			if v, ok := fieldsMap[colId+1]; ok {
				rowValue[v] = colVal
			}
		}

		// 回调处理值
		err = cbHandler(rowValue)
		if err != nil {
			return
		}
	}

	return
}

// 覆盖本地解析参数
func (local *ExcelOption) parse(opts ...ExcelOption) {
	for _, opt := range opts {
		if opt.Sheet != "" {
			local.Sheet = opt.Sheet
		}
		if opt.StartRow >= 1 {
			local.StartRow = opt.StartRow
		}
	}
}

// 获取tag值
// 例如：tagValue(`,col=1,default=2`, `col`) => 1
func tagValue(excelTag string, key string) (val string) {
	arr := strings.Split(excelTag, ",")
	for _, i := range arr {
		index := strings.Index(i, key)
		if index > -1 {
			keyLen := len(key)
			// 包含,且长度大于
			if len(i) > keyLen+1 {
				return i[index+keyLen+1:]
			}
			return
		}
	}

	return
}

// 计算结excel的列数对应的构体字段
// map [int=>列数] string=>结构体json字段
func fieldsOrderMap(v interface{}) map[int]string {
	fieldsMap := make(map[int]string, 0)

	// 映射字段
	t := reflect.TypeOf(v)
	for i := 0; i < t.NumField(); i++ {
		// tag
		field := t.Field(i)
		tag := field.Tag

		// json tag
		jsonTag := tag.Get("json")
		jsonTagName := jsonTag[:strings.Index(jsonTag, ",")]

		// excel tag
		excelTag := tag.Get("excel")
		excelTagCol := tagValue(excelTag, "col")
		// 计算列
		var excelTagColNum int
		if excelTagCol != "" {
			excelTagColNum, _ = strconv.Atoi(excelTagCol)
		}
		if excelTagColNum <= 0 {
			excelTagColNum = i + 1 // 取字段的顺序
		}

		if jsonTagName != "" && excelTagColNum > 0 {
			fieldsMap[excelTagColNum] = jsonTagName
		}
	}

	return fieldsMap
}
