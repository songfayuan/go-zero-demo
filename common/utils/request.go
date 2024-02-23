package utils

import (
	"go-zero-demo/common/errors/errorx"
	"mime/multipart"
	"net/http"
)

// ParseFile 处理单个文件
func ParseFile(r *http.Request, field string) (file multipart.File, fileHeader *multipart.FileHeader, err error) {
	mf := r.MultipartForm.File
	if mfVal, ok := mf[field]; !ok {
		err = errorx.New("未选择文件")
		return

	} else if len(mfVal) > 1 {
		err = errorx.New("只允许上传单个文件")
		return
	}

	file, fileHeader, err = r.FormFile(field)
	return
}

// ParseFiles 处理多个文件 todo
func ParseFiles(r *http.Request, field string) {

}
