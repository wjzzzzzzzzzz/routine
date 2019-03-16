package string

import "unicode"

//给定一个字符串，验证它是否是回文串，只考虑字母和数字字符，可以忽略字母的大小写。
//
//说明：本题中，我们将空字符串定义为有效的回文串。
//ok
func isPalindrome1(s string) bool {
	var i, j int
	j = len(s) - 1
	for i < j {
		if isV(s[i]) {
			for {
				if isV(s[j]) {
					if equal(s[i], s[j]) {
						j--
						break
					} else {
						return false
					}
				}
				j--
			}
		}
		i++
	}
	return true
}
func isV(i uint8) bool {
	return unicode.IsDigit(rune(i)) || unicode.IsLetter(rune(i))
}
func equal(a, b uint8) bool {
	if a==b{
		return true
	}
	if unicode.IsLetter(rune(a))&&unicode.IsLetter(rune(b))&&(a-b==32||b-a==32){
		return true
	}
	return false
}
