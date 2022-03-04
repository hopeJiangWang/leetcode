package top100

import "sort"

/*
输入: intervals = [[1,2],[2,3],[3,4],[1,3]]
输出: 1
解释: 移除 [1,3] 后，剩下的区间没有重叠。
*/

func EraseOverlapIntervals(intervals [][]int) int {
	/*
		1.按照右边界排序，从左向右记录非交叉区间的个数;
		2.最后用区间总数减去非交叉区间的个数就是需要移除的区间个数了
	*/

	lenOfIntervals := len(intervals)

	// 右边界排序
	sort.Slice(intervals, func(i, j int) bool {
		return intervals[i][1] < intervals[j][1]
	})

	count := 1 // 非交叉区间的个数
	pre := intervals[0]
	for i := 1; i < lenOfIntervals; i++ {
		cur := intervals[i]
		// 没有重合
		if pre[1] <= cur[0] {
			count++
			pre = cur // 更新结果集最大值
		} 
	}
	return lenOfIntervals - count
}