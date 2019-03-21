package num

import (
	"testing"
	"log"
	"encoding/hex"
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
	//s1 := []int{3, 2, 0, 4, 5}

	//log.Println(combinationSum(s1, 15))
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
		4, 5, 6, 7, 8, 7, 6, 5, 4}
	log.Println(singleNumber(s1))
	a := []byte{24, 0}
	log.Println(hex.EncodeToString(a))

}
func Test_mySqrt(t *testing.T) {
	log.Println(mySqrt(169))
}
func Test_climbStairs(t *testing.T) {
	log.Println(climbStairs(4))
}
func Test_majorityElement(t *testing.T) {
	log.Println(majorityElement([]int{1, 2, 2}))
}
func Test_trailingZeroes(t *testing.T) {
	var result int
	var ge int
	for i := 1; i <= 100; i++ {
		result *= i
	}
	log.Println(result)
	for result%10 == 0 {
		result = result / 10
		ge++
	}
	log.Println(ge)

}
func Test_trailingZeroes1(t *testing.T) {
	log.Println(trailingZeroes1(35))
}
func Test_twoSum1(t *testing.T) {
	log.Println(twoSum1([]int{1, 3, 4, 6, 7, 9}, 10))
}
func Test_dichotomy(t *testing.T) {
	log.Println(dichotomy([]int{1, 3, 4, 6, 7, 9}, 4))
}
func Test_containsDuplicate(t *testing.T) {
	log.Println(containsDuplicate([]int{}))
}
func Test_containsNearbyDuplicate(t *testing.T) {
	log.Println(containsNearbyDuplicate([]int{99, 99}, 2))
}
func Test_getCommKey(t *testing.T) {
	//log.Println(Test_getCommKey)
}
func Test_intersection(t *testing.T) {
	ints := make(map[int]int)
	log.Println(ints[1], ints)
	log.Println(intersection([]int{1, 2, 3, 4}, []int{3, 5}))
}
func Test_merge(t *testing.T) {
	ints := []int{1, 3, 5, 0, 0}
	merge(ints, 3,[]int{2,4}, 2)
	log.Println(ints)
}


