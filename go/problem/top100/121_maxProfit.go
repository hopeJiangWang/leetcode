package top100

import (
	"fmt"
)

/*
你只能选择 某一天 买入这只股票，并选择在 未来的某一个不同的日子
卖出该股票。设计一个算法来计算你所能获取的最大利润。

输入：[7,1,5,3,6,4]
输出：5
解释：在第 2 天（股票价格 = 1）的时候买入，在第 5 天（股票价格 = 6）的时候卖出，最大利润 = 6-1 = 5 。
     注意利润不能是 7-1 = 6, 因为卖出价格需要大于买入价格；同时，你不能在买入前卖出股票。

输入：prices = [7,6,4,3,1]
输出：0
解释：在这种情况下, 没有交易完成, 所以最大利润为 0。

*/

func MaxProfit(prices []int) int {
	/*
		minPrice：第i日卖的时候，前面i-1天里面的最小价格
		maxProfit：第i日卖的时候，此时的最大利润，也即price[i]-min(price[0:i-1])
	*/

	minPrice, maxProfit := prices[0], 0

	for i := 0; i < len(prices); i++ {
		minPrice = min(minPrice, prices[i])
		maxProfit = max(maxProfit, prices[i]-minPrice)
		fmt.Printf("maxProfit: %v, minPrice: %v \n", maxProfit, minPrice)
	}

	return maxProfit
}
