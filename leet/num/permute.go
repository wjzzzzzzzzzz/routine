package num

//给定一个没有重复数字的序列，返回其所有可能的全排列。

func permute(nums []int) [][]int {
	if len(nums) == 0 {
		return [][]int{}
	}
	//var res [][]int

	//for _, v := range nums {
	//	for _, value := range res {
	//		//temp := insert(v, value)
	//	}
	//}

	return nil
}

func insert(a int, m []int) [][]int {
	var res [][]int
	for i := 0; i <= len(m); i++ {
		var temp []int
		temp = append([]int{a}, m[i:]...)
		temp = append(m[:i], temp...)
		res = append(res, temp)
	}
	return res
}
