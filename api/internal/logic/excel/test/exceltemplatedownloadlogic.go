package test

import (
	"context"
	"go-zero-demo/common/errors/errorx"
	"net/http"
	"os"

	"github.com/zeromicro/go-zero/core/logx"
	"go-zero-demo/api/internal/svc"
)

type ExcelTemplateDownloadLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
	writer http.ResponseWriter
}

func NewExcelTemplateDownloadLogic(ctx context.Context, svcCtx *svc.ServiceContext, writer http.ResponseWriter) *ExcelTemplateDownloadLogic {
	return &ExcelTemplateDownloadLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
		writer: writer,
	}
}

func (l *ExcelTemplateDownloadLogic) ExcelTemplateDownload() (err error) {
	SavePath := l.svcCtx.Config.UploadFile.TemplatePath
	filePath := "demo_excel_template.xlsx"

	fullPath := SavePath + filePath
	fileName := "Excel导入模板.xlsx"

	//fullPath = "/Users/songfayuan/GolandProjects/go-zero-demo/template/excel/demo_excel_template.xlsx"  //测试地址，绝对路径
	_, err = os.Stat(fullPath)
	if err != nil || os.IsNotExist(err) {
		return errorx.New("文件不存在")
	}
	bytes, err := os.ReadFile(fullPath)
	if err != nil {
		return errorx.New("读取文件失败")
	}

	l.writer.Header().Add("Content-Type", "application/octet-stream")
	l.writer.Header().Add("Content-Disposition", "attachment; filename= "+fileName)
	l.writer.Write(bytes)

	return
}
