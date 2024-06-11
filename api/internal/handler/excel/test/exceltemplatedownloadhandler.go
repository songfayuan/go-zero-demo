package test

import (
	"net/http"

	"github.com/zeromicro/go-zero/rest/httpx"
	"go-zero-demo/api/internal/logic/excel/test"
	"go-zero-demo/api/internal/svc"
)

func ExcelTemplateDownloadHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		l := test.NewExcelTemplateDownloadLogic(r.Context(), svcCtx, w)
		err := l.ExcelTemplateDownload()
		if err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
		} else {
			httpx.Ok(w)
		}
	}
}
