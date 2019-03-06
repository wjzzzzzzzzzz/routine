package num

import (
	"time"
	"log"
)

//假设你正在爬楼梯。需要 n 阶你才能到达楼顶。
//
//每次你可以爬 1 或 2 个台阶。你有多少种不同的方法可以爬到楼顶呢？
//
//注意：给定 n 是一个正整数。

//f(n)=f(n-1)+f(n-2) f(1)=1
//ok
func climbStairs1(n int) int {
	if n <= 3 {
		return n
	}

	log.Println(n)
	time.Sleep(time.Second)
	return climbStairs(n-1) + climbStairs(n-2)
}

func climbStairs(n int) int {
	if n <= 3 {
		return n
	}
	ints := make(map[int]int)
	return get(n-1,ints) + get(n-2,ints)
}

func get(n int, m map[int]int) int {
	if n <= 3 {
		return n
	}
	if v, ok := m[n]; ok {
		return v
	}
	result := get(n-1, m) + get(n-2, m)
	m[n] = result
	return result
}
