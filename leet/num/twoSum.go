package num

//给定一个整数数组 nums 和一个目标值 target，请你在该数组中找出和为目标值的那 两个 整数，并返回他们的数组下标。
//你可以假设每种输入只会对应一个答案。但是，你不能重复利用这个数组中同样的元素
//ok
func twoSum(nums []int, target int) []int {
	var result []int
	for in, va := range nums {
		va = target - va
		for i, v := range nums {
			if i <= in {
				continue
			}
			if v == va {
				return append(result, i, in)
			}
		}
	}
	return nil
}


