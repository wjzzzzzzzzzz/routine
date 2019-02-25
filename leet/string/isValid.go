package string

import "strings"

//给定一个只包括 '('，')'，'{'，'}'，'['，']' 的字符串，判断字符串是否有效。
//有效字符串需满足：
//左括号必须用相同类型的右括号闭合。
//左括号必须以正确的顺序闭合。
//注意空字符串可被认为是有效字符串。
//

func isValid(s string) bool {
	if s == "" {
		return true
	}

	strings.Index(s,"(")

	split := strings.Split(s, "")

}
