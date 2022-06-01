package juneinterval

/*
输入：nums = [1, 7, 3, 6, 5, 6]
输出：3
解释：
中心下标是 3 。
左侧数之和 sum = nums[0] + nums[1] + nums[2] = 1 + 7 + 3 = 11 ，
右侧数之和 sum = nums[4] + nums[5] = 5 + 6 = 11 ，二者相等。

*/

func PivotIndex(nums []int) int {
	/*
		典型前缀和：
		preSum = nums[0] + nums[1] + nums[2] + ... + nums[n-1]
		sumLeft[i] = preSum[i] - nums[i]
		sumRight[i] = preSum[n-1] - preSum[i]
		==> preSum[i] - nums[i] = preSum[n-1] - preSum[i]
		==> 2 * preSum[i] = preSum[n-1] + nums[i]

		或
		记数组的全部元素之和为 total，当遍历到第 ii 个元素时，设其左侧元素之和为 sum
		则其右侧元素之和为 total-nums_i-sum
		==> total-nums[i]-sum=sum
		==> 2 * sum + nums[i] = total
	*/

	/*
		var preSum []int

		tmpSum := 0
		for _, v := range nums {
			tmpSum += v
			preSum = append(preSum, tmpSum)
		}

		len1 := len(nums)
		for i := 0; i < len1; i++ {
			if 2 * preSum[i] == preSum[len1-1] + nums[i] {
				return i
			}
		}
		return -1
	*/


	totalSum := 0
	for _, v := range nums {
		totalSum += v
	}

	len1 := len(nums)
	sum := 0
	for i := 0; i < len1; i++ {
		if 2 * sum + nums[i] == totalSum {
			return i
		}
		sum += nums[i]
	}
	return -1
}