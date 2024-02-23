package types

const (
	//DsmsConfig string = "dsms_config"

	//clickhouse数据表
	HttpCaptureTableName       string = "http_capture"
	HttpCaptureFilterTableName string = "http_capture_filter"

	//请求内容类型
	JsonType   string = "application/json"
	TextType   string = "text/plain;charset=UTF-8"
	FormType   string = "application/x-www-form-urlencoded"
	StreamType string = "application/octet-stream"
)
