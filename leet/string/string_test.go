package string

import (
	"testing"
	"log"
)
func init(){
	log.SetFlags(log.Lshortfile | log.LstdFlags)
}
func Test_isPalindrome(t *testing.T) {
	palindrome := isPalindrome(123)
	log.Println(palindrome)

}
func Test_longestCommonPrefix(t *testing.T) {
	test := []string{
		"ca", "caaa", "ca",
	}
	prefix := longestCommonPrefix(test)
	log.Println(prefix)
}
