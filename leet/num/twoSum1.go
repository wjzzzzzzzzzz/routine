package num

/*给定一个已按照升序排列 的有序数组，找到两个数使得它们相加之和等于目标数。

函数应该返回这两个下标值 index1 和 index2，其中 index1 必须小于 index2。

说明:

返回的下标值（index1 和 index2）不是从零开始的。
你可以假设每个输入只对应唯一的答案，而且你不可以重复使用相同的元素。*/
//很慢
func twoSum1(numbers []int, target int) []int {
	var i, j int
	for i = 0; i < len(numbers); i++ {
		for j = i + 1; j < len(numbers); j++ {
			if numbers[i]+numbers[j] == target {
				return append([]int{i+1}, j+1)
			}
		}
	}
	return nil
}
