package num

//编写一个算法来判断一个数是不是“快乐数”。
//
//一个“快乐数”定义为：对于一个正整数，每一次将该数替换为它每个位置上的数字的平方和，然后重复这个过程直到这个数变为 1，也可能是无限循环但始终变不到 1。如果可以变为 1，那么这个数就是快乐数。
//
func isHappy(n int) bool {
	for i := 0; i < 100; i++ {
		if n == 1 {
			return true
		}
		n = square(Spilt(n))
	}
	return false
}

func Spilt(n int) []int {
	var res []int
	for {
		s := n / 10
		if s >= 10 {
			res = append(res, n%10)
			n = s
		} else {
			res = append(res, n/10)
			res = append(res, n%10)
			break
		}
	}
	return res
}
func square(s []int) int {
	var res int
	for _, v := range s {
		res += v * v
	}
	return res
}
