package utils

import "strings"

// EndBracket 获取后括号
func EndBracket(start string) string {
	switch start {
	case "(":
		start = ")"
	case "[":
		start = "]"
	case "{":
		start = "}"
	}

	return ""
}

// EnTransBrackets ( => { ,  [ => {{,  { => {{{
func EnTransBrackets(s string) string {
	switch s {
	case "(":
		s = "{"
	case "[":
		s = "{{"
	case "{":
		s = "{{{"
	case ")":
		s = "}"
	case "]":
		s = "}}"
	case "}":
		s = "}}}"
	}

	return s
}

// DeTransBrackets { => (,  {{ => [,  {{{ => {
func DeTransBrackets(s string) string {
	switch s {
	case "{":
		s = "("
	case "{{":
		s = "["
	case "{{{":
		s = "("
	case "}":
		s = ")"
	case "}}":
		s = "]"
	case "}}}":
		s = "}"
	}

	return s
}

// BracketsIsMatch 括号是否匹配
func BracketsIsMatch(all []string) bool {
	if len(all) == 0 {
		return true
	}

	s := strings.Join(all, "")

	m := map[string]string{")": "(", "]": "[", "}": "{"}
	var stack []string
	// 把字符串的每个字符放进栈中，每放一个就判断与前一个是不是配对的
	for i := 0; i < len(s); i++ {
		if len(stack) == 0 {
			stack = append(stack, string(s[i]))
		} else {
			// 判断是否配对
			// 如果是相同的话，那就去除栈的最后一个元素
			// 如果不相同的话，那就把源字符串的对应元素加进栈中
			if stack[len(stack)-1] == m[string(s[i])] {
				stack = stack[:len(stack)-1]
			} else {
				stack = append(stack, string(s[i]))
			}
		}
	}

	if len(stack) == 0 {
		return true
	} else {
		return false
	}
}
