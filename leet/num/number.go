package num

//二分法

//给定一个数字和排好序的切片使用二分法 找到下标
func dichotomy(s []int, target int) int {
	var left, right, Middle int
	left = 0
	right = len(s)-1
	for left < right {
		Middle = (left + right) / 2
		if target > s[Middle] {
			left = Middle+1
		} else if target<s[Middle]{
			right = Middle-1
		}else{
			return Middle
		}
	}
	return -1
}
