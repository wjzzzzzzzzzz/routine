package num


//给定一个数组，它的第 i 个元素是一支给定股票第 i 天的价格。
//
//如果你最多只允许完成一笔交易（即买入和卖出一支股票），设计一个算法来计算你所能获取的最大利润。
//
//注意你不能在买入股票前卖出股票。
//示例 1:
//
//输入: [7,1,5,3,6,4]
//输出: 5
//解释: 在第 2 天（股票价格 = 1）的时候买入，在第 5 天（股票价格 = 6）的时候卖出，最大利润 = 6-1 = 5 。
//注意利润不能是 7-1 = 6, 因为卖出价格需要大于买入价格。
//f(n) = max(f(n-1) + A[n], A[n]);
//在n点卖出的话最大利润f（n）=max（f(n-1),前面最大的值-index（i）)）//
//动态规划 前i天的最大收益 = max{前i-1天的最大收益，第i天的价格-前i-1天中的最小价格}
func maxProfit(prices []int) int {
	if len(prices)==0{
		return 0
	}
	var maxPro int
	var Pro int
	var temp int
	min := prices[0]
	for i := 1; i < len(prices); i++ {
		if prices[i-1] < min {
			min = prices[i-1]
		}
		temp = prices[i] - min
		if Pro < temp {
			Pro = temp
		}
		if maxPro < Pro {
			maxPro = Pro
		}
	}
	return maxPro
}
