package await

//给定一个含有 n 个正整数的数组和一个正整数 s ，找出该数组中满足其和 ≥ s 的长度最小的连续子数组。如果不存在符合条件的连续子数组，返回 0。
func minSubArrayLen(s int, nums []int) int {
	var size int
	for size = 1; size <= len(nums); size++ {
		var i int
		res := sum(nums[i : i+size])
		for i = 1; i <= len(nums); i++ {
		}

	}
	return 0
}
func sum(nums []int) int {
	var res int
	for i := 0; i < len(nums); i++ {
		res += nums[i]
	}
	return res
}
