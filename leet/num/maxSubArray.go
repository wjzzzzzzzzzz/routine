package num
//给定一个整数数组 nums ，找到一个具有最大和的连续子数组（子数组最少包含一个元素），返回其最大和。
//示例:
//输入: [-2,1,-3,4,-1,2,1,-5,4],
//输出: 6
//解释: 连续子数组 [4,-1,2,1] 的和最大，为 6。
//
//方法如下：
//
//定义一个函数f(n)，以第n个数为结束点的子数列的最大和，存在一个递推关系f(n) = max(f(n-1) + A[n], A[n]);
//将这些最大和保存下来后，取最大的那个就是，最大子数组和。因为最大连续子数组 等价于 最大的以n个数为结束点的子数列和 附代码
//ok
func maxSubArray(nums []int) int {
	if len(nums) <= 1 {
		return nums[0]
	}
	var SumNum int
	MaxNum := nums[0]
	for _, num := range nums {
		SumNum = SumNum + num
		if SumNum < num {//如果加上num还比num小，则抛弃之前的积累,从num 开始 f(n) = max(f(n-1) + A[n], A[n]);
			SumNum = num
		}

		if SumNum > MaxNum {//取下最大的值
			MaxNum = SumNum
		}
	}

	return MaxNum
}















