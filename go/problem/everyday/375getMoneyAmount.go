package problem

import (
	"math"
	"mycode/leetcode/go/problem/common"
)

func GetMoneyAmount(n int) int {
	/*
		设dp[i][j] 为从[i,j]中猜出目标值的最小花费，dp[i][i] = 0
		则从[i,j]中猜出x的最小花费为：max(dp[i][x-1], dp[x+1][j]) + x
	*/
	const INT_MAX = math.MaxInt32
	var dp [210][210]int

	for i := n - 1; i >= 1; i-- {
		for j := i + 1; j <= n; j++ {
			dp[i][j] = INT_MAX
			for x := i; x < j; x++ {
				dp[i][j] = common.MyMinInt(dp[i][j], common.MyMaxInt(dp[i][x-1], dp[x+1][j])+x)
			}
		}
	}

	return dp[1][n]
}

func GetMoneyAmoun2t(n int) int {
	/*
		设dp[i][j] 为从[i,j]中猜出目标值的最小花费，dp[i][i] = 0
		则从[i,j]中猜出x的最小花费为：max(dp[i][x-1], dp[x+1][j]) + x
		由于状态转移方程为根据规模小的子问题计算规模大的子问题，因此计算子问题的顺序为先计算规模小的子问题，
		后计算规模大的子问题，需要注意循环遍历的方向
	*/
	const INT_MAX = math.MaxInt32
	var dp [210][210]int

	for i := n - 1; i >= 1; i-- {
		for j := i + 1; j <= n; j++ {
			minCost := INT_MAX
			for x := i; x < j; x++ {
				biggerCost := dp[i][x-1]  
				if dp[x+1][j] > dp[i][x-1] {
					biggerCost = dp[x+1][j]
				}
				tmpCost := biggerCost + x
				if tmpCost < minCost {
					minCost = tmpCost
				}
			}
			dp[i][j] = minCost
		}
	}

	return dp[1][n]
}
