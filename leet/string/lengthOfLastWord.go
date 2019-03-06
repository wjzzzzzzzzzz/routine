package string

import (
	"strings"
)

//给定一个仅包含大小写字母和空格 ' ' 的字符串，返回其最后一个单词的长度。
//
//如果不存在最后一个单词，请返回 0 。
//先把最后面的空格替换 ,然后再找前面的空格
//ok
func lengthOfLastWord(s string) int {
	if s == "" {
		return 0
	}

	for {
		index := strings.LastIndex(s, " ")
		if index == len(s)-1 {
			s = s[:len(s)-1]
			if len(s) == 0 {
				return 0
			}
		} else {
			return len(s) - index - 1
		}
	}
}
