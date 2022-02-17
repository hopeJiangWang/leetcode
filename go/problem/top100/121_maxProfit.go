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
		设在第i天，minPrice 代表前i-1天的最低价，那么第i天的最大利润：
		如果需要卖出，则 maxProfit = price[i] - minPrice
		如果不卖出，则 maxProfit = 第i-1 的最大利润，即maxProfit不变
	*/
	len := len(prices)

	minPrice, maxProfit := prices[0], 0
	for i := 1; i < len; i++ {
		maxProfit = max(prices[i] - minPrice, maxProfit)
		minPrice = min(minPrice, prices[i])
		fmt.Printf("index: %d, maxProfit: %d, minPrice: %d \n", i, maxProfit, minPrice)
	}
	return maxProfit
}

func minInArry(arr []int, length int) int {
	minNum := 0

	for i := 0; i < length; i++ {
		if minNum >= arr[i] {
			minNum = arr[i]
		}
	}

	return minNum
}