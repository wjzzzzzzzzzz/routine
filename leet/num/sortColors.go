package num



//冒泡排序
func sortColors(nums []int) {
	for i := 0; i < len(nums); i++  {
		for j := 0; j< len(nums)-i-1; j++ {
			if nums[j]>nums[j+1]{
				nums[j],nums[j+1]=nums[j+1],nums[j]
			}

		}
	}
}
