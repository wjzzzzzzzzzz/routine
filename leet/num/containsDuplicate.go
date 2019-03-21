package num
//给定一个整数数组，判断是否存在重复元素。
//
//如果任何值在数组中出现至少两次，函数返回 true。如果数组中每个元素都不相同，则返回 false。
func containsDuplicate(nums []int) bool {
	ints := make(map[int]int)
	for i := 0; i < len(nums); i++ {
		if v, ok := ints[nums[i]]; !ok {
			ints[nums[i]] = v
		} else {
			return true
		}
	}
	return false
}
