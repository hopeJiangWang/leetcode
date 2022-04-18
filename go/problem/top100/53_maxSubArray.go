package top100

import (
	"fmt"
)

/*
输入：nums = [-2,1,-3,4,-1,2,1,-5,4]
输出：6
解释：连续子数组 [4,-1,2,1] 的和最大，为 6 。

输入：nums = [1]
输出：1

输入：nums = [5,4,-1,7,8]
输出：23
*/

func MaxSubArray3(nums []int) int {
	/*
		动态规划：
		设i的结尾的连续子数组的和为dp[i]:
		dp[i] = max(dp[i-1] + nums[i], nums[i])
	*/
	var res int
	numsLen := len(nums)

	dp := make([]int, numsLen)
	res, dp[0] = nums[0], nums[0]

	for i := 1; i < numsLen; i++ {
		dp[i] = max(dp[i-1] + nums[i], nums[i])
		res = max(res, dp[i])
	}
	return res
}

func MaxSubArray1(nums []int) int {
	/*
		最简单直观的：直接双指针查找即可（O(n^2)）
		超过时间限制~~
	*/
	var res int = -10001

	numsLen := len(nums)
	if numsLen == 0 {
		return 0
	}

	for i := 0; i < numsLen; i++ {
		for j := i; j < numsLen; j++ {
			res = max(res, sumArr(nums, i, j))
		}
	}
	return res
}

func sumArr(nums []int, begin, end int) int {
	var res int = 0
	for i := begin; i <= end; i++ {
		res += nums[i]
	}
	return res
}

func MaxSubArray(nums []int) int {
	/*
		不能直接使用双指针来进行查找
		动态规划：nums = [-2,1,-3,4,-1,2,1,-5,4]
		假设dp[i]表示以i结尾的连续子序列的最大和，dp[0]=nums[0]
		dp[i] = max(dp[i-1]+nums[i], nums[i])
	*/
	var res int
	numsLen := len(nums)
	if numsLen == 0 {
		return res
	}

	dp := make([]int, numsLen)
	dp[0], res = nums[0], nums[0]

	for i := 1; i < numsLen; i++ {
		dp[i] = max(dp[i-1]+nums[i], nums[i])
		res = max(res, dp[i])
		//fmt.Println("i: ", i, "dp: ", dp[i], "nums: ", nums[i])
	}

	fmt.Println(dp)
	return res
}


