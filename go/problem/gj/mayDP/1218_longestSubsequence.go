package mayDP

/*
输入：arr = [1,5,7,8,5,3,4,2,1], difference = -2
输出：4
解释：最长的等差子序列是 [7,5,3,1]。

输入：arr = [1,3,5,7], difference = 1
输出：1
解释：最长的等差子序列是任意单个元素。

输入：arr = [1,2,3,4], difference = 1
输出：4
解释：最长的等差子序列是 [1,2,3,4]。
*/

func LongestSubsequence(arr []int, difference int) int {
	/*
		dp[i]表示以arr[i]结尾的等差子序列的最长长度，
		（1）那么如果我们可以在arr[i]左侧找到满足arr[j]：arr[j]=arr[i]-difference, 
			dp[i] = dp[j] + 1
		以dp[v]表示以v为结尾的最长等差子序列的长度，那么：
			dp[v] = dp[v-d] + 1
	*/
	res := 0
	dp := map[int]int{}
	for _, v := range arr {
		dp[v] = dp[v-difference] + 1
		res = max(res, dp[v])
	}
	return res
}

func max(x, y int) int {
	if x < y {
		return y
	}
	return x
}