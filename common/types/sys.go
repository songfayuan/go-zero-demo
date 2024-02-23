package types

// SysConfigTag 系统设置：标识符
type SysConfigTag string

// PwdSymbol 密码符号
type PwdSymbol string

const (
	TagForSafety     SysConfigTag = "safety"             // 安全设置
	TagForMail       SysConfigTag = "mail"               // 发件箱设置
	TagForMemberSync SysConfigTag = "syncMemberAdSource" // 终端-人员：同步AD源

	PwdSymbolForDigit     PwdSymbol = "digit"     // 数字
	PwdSymbolForLowercase PwdSymbol = "lowercase" // 小写字母
	PwdSymbolForUppercase PwdSymbol = "uppercase" // 大写字母
	PwdSymbolForSpecial   PwdSymbol = "special"   // 特殊符号
)

func (v SysConfigTag) String() (s string) {
	switch v {
	case TagForSafety:
		s = "安全设置"
	case TagForMail:
		s = "发件箱设置"
	case TagForMemberSync:
		s = "人员管理-同步AD源"
	}
	return
}

func (v PwdSymbol) String() (s string) {
	switch v {
	case PwdSymbolForDigit:
		s = "数字"
	case PwdSymbolForLowercase:
		s = "小写字母"
	case PwdSymbolForUppercase:
		s = "大写字母"
	case PwdSymbolForSpecial:
		s = "特殊符号"
	}
	return
}
