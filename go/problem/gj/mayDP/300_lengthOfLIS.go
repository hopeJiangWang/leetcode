package mayDP

/*
输入：nums = [10,9,2,5,3,7,101,18]
输出：4
解释：最长递增子序列是 [2,3,7,101]，因此长度为 4 。

*/

/*
	方法一：动态规划
	O(n^2)
*/
func LengthOfLIS(nums []int) int {
	/*
		dp[i]表示以第i个数字结尾的最长递增子序列的长度
		dp[i] = max(dp[i], dp[j]+1)(0≤j<i且num[j]<num[i])

		res = max(dp[i])
	*/
	len1, res := len(nums), 0
	dp := make([]int, len1)

	for i := 0; i < len1; i++ {
		dp[i] = 1
		for j := 0; j < i; j++ {
			if nums[j] < nums[i] {
				dp[i] = max(dp[i], dp[j]+1)
			}
		}
		res = max(res, dp[i])
	}
	return res
}

/*
	方法二：贪心+二分查找
*/
func LengthOfLIS2(nums []int) int {
	/*
		设当前已求出的最长上升子序列的长度为len（初始时为 1），从前往后遍历数组nums，在遍历到nums[i] 时：
		(1)如果nums[i]>d[len] ，则直接加入到 d 数组末尾，并更新 len=len+1；
		(2)否则，在 d 数组中二分查找，找到第一个比 nums[i] 小的数 d[k]，并更新d[k+1]=nums[i]。
	
		以输入序列 [0, 8, 4, 12, 2] 为例：
			第一步插入 0，d = [0]；
			第二步插入 8，d = [0, 8]；
			第三步插入 4，d = [0, 4]；
			第四步插入 12，d = [0, 4, 12]；
			第五步插入 2，d = [0, 2, 12].
	*/
	len1, n := 1, len(nums)

	if n == 0 {
		return 0
	}

	d := make([]int, n+1)
	d[len1] = nums[0]

	for i := 1; i < n; i++ {
		// 如果nums[i]>d[len] ，则直接加入到 d 数组末尾，并更新 len=len+1；
		if nums[i] > d[len1] {
			len1++
			d[len1] = nums[i]
		} else {
			// 否则，在 d 数组中二分查找，找到第一个比 nums[i] 小的数 d[k]，并更新d[k+1]=nums[i]。
			l, r, k := 1, len1, 0
			for l <= r {
				mid := (l + r) >> 1
				if d[mid] < nums[i] {
					k = mid
					l = mid + 1
				} else {
					r = mid - 1
				}
			}
			d[k+1] = nums[i]
		}
	}

	return len1
}