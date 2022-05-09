package mayDP

import "sort"

/*
输入：events = [[1,2,4],[3,4,3],[2,3,1]], k = 2
输出：7
解释：选择绿色的活动会议 0 和 1，得到总价值和为 4 + 3 = 7 。

输入：events = [[1,2,4],[3,4,3],[2,3,10]], k = 2
输出：10
解释：参加会议 2 ，得到价值和为 10 。
你没法再参加别的会议了，因为跟会议 2 有重叠。你 不 需要参加满 k 个会议。

输入：events = [[1,1,1],[2,2,2],[3,3,3],[4,4,4]], k = 3
输出：9
解释：尽管会议互不重叠，你只能参加 3 个会议，所以选择价值最大的 3 个会议。
*/

func maxValue(events [][]int, k int) int {
	/*
		定义dp[i][j]为考虑前i个事件，选择不超过j的最大价值
		（1）不选 i，dp[i][j]=dp[i-1][j]
		（2）选择 i，找到第 i 件事件之前，与第 i 件事件不冲突的事件，记为last，
		则 dp[i][j]=dp[last][j-1]+value_i
		--> dp[i][j] = max(dp[i-1][j], dp[last][j-1]+value_i)

		其实就是找last，对events的结束时间排序，然后从右往左找，找到第一个满足结束时间小于当前事件的开始时间的事件，
		就是last
	*/
	sort.Slice(events, func(i, j int) bool {
		return events[i][1] < events[j][1]
	})

	len1 := len(events)
	dp := make([][]int, len1+1)
	for i := range dp {
		dp[i] = make([]int, k+1)
	}

	for i := 1; i <= len1; i++ {
		start, value := events[i-1][0], events[i-1][2]

		last := 0
		for j := i-1; j >= 1; j-- {
			if events[j-1][1] < start {
				last = j
				break
			}
		}

		for p := 1; p <= k; p++ {
			dp[i][p] = max(dp[i-1][p], dp[last][p-1]+value)
		}
	}

	return dp[len1][k]
}

func maxValue2(events [][]int, k int) int {
	/*
		定义dp[i][j]为考虑前i个事件，选择不超过j的最大价值
		（1）不选 i，dp[i][j]=dp[i-1][j]
		（2）选择 i，找到第 i 件事件之前，与第 i 件事件不冲突的事件，记为last，
		则 dp[i][j]=dp[last][j-1]+value_i
		--> dp[i][j] = max(dp[i-1][j], dp[last][j-1]+value_i)

		其实就是找last，对events的结束时间排序，然后从右往左找，找到第一个满足结束时间小于当前事件的开始时间的事件，
		就是last

		可以使用二分去查找
	*/
	sort.Slice(events, func(i, j int) bool {
		return events[i][1] < events[j][1]
	})

	len1 := len(events)
	dp := make([][]int, len1+1)
	for i := range dp {
		dp[i] = make([]int, k+1)
	}

	for i := 1; i <= len1; i++ {
		start, value := events[i-1][0], events[i-1][2]

		last := 0
		// for j := i-1; j >= 1; j-- {
		// 	if events[j-1][1] < start {
		// 		last = j
		// 		break
		// 	}
		// }
		l, r := 1, i-1
		for l < r {
			mid := (l + r + 1) >> 1
			if events[mid-1][1] < start {
				l = mid
			} else {
				r = mid - 1
			}
		}

		if r > 0 && events[r-1][1] < start {
			last = r
		}

		for p := 1; p <= k; p++ {
			dp[i][p] = max(dp[i-1][p], dp[last][p-1]+value)
		}
	}

	return dp[len1][k]
}