package main

import (
	"strconv"
)

//判断一个整数是否是回文数。回文数是指正序（从左向右）和倒序（从右向左）读都是一样的整数。
func isPalindrome(x int) bool {
	if x < 0 {
		return false
	} else {
		s := []byte(strconv.Itoa(x))
		length := len(s) - 1
		for i, v := range s {
			if s[length-i] != v {
				return false
			}
		}
		return true
	}
}
func main() {
	isPalindrome(123)
}
