package test

import (
	"go-zero-demo/common/utils"
	"net/http"

	"github.com/zeromicro/go-zero/rest/httpx"
	"go-zero-demo/api/internal/logic/excel/test"
	"go-zero-demo/api/internal/svc"
	"go-zero-demo/api/internal/types"
)

func ExcelImportHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.ExcelUploadReq
		req.DeptId = r.FormValue("deptId")
		f, fh, e := utils.ParseFile(r, "file")
		if e != nil {
			httpx.Error(w, e)
			return
		}
		req.File = &types.File{File: f, FileHeader: fh}

		l := test.NewExcelImportLogic(r.Context(), svcCtx)
		resp, err := l.ExcelImport(&req)
		if err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
		} else {
			httpx.OkJson(w, resp)
		}
	}
}
