package num

import (
	"testing"
	"log"
)

func init() {
	log.SetFlags(log.Lshortfile | log.LstdFlags)
}
func Test_reverse(t *testing.T) {
	i := reverse(1234)
	log.Println(i)

}
func Test_twoSum(t *testing.T) {
	sum := twoSum([]int{3, 2, 0, 4, 5, 3}, 6)
	log.Println(sum)
}
func Test_removeElement(t *testing.T) {
	s1 := []int{3, 3, 2, 0, 4, 5, 3, 3}
	s2 := []int{1}
	s3 := []int{3, 3, 2, 0, 4, 5, 3, 3}
	log.Println(removeElement(s1, 3), s1)
	log.Println(removeElement(s2, 5), s2)
	log.Println(removeElement(s3, 4), s3)
}
func Test_maxArea(t *testing.T) {
	s1 := []int{3, 3, 2, 0, 4, 5, 3, 3}
	//s2 := []int{1}
	//s3 := []int{4, 3, 3}

	log.Println(maxArea(s1))
}
func Test_combinationSum(t *testing.T) {
	s1 := []int{3, 2, 0, 4, 5}

	log.Println(combinationSum(s1, 15))
}
func Test_maxSubArray(t *testing.T) {
	s1 := []int{3, 2, 0, 4, 5}

	log.Println(maxSubArray(s1))
}
func Test_maxProfit(t *testing.T) {
	s1 := []int{7, 6, 5, 4, 3, 2}
	log.Println(maxProfit(s1))
}
func Test_maxProfit2(t *testing.T) {
	s1 := []int{7, 1, 5, 3, 6, 4}
	log.Println(maxProfit2(s1))
}
func Test_singleNumber(t *testing.T) {
	s1 := []int{
		4, 5, 6, 7,8 ,7, 6, 5, 4}
	log.Println(singleNumber(s1))

}
