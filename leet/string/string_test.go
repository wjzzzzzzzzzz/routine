package string

import (
	"testing"
	"log"
)

func init() {
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
func Test_isValid(t *testing.T) {
	log.Println(isValid("[(){}]"))
}
func Test_removeDuplicates(t *testing.T) {
	var s1 = []int{0, 0, 0, 0, 0, 0, 1, 2, 2, 2}
	var s2 = []int{15, 15, 17, 17}
	var s3 = []int{1, 2, 2, 3, 3, 3, 3}
	var s4 = []int{3, 5, 5, 5, 6}
	log.Println(removeDuplicates(s1), s1)
	log.Println(removeDuplicates(s2), s2)
	log.Println(removeDuplicates(s3), s3)
	log.Println(removeDuplicates(s4), s4)
}
func Test_searchInsert(t *testing.T) {
	var s1 = []int{0, 0, 0, 0, 0, 0, 1, 2, 2, 2}
	var s2 = []int{15, 15, 17, 17}
	var s3 = []int{1, 2}
	var s4 = []int{1,3, 5, 5, 5, 6}
	log.Println(searchInsert(s1, 5))
	log.Println(searchInsert(s2, 6))
	log.Println(searchInsert(s3, 7))
	log.Println(searchInsert(s4, 3))
}
func Test_searchInsert1(t *testing.T) {
	var s1 = []int{0, 0, 0, 0, 0, 0, 1, 2, 2, 2}
	//var s2 = []int{15, 15, 17, 17}
	var s3 = []int{1, 2, 2, 3, 3, 3, 3}
	//var s4 = []int{1,3, 5, 5, 5, 6}
	searchInsert(s1, 5)
	//log.Println(searchInsert(s2, 6))
	searchInsert(s3, 7)
	//log.Println(searchInsert(s4, 3))
	log.Println(s1, s3)
}
