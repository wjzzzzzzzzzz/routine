package num

import (
	"log"
	"math"
)

//给定 n 个非负整数 a1，a2，...，an，每个数代表坐标中的一个点 (i, ai) 。在坐标内画 n 条垂直线，垂直线 i 的两个端点分别为 (i, ai) 和 (i, 0)。找出其中的两条线，使得它们与 x 轴共同构成的容器可以容纳最多的水。
//说明：你不能倾斜容器，且 n 的值至少为 2。

func maxArea1(height []int) int {
	result := 0
	for i := 0; i < len(height); i++ {
		for j := len(height) - 1; j > i; j-- {
			log.Println(i, j)
			temp := (j - i) * int(math.Min(float64(height[j]), float64(height[i])))
			if temp > result {
				result = temp
			}
		}
	}
	return result
}
func maxArea(height []int) int {
	i := 0
	j := len(height) - 1
	var result float64
	for j>i  {
		result=math.Max(float64(result),float64(j - i) * (math.Min(float64(height[j]), float64(height[i]))))
		if height[j]>height[i]{
			i++
		}else{
			j--
		}

	}
	return int(result)
}
