package string
//给定一个排序数组，你需要在原地删除重复出现的元素，使得每个元素只出现一次，返回移除后数组的新长度。
//
//不要使用额外的数组空间，你必须在原地修改输入数组并在使用 O(1) 额外空间的条件下完成。
func removeDuplicates1(nums []int) int {
	if len(nums) == 0 {
		return 0
	}
	if len(nums) == 1 {
		return 1
	}
	var temp, result int
	temp = nums[0]
	for _, v := range nums {
		if temp == v {
		} else {
			temp = v
			result++
		}
		nums[result] = temp //如果跟上一个相等 就把上个值按照result赋给他 ，如果不等则result++
	}
	return result + 1
}


func removeDuplicates(nums []int) int {
	if len(nums) == 0 {
		return 0
	}
	i := 0
	for j := 1; j < len(nums); j++ {
		if nums[j] != nums[i] {
			i++
			nums[i] = nums[j]
		}
	}
	return i + 1

}
