package string

import "log"
//给定一个排序数组和一个目标值，在数组中找到目标值，并返回其索引。如果目标值不存在于数组中，返回它将会被按顺序插入的位置。
//ok
func searchInsert1(nums []int, target int) int {
	defer log.Println(nums)
	var index, value int
	if len(nums) == 0 {
		return 0
	}
	if nums[0] > target {
		nums = append([]int{target}, nums...)

	}
	for index, value = range nums {
		if value > target {
			temp := append([]int{}, nums[:index]...)
			nums = append(nums[:index], target)
			nums = append(nums, temp...)
			return index
		}
	}
	nums = append(nums, target)

	log.Println(nums)
	return index

}
func searchInsert2(nums []int, target int) int {
	var index, value int
	if len(nums) == 0 {
		return 0
	}
	for index, value = range nums {
		if value >= target {
			return index
		}
	}
	return index + 1

}

func searchInsert(nums []int, target int) int {
	l := 0
	r := len(nums)
	for l<r{
		m:=(l+r)/2
		if nums[m]==target{
			return m
		}else if nums[m]<target{
			if m+1 >= r || nums[m+1] > target{//如果两个以下 或者nums  找到 两者之间
				return m+1
			}else{
				l=m+1 //继续二分查找
			}
		}else {
			if m==0 ||nums[m-1]<target{
				return m
			}else{
				r=m
			}
		}
	}
	return 0
}
