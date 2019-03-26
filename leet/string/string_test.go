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
	var s4 = []int{1, 3, 5, 5, 5, 6}
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
func Test_lengthOfLastWord(t *testing.T) {
	s := "abcdefg"
	log.Println(len(s))
	s = s[:len(s)-1]
	log.Println(s, len(s))
	a := " a"
	word := lengthOfLastWord(a)
	log.Println(word)
}
func Test_plusOne(t *testing.T) {
	a := []int{1}
	//word := plusOne(a)
	one := plusOne(a)
	log.Println(one)

}

func Test_addBinary(t *testing.T) {
	//切片包含左边,不包含右边
	a := "123"
	log.Println(a[0:3])
	b := []int{1, 2, 3}
	log.Printf("%T%v", b[2], b[2])
	log.Println(b[2])
}

func Test_isPalindrome1(t *testing.T) {

	s := "A"
	//log.Println(s[0],s[1])
	log.Println(string(s[0] + 1))
	log.Printf("T", s[0])
}

func Test_convertToTitle(t *testing.T) {
	s := "saaasss"
	log.Printf("%T", s[1])
}
func Test_isIsomorphic(t *testing.T) {
	log.Println(isIsomorphic("b", "a"))
}
func Test_lengthOfLongestSubstring(t *testing.T) {
	//log.Println( lengthOfLongestSubstring("aaudfac"))	//
	////log.Println( all("abcd"))
}
func Test_reverseString(t *testing.T) {
	bytes := []byte("123745")
	reverseString(bytes)
	log.Println(string(bytes))
	//log.Println( all("abcd"))
}
func Test_isAnagram(t *testing.T) {
	println(isAnagram("aaacccbbb", "abcabcac"))
}


func Test_letterCasePermutation(t *testing.T) {
	log.Println(letterCasePermutation("abcd"))
}