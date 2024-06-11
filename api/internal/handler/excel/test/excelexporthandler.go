package test

import (
	"net/http"

	"github.com/zeromicro/go-zero/rest/httpx"
	"go-zero-demo/api/internal/logic/excel/test"
	"go-zero-demo/api/internal/svc"
	"go-zero-demo/api/internal/types"
)

func ExcelExportHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.ExcelExportlReq
		if err := httpx.Parse(r, &req); err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
			return
		}

		l := test.NewExcelExportLogic(r.Context(), svcCtx, w)
		resp, err := l.ExcelExport(&req)
		if err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
		} else {
			httpx.OkJsonCtx(r.Context(), w, resp)
		}
	}
}
