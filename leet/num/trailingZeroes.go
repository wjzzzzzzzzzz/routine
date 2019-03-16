package num

import (
	"math"
)
//ok
//给定一个整数 n，返回 n! 结果尾数中零的数量。
func trailingZeroes(n int) int {
	if n <= 0 {
		return 0
	}
	i := MaLog(n, 5)
	var result int
	for i != -1 {
		result += n / MaPow(5, i)
		i--
	}
	return result
}
func MaLog(x int, di int) int {
	return int(math.Log(float64(x)) / math.Log(float64(di)))
}
func MaPow(x, y int) int {
	result := x
	for y != 0 {
		result *= x
		y--
	}
	return result
}
func trailingZeroes1(n int) int {
	res := 0
	for n > 0 {
		n = n / 5 //有多少个5 是这个意思
		res += n
	}

	return res
}
