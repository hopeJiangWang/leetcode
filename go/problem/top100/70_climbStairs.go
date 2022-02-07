package top100

import (
	"math"
)

/*
	每次你可以爬 1 或 2 个台阶。你有多少种不同的方法可以爬到楼顶（n阶）呢？
*/

// 普通递归超时 44 
func ClimbStairs1(n int) int {
	/*
		这个典型的递归、动规场景：
		1. 递归：f[n] 代表爬n阶台阶可能的方法，f[1] = 1, f[2] = 2, f[3] = 3
		只能从第n-1或者n-2阶走上去，即：f[n] = f[n-1] + f[n-2]
	*/
	if n == 1 {
		return 1
	} else if n == 2 {
		return 2
	}

	return ClimbStairs1(n - 1) + ClimbStairs1(n - 2)
}

// 记忆优化
func ClimbStairs2(n int) int {
	/*
		这个典型的递归、动规场景：
		1. 递归：f[n] 代表爬n阶台阶可能的方法，f[1] = 1, f[2] = 2, f[3] = 3
		只能从第n-1或者n-2阶走上去，即：f[n] = f[n-1] + f[n-2]
	*/
	if n == 1 {
		return 1
	} else if n == 2 {
		return 2
	}

	dp := make([]int64, n + 1)

	dp[1] = 1
	dp[2] = 2

	for i := 3; i <= n; i++ {
		dp[i] = dp[i-1] + dp[i-2]
	}

	return int(dp[n])
}

// 存储优化
func ClimbStairs3(n int) int {
	/*
		这个典型的递归、动规场景：
		1. 递归：f[n] 代表爬n阶台阶可能的方法，f[1] = 1, f[2] = 2, f[3] = 3
		只能从第n-1或者n-2阶走上去，即：f[n] = f[n-1] + f[n-2]
	*/
	if n == 1 {
		return 1
	} else if n == 2 {
		return 2
	}
	
	tmp, a, b := 0, 1, 2
	for i := 3; i <= n; i++ {
		tmp = a
		a = b
		b = tmp + a
	}

	return b
}

// 数学方法：二阶线性递推数列直接计算（注意这里的n，需要多加一）
func ClimbStairs(n int) int {
	sqrt5 := math.Sqrt(5)
	a := math.Pow((1 + sqrt5) / 2, float64(n + 1))
	b := math.Pow((1 - sqrt5) / 2, float64(n + 1))
	return int(math.Round((a - b) / sqrt5))
}