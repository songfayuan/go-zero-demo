package shared

/* row 标记 */

// DelNo 正常使用
var DelNo int64 = 0

// DelYes 软删除
var DelYes int64 = 1

/* 时间格式化模版 */

// DateTimeFormatTplStandardDateTime 模板1  Y-m-d H:i:s
var DateTimeFormatTplStandardDateTime = "2006-01-02 15:04:05"

// DateTimeFormatTplStandardDate 模板3  Y-m-d
var DateTimeFormatTplStandardDate = "2006-01-02"

// DateTimeFormatTplStandardTime 模板3  H:i:s
var DateTimeFormatTplStandardTime = "15:04:05"

const (
	CodeOk int64  = 200 //ok
	MsgOk  string = "success"
)
