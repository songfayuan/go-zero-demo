package test

import (
	"context"
	"fmt"
	"github.com/xuri/excelize/v2"
	"github.com/zeromicro/go-zero/core/mapping"
	"go-zero-demo/common/errors/errorx"
	"go-zero-demo/common/utils"
	"path/filepath"
	"strings"

	"go-zero-demo/api/internal/svc"
	"go-zero-demo/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type ExcelImportLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

type excelDataForDept struct {
	DeptId       string `json:"DeptId,optional" excel:"col=1"`       // 第1列：部门id
	ParentDeptId string `json:"ParentDeptId,optional" excel:"col=2"` // 第2列：上级部门id
	DeptName     string `json:"DeptName,optional" excel:"col=3"`     // 第3列：部门名称
	Level        string `json:"Level,optional" excel:"col=4"`        // 第4列：部门等级（分级名称）
}

type excelDataForMember struct {
	DeptId  string `json:"DeptId,optional" excel:"col=1"`  // 第1列：部门
	Name    string `json:"Name,optional" excel:"col=2"`    // 第2列：姓名
	Account string `json:"Account,optional" excel:"col=3"` // 第3列：帐号
	Level   string `json:"Level,optional" excel:"col=4"`   // 第4列：等级（分级名称）
	IpAddr  string `json:"IpAddr,optional" excel:"col=5"`  // 第5列：IP
	MacAddr string `json:"MacAddr,optional" excel:"col=6"` // 第6列：MAC
}

var (
	validUploadFileExt = map[string]any{
		".xlsx": nil,
		".xls":  nil,
	}
)

func NewExcelImportLogic(ctx context.Context, svcCtx *svc.ServiceContext) *ExcelImportLogic {
	return &ExcelImportLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *ExcelImportLogic) ExcelImport(req *types.ExcelUploadReq) (resp *types.ExcelImportResp, err error) {

	if _, ok := validUploadFileExt[strings.ToLower(filepath.Ext(req.File.FileHeader.Filename))]; !ok {
		return nil, errorx.New("无效的文件格式")
	}
	// 打开文件
	f, err := excelize.OpenReader(req.File.File)
	if err != nil {
		return nil, errorx.New("无效的文件")
	}

	/* 解析部门Sheet数据 start */
	// 解析文件参数
	var excelDept []excelDataForDept
	if excelDept, err = parseFileDept(f); err != nil {
		return
	}
	// format
	for _, i := range excelDept {
		fmt.Printf("Excel数据：%v/%v/%v/%v", i.DeptId, i.ParentDeptId, i.DeptName, i.Level)
	}
	/* 解析部门Sheet数据 end */

	/* 解析用户Sheet数据 start */
	// 解析文件参数
	var excelMember []excelDataForMember
	if excelMember, err = parseFileUser(f); err != nil {
		return
	}
	// format
	for _, i := range excelMember {
		fmt.Printf("Excel数据：%v/%v/%v/%v/%v/%v", i.DeptId, i.Name, i.Account, i.Level, i.IpAddr, i.MacAddr)
	}
	/* 解析用户Sheet数据 end */

	return &types.ExcelImportResp{
		Code:    200,
		Message: "导入成功",
		Data: types.ExcelImportData{
			Total:   10,
			Success: 10,
			Msg:     "成功",
		},
	}, nil
}

// 解析部门Sheet数据
func parseFileDept(f *excelize.File) ([]excelDataForDept, error) {

	// 解析参数（可选）
	excelOption := utils.ExcelOption{Sheet: "部门", StartRow: 2}

	// 映射回调
	all := make([]excelDataForDept, 0)
	cbHandler := func(data map[string]interface{}) error {
		temp := excelDataForDept{}
		err := mapping.UnmarshalJsonMap(data, &temp)
		if err != nil {
			return err
		}
		all = append(all, temp)
		return nil
	}

	// 映射
	if err := utils.ParseExcel(f, excelDataForDept{}, cbHandler, excelOption); err != nil {
		return nil, errorx.New("解析文件时出错：" + err.Error())
	}

	if len(all) == 0 {
		return nil, errorx.New("文件中无有效数据")
	}

	return all, nil
}

// 解析用户Sheet数据
func parseFileUser(f *excelize.File) ([]excelDataForMember, error) {
	// 解析参数（可选）
	excelOption := utils.ExcelOption{Sheet: "用户", StartRow: 2}

	// 映射回调
	all := make([]excelDataForMember, 0)
	cbHandler := func(data map[string]interface{}) error {
		temp := excelDataForMember{}
		err := mapping.UnmarshalJsonMap(data, &temp)
		if err != nil {
			return err
		}
		all = append(all, temp)
		return nil
	}

	// 映射
	if err := utils.ParseExcel(f, excelDataForMember{}, cbHandler, excelOption); err != nil {
		return nil, errorx.New("解析文件时出错：" + err.Error())
	}

	if len(all) == 0 {
		return nil, errorx.New("文件中无有效数据")
	}

	return all, nil
}
