package everyday

import "sort"

/*
输入: matchsticks = [1,1,2,2,2]
输出: true
解释: 能拼成一个边长为2的正方形，每边两根火柴。

输入: matchsticks = [3,3,3,3,4]
输出: false
解释: 不能用所有火柴拼成一个正方形。

*/

func Makesquare(matchsticks []int) bool {
	// 先统计所有火柴的长度
	totalLen := 0
	for _, l := range matchsticks {
		totalLen += l
	}

	// 如果长度不是4的整数倍，直接返回失败即可
	if totalLen%4 != 0 {
		return false
	}
	
	sort.Sort(sort.Reverse(sort.IntSlice(matchsticks))) // 减少搜索量

	edges := [4]int{}
	var dfs func(int) bool
	dfs = func(i int) bool {
		// 到最后一根火柴了
		if i == len(matchsticks) {
			return true
		}

		for e := range edges {
			edges[e] += matchsticks[i]
			if edges[e] <= totalLen/4 && dfs(i+1) {
				return true
			}
			edges[e] -= matchsticks[i]
		}
		return false
	}
	
	return dfs(0)


	// edges := [4]int{}
	// var dfs func(int) bool
	// dfs = func(idx int) bool {
	// 	if idx == len(matchsticks) {
	// 		return true
	// 	}
	// 	for i := range edges {
	// 		edges[i] += matchsticks[idx]
	// 		// 如果当前边还未满，那就继续试下一根火柴
	// 		if edges[i] <= totalLen/4 && dfs(idx+1) {
	// 			return true
	// 		}
	// 		// 否则，这个火柴不匹配这条边，需要剔除
	// 		edges[i] -= matchsticks[idx]
	// 	}
	// 	return false
	// }
	// return dfs(0)
}