package string
//给定一个由整数组成的非空数组所表示的非负整数，在该数的基础上加一。
//
//最高位数字存放在数组的首位， 数组中每个元素只存储一个数字。
//
//你可以假设除了整数 0 之外，这个整数不会以零开头。
//ok
func plusOne(digits []int) []int {
	var index int
	if len(digits) == 0 {
		return nil
	}
	index = len(digits) - 1
	for digits[index] == 9 {
		digits[index] = 0
		index--
		if index==-1 {
			break
		}
	}
	if index==-1{
		digits = append([]int{1}, digits...)
		return digits
	}
	digits[index] += 1
	return digits
}
