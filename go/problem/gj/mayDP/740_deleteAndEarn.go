package mayDP

func deleteAndEarn(nums []int) int {
	/*
		假设dp[i]代表i此时获取到的最大总点数，那么：dp[i] = max(dp[i-1], dp[i-2]+i*numsOfI)
		即为不选i或者选择i两种情况中的最大值
	*/

	// 先使用hash数组，记录数组中每个数字的个数
	hashSlice := make([]int, 10001)
	for i := 0; i < len(nums); i++ {
		hashSlice[nums[i]] += 1
	}

	// 初始化dp数组，0无正向作用
	dp := make([]int, 10001)
	dp[0], dp[1] = 0, hashSlice[1]

	for i := 2; i < len(hashSlice); i++ {
		dp[i] = max(dp[i-1], dp[i-2]+i*hashSlice[i])
	}
	return dp[len(dp)-1]
}

// 空间优化：因为每次只需要当前位置的前两个数据，我们用常量来记录更新即可
func deleteAndEarn2(nums []int) int {
	/*
		假设dp[i]代表i此时获取到的最大总点数，那么：dp[i] = max(dp[i-1], dp[i-2]+i*numsOfI)
		即为不选i或者选择i两种情况中的最大值
	*/

	// 先使用hash数组，记录数组中每个数字的个数
	hashSlice := make([]int, 10001)
	for i := 0; i < len(nums); i++ {
		hashSlice[nums[i]] += 1
	}

	// 初始化dp数组，0无正向作用
	// dp := make([]int, 10001)
	// dp[0], dp[1] = 0, hashSlice[1]
	first, second := 0, hashSlice[1]

	for i := 2; i < len(hashSlice); i++ {
		first, second = second, max(second, first+i*hashSlice[i])
		// dp[i] = max(dp[i-1], dp[i-2]+i*hashSlice[i])
	}
	return second
}