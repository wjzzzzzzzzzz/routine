package num

//给定一个大小为 n 的数组，找到其中的众数。众数是指在数组中出现次数大于 ⌊ n/2 ⌋ 的元素。
//
//你可以假设数组是非空的，并且给定的数组总是存在众数。
//
//ok
func majorityElement(nums []int) int {
	if len(nums)==1{
		return nums[0]
	}
	l := len(nums) / 2
	ints := make(map[int]int)

	for _, va := range nums {
		if _, ok := ints[va];!ok{
			ints[va]=1
		}else {
			ints[va]++
			if ints[va]>l{
				return va
			}
		}

	}
	return 0
}
