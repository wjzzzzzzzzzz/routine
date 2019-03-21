package num
//给定一个整数数组和一个整数 k，判断数组中是否存在两个不同的索引 i 和 j，使得 nums [i] = nums [j]，并且 i 和 j 的差的绝对值最大为 k。
//
//示例 1:
func containsNearbyDuplicate(nums []int, k int) bool {
	ints := make(map[int][]int)
	for i := 0; i < len(nums); i++ {
		if v, ok := ints[nums[i]]; !ok {
			var in []int
			ints[nums[i]] = append(in, i)
		} else {
			t := v
			for j := 0; j < len(t); j++ {
				if i-t[j] <= k {
					return true
				}
			}
			ints[nums[i]] = append(v, i)
		}
	}
	return false
}
