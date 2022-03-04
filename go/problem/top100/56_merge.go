package top100

import (
	"fmt"
	"sort"
)

/*
输入：intervals = [[1,3],[2,6],[8,10],[15,18]]
输出：[[1,6],[8,10],[15,18]]
解释：区间 [1,3] 和 [2,6] 重叠, 将它们合并为 [1,6].

输入：intervals = [[1,4],[4,5]]
输出：[[1,5]]
解释：区间 [1,4] 和 [4,5] 可被视为重叠区间。
*/

func Merge1(intervals [][]int) [][]int {
	var res [][]int

	/*
		1.按区间起点来排序；
		2.然后从左至右，判断相邻区间能否合并：
			intervals[i][0] <= intervals[j][0]，则可以，更新新边界：end = max(end,intervals[j][1])
		
	*/
	lenOfIntervals := len(intervals)
	if lenOfIntervals == 0 {
		return res
	}

	sort.Slice(intervals, func(i, j int) bool { 
		return intervals[i][0] <= intervals[j][0]
	})

	fmt.Println("intervals: ", intervals)
	res = append(res, []int{intervals[0][0], intervals[0][1]})
	lenOfRes := 1
	for i := 1; i < lenOfIntervals; i++ {
		begin := intervals[i][0]
		end := intervals[i][1]
		// 当前区间的最小值大于结果集中的最大值，则不会重合，直接插入即可
		if begin > res[lenOfRes-1][1] {
			res = append(res, []int{begin, end})
			lenOfRes++
		} else {
			// 否则，则能重合，更新结果集的最大值即可
			res[lenOfRes-1][1] = max(res[lenOfRes-1][1], end)
		}
		
		fmt.Println("res: ", res)
	}

	return res
}

func Merge(intervals [][]int) [][]int {
	var res [][]int
	lenOfIntervals := len(intervals)
	if lenOfIntervals == 0 {
		return res
	}

	sort.Slice(intervals, func(i, j int) bool { 
		return intervals[i][0] <= intervals[j][0]
	})

	pre := intervals[0]

	for i := 1; i < lenOfIntervals; i++ {
		cur := intervals[i]
		// 没有重合
		if pre[1] < cur[0] {
			res = append(res, pre)
			pre = cur // 更新结果集最大值
		} else {
			// 有重合，更新结果集最大值
			pre[1] = max(pre[1], cur[1])
		}
	}

	// 最后一个要记得存入结果集
	res = append(res, pre)

	return res
}
