package mayDP

/*
输入：target = [5,1,3], arr = [9,4,2,3,4]
输出：2
解释：你可以添加 5 和 1 ，使得 arr 变为 [5,9,4,1,2,3,4] ，target 为 arr 的子序列。

输入：target = [6,4,8,1,3,2], arr = [4,7,6,2,3,8,6,1]
输出：3
*/

func MinOperations(target []int, arr []int) int {
	/*
		dp[i][j]：表示长度为[0,i-1]的字符串text1与长度为[0,j-1]的字符串text2的最长公共子序列的长度；
		1、text1[i-1] == text2[j-1]：那么找到了一个共同元素，dp[i][j] = dp[i-1][j-1] + 1
		2、text1[i-1] != text2[j-1]：那么这时候就要找之前的最大值了，dp[i][j] = max(dp[i][j-1], dp[i-1][j])
	
		可以先找出两个序列的最长公共子序列的长度，然后再看target中还有几个剩余就是还需要添加几个了

		https://leetcode-cn.com/problems/minimum-operations-to-make-a-subsequence/solution/gong-shui-san-xie-noxiang-xin-ke-xue-xi-oj7yu/
		思路：找到最大公共子序列的长度len，然后res=len1-len
		思路没问题，但是时间复杂度为O(mn)，会超时
		注意到target中的数据互不相同：即将LCS（最长公共子序列）->LIS（最长上升子序列），O(mn)->O(nlogn)
		同时最长上升子序列问题（LIS）存在使用「维护单调序列 + 二分」的贪心解法，复杂度为 O(nlogn)。

		因此本题可以通过「抽象成 LCS 问题」->「利用 target 数组元素各不相同，转换为 LIS 问题」->「使用 LIS 的贪心解法」，做到O(nlogn) 的复杂度。
		==》
		*/
	len1, len2 := len(target), len(arr)
	
	/*
		思路：把arr中包含target中的整数及其在target中的索引按顺序缓存下来一个数组，然后计算这个数组的最长递增子序列即可。
		确定基本思路之后，就可以先将LCS转换成LIS
		输入：target = [6,4,8,1,3,2], arr = [4,7,6,2,3,8,6,1]
		==>
		arr* = [1,0,5,4,2,0,3]
		target* = [0,1,2,3,4,5]
		==> 即转换成求新的两个数组的最长公共子序列
		因为target是严格递增的，那么等价于求新的arr*的最长递增子序列
	*/

	// 获取target中的数据及其对应的索引的关系
	hashMap := make(map[int]int)
	for i := 0; i < len1; i++ {
		hashMap[target[i]] = i
	}

	// 把arr中包含target中的整数及其在target中的索引按顺序缓存下来一个数组
	var newIndexArr []int
	for i := 0; i < len2; i++ {
		_, ok := hashMap[arr[i]] 
		if ok {
			newIndexArr = append(newIndexArr, hashMap[arr[i]])
		}
	}

	// 如果arr中没有一个target中的数据，那么说明全部需要添加
	if len(newIndexArr) == 0 {
		return len1
	}

	// 否则，就是找新数组的最长递增子序列
	return len1 - lengthOfLIS(newIndexArr)
}

func lengthOfLIS(nums []int) int {
	/*
		设当前已求出的最长上升子序列的长度为len（初始时为 1），从前往后遍历数组nums，在遍历到nums[i] 时：
		(1)如果nums[i]>d[len] ，则直接加入到 d 数组末尾，并更新 len=len+1；
		(2)否则，在 d 数组中二分查找，找到第一个比 nums[i] 小的数 d[k]，并更新d[k+1]=nums[i]。
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