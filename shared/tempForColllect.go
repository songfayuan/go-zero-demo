package shared

// 级联管理-采集协议-临时 map

const (
	CacheKeyForFormCollect = "cache:form:collect"
)

type A2 int
type A5 int

type A3 int
type A6 int

const (
	A2Default A2 = iota
	A2Http
	A2POP3
	A2FTP
)

const (
	A3_0 A3 = iota
	A3_1
	A3_2
	A3_3
	A3_4
	A3_5
	A3_6
)

const (
	A5Default A5 = iota
	A5SqlServer
	A5MySql
	A5Oracle
)

const (
	A6_0 A6 = iota
	A6_1
	A6_2
)

func (s A2) String() string {
	switch s {
	case A2Http:
		return "HTTP"
	case A2POP3:
		return "POP3"
	case A2FTP:
		return "FTP"
	default:
		return "-"
	}
}

func (s A2) StringToC() string {
	switch s {
	case A2Http:
		return "HTTP"
	case A2POP3:
		return "POP3"
	case A2FTP:
		return "FTP"
	default:
		return ""
	}
}

func (s A3) String() string {
	switch s {
	case A3_1:
		return "请求信息"
	case A3_2:
		return "响应信息"
	case A3_3:
		return "文件信息"
	case A3_4:
		return "请求头信息"
	case A3_5:
		return "响应头信息"
	case A3_6:
		return "URL"
	default:
		return "-"
	}
}

func (s A3) StringToC() string {
	switch s {
	case A3_1:
		return "request_body"
	case A3_2:
		return "response_body"
	case A3_3:
		return "file_info"
	case A3_4:
		return "request_header"
	case A3_5:
		return "response_header"
	case A3_6:
		return "url"
	default:
		return ""
	}
}

func (s A5) String() string {
	switch s {
	case A5SqlServer:
		return "SqlServer"
	case A5MySql:
		return "MySql"
	case A5Oracle:
		return "Oracle"
	default:
		return "-"
	}
}

func (s A5) StringToC() string {
	switch s {
	case A5SqlServer:
		return "SqlServer"
	case A5MySql:
		return "MySql"
	case A5Oracle:
		return "Oracle"
	default:
		return ""
	}
}

func (s A6) String() string {
	switch s {
	case A6_1:
		return "请求信息"
	case A6_2:
		return "响应信息"
	default:
		return "-"
	}
}

func (s A6) StringToC() string {
	switch s {
	case A6_1:
		return "request_body"
	case A6_2:
		return "response_body"
	default:
		return ""
	}
}

type TempForCollect struct {
	A1 int   `json:"a1"`
	A2 []int `json:"a2"`
	A3 []int `json:"a3"`
	A4 int   `json:"a4"`
	A5 []int `json:"a5"`
	A6 []int `json:"a6"`
}

func (t *TempForCollect) AppResult() (int, []string, []string) {
	if t.A1 == 1 {
		var str1 []string
		for _, v := range t.A2 {
			if vv := A2(v).StringToC(); vv != "-" {
				str1 = append(str1, vv)
			}
		}

		var str2 []string
		for _, v := range t.A3 {
			if vv := A3(v).StringToC(); vv != "-" {
				str2 = append(str2, vv)
			}
		}
		return 1, str1, str2
	}
	return 0, nil, nil
}

func (t *TempForCollect) DBResult() (int, []string, []string) {
	if t.A4 == 1 {
		var str1 []string
		for _, v := range t.A5 {
			if vv := A5(v).StringToC(); vv != "-" {
				str1 = append(str1, vv)
			}
		}

		var str2 []string
		for _, v := range t.A6 {
			if vv := A6(v).StringToC(); vv != "-" {
				str2 = append(str2, vv)
			}
		}
		return 1, str1, str2
	}
	return 0, nil, nil
}
