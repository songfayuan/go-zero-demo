package types

type OptCompare string

type OptSize string

type OptLogical string

type OptBracket string

const (
	OptGt      OptCompare = "gt"      // 大于
	OptLt      OptCompare = "lt"      // 小于
	OptBetween OptCompare = "between" // 在两者之间

	OptSizeKb OptSize = "KB"
	OptSizeMb OptSize = "MB"
	OptSizeGb OptSize = "GB"

	OptAnd OptLogical = "and"
	OptOr  OptLogical = "or"

	OptLeftBracket1  OptBracket = "1"
	OptLeftBracket2  OptBracket = "2"
	OptLeftBracket3  OptBracket = "3"
	OptRightBracket1 OptBracket = "-1"
	OptRightBracket2 OptBracket = "-2"
	OptRightBracket3 OptBracket = "-3"
)

func (v OptCompare) String() (s string) {
	switch v {
	case OptGt:
		s = "大于"
	case OptLt:
		s = "小于"
	case OptBetween:
		s = "在两者之间"
	}
	return
}

func (v OptSize) String() (s string) {
	return string(v)
}

func (v OptLogical) String() (s string) {
	switch v {
	case OptAnd:
		s = "和"
	case OptOr:
		s = "且"
	}
	return
}

func (v OptBracket) String() (s string) {
	switch v {
	case OptLeftBracket1:
		s = "{"
	case OptLeftBracket2:
		s = "{{"
	case OptLeftBracket3:
		s = "{{{"
	case OptRightBracket1:
		s = "}"
	case OptRightBracket2:
		s = "}}"
	case OptRightBracket3:
		s = "}}}"
	}
	return
}
