package top100

/*
输入：candidates = [2,3,6,7], target = 7
输出：[[2,2,3],[7]]

输入: candidates = [2,3,5], target = 8
输出: [[2,2,2,2],[2,3,3],[3,5]]

输入: candidates = [2], target = 1
输出: []
*/

func CombinationSum(candidates []int, target int) [][]int {
	res := new([][]int)

	var tmpArr []int
	combinationSumBacktrack(res, target, 0, tmpArr, candidates)
	return *res
}

func combinationSumBacktrack(res *[][]int, target, nowSum int, tmpArr, candidates []int) {
	/*
		回溯跳出的条件为：和为target
	*/
	if nowSum == target {
		*res = append(*res, tmpArr)
		return
	}
	for _, v := range candidates {
		if nowSum < target {
			combinationSumBacktrack(res, target, nowSum+v, append(tmpArr, v), candidates)
		}
	}

}

