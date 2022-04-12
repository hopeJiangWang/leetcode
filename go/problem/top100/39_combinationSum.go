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
	var res [][]int
	var tmpArr []int

	backtrackCombinationSum(candidates, tmpArr, target, 0, 0, &res)
	return res
}

func backtrackCombinationSum(candidates, tmpArr []int, target, startIndex, nowSum int, res *[][]int) {
	/*
		回溯出口：和达到目标值即可
	*/
	if nowSum == target {
		// 因为这里用的切片，有指针指向底层数据，后续的修改会对tmpArr也造成影响
		// 所以需要新建一个tmp
		tmp := make([]int, len(tmpArr))
		copy(tmp, tmpArr)        // 拷贝
		*res = append(*res, tmp) // 放入结果集
		tmpArr = []int{}         // 清空临时数组
		return
	}

	// 剪枝
	if nowSum > target {
		return
	}

	// 这里保存初始下标，为了避免重复
	for i := startIndex; i < len(candidates); i++ {
		tmpArr = append(tmpArr, candidates[i])
		nowSum += candidates[i]
		backtrackCombinationSum(candidates, tmpArr, target, i, nowSum, res)
		// 回溯
		tmpArr = tmpArr[:len(tmpArr)-1]
		nowSum -= candidates[i]
	}
}
