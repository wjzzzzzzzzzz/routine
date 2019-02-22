package main

import (
	"log"
	"strings"
)

//编写一个函数来查找字符串数组中的最长公共前缀。
//如果不存在公共前缀，返回空字符串 ""。
func longestCommonPrefix(strs []string) string {
	if len(strs) == 0 {
		return ""
	}
	var temp, result string
	one := strings.Split(strs[0], "")
	for j, _ := range one {
		temp = temp + one[j]
		for _, v := range strs {
			if !strings.HasPrefix(v, temp) {
				return result
			}
		}
		result = temp
	}
	return result
}
func main() {
	log.SetFlags(log.Lshortfile | log.LstdFlags)
	test := []string{
		"ca", "caaa", "ca",
	}
	prefix := longestCommonPrefix(test)
	log.Println(prefix)

}
