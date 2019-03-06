package num


//ok

//给定一个非空整数数组，除了某个元素只出现一次以外，其余每个元素均出现两次。找出那个只出现了一次的元素。
//
//说明：
//
//你的算法应该具有线性时间复杂度。 你可以不使用额外空间来实现吗？
//必须全部遍历
//func singleNumber(nums []int) int {
//	if len(nums) == 1 {
//		return nums[0]
//	}
//	sort.Ints(nums)
//	log.Println(nums)
//	for i := 1; i < len(nums); i+=2 {
//		if nums[i-1] != nums[i] {
//			return nums[i-1]
//		}
//	}
//	return nums[len(nums)-1]
//}
//可以这样异或么？？？
func singleNumber(nums []int) int {
	var result int
	for i := 0; i < len(nums); i++ {
		result ^= nums[i]

	}
	return result
}
