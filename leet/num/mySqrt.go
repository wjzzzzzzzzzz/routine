package num
//开平方
//ok
func mySqrt(x int) int {
	if x < 2 {
		return x
	}
	l := 0
	r := x
	m := 0
	m2 := 0
	for l < r {
		m = (l + r) / 2
		if l == m {
			return m
		}
		m2 = m * m
		if m2 == x {
			return m
		} else if m2 > x {
			r = m
		} else {
			l = m
		}
	}
	return 0
}
