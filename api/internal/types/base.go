package types

import "mime/multipart"

type File struct {
	File       multipart.File
	FileHeader *multipart.FileHeader
}

type ExcelUploadReq struct {
	DeptId string `json:"deptId"` // 部门id
	File   *File  `json:"file"`   // excel文件
}
