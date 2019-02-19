package main

import (
	"strconv"
)

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

}
