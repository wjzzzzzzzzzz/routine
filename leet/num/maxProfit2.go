package num

//给定一个数组，它的第 i 个元素是一支给定股票第 i 天的价格。
//
//设计一个算法来计算你所能获取的最大利润。你可以尽可能地完成更多的交易（多次买卖一支股票）。
//
//注意：你不能同时参与多笔交易（你必须在再次购买前出售掉之前的股票）。

//跳过中间的点会使得利润降低 。因此只要在降价之前卖出就可以
//一直大于 不给index赋值 ，小于 price[i-1]-price[index] //检测不到最后一个值
//[7,1,5,3,6,4]
//ok

func maxProfit2(prices []int) int {
	if len(prices) == 0 {
		return 0
	}
	var Pro int
	for i := 1; i < len(prices); i++ {
		temp := prices[i] - prices[i-1]
		if temp>0{
			Pro+=temp
		}

	}
	return Pro
}

//func maxProfit2(prices []int) int {
//	if len(prices) == 0 {
//		return 0
//	}
//	var Pro int
//	var index int
//	var peak,valley int
//
//	l:=len(prices)
//	for i := 1; i < l-1; i++ {
//		if prices[i-1] > prices[i] && prices[i] < prices[i+1] {//先找谷
//			valley=i
//		}
//
//		if prices[i-1] < prices[i] && prices[i] > prices[i+1] {//先找峰
//			peak=i
//		}
//		if peak*valley!=0{
//			Pro+=prices[peak]-prices[valley]
//			peak=0
//			valley=0
//		}
//
//
//
//	}
//	if index==0&&prices[0]<prices[l-1]{
//		return  prices[l-1]-prices[0]
//	}
//	return Pro
//}
