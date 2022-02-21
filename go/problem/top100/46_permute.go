package top100

import "fmt"

/*
输入：nums = [1,2,3]
输出：[[1,2,3],[1,3,2],[2,1,3],[2,3,1],[3,1,2],[3,2,1]]

输入：nums = [0,1]
输出：[[0,1],[1,0]]

输入：nums = [1]
输出：[[1]]
*/

func Permute(nums []int) [][]int {
	var res [][]int
	var tmpArr []int
	used := make([]int, 7)
	backtrackPermute(nums, tmpArr, used, 0, &res)

	return res
}

func backtrackPermute(nums, tmpArr, used []int, cnt int, res *[][]int) {
	/*
		回溯终点：cnt == len(nums)
	*/
	if cnt == len(nums) {
		tmp := make([]int, len(tmpArr))
		copy(tmp, tmpArr)
		*res = append(*res, tmp)
		return
	}

	for i := 0; i < len(nums); i++ {
		// 如果这个数字已经用过了，就跳过它
		if used[nums[i]] == 1 {
			continue
		}

		cnt++
		tmpArr = append(tmpArr, nums[i])
		used[nums[i]] = 1
		fmt.Printf("used: %v, cnt: %v, tmpArr: %v, res: %v \n", used, cnt, tmpArr, res)
		backtrackPermute(nums, tmpArr, used, cnt, res)
		tmpArr = tmpArr[:len(tmpArr)-1]
		cnt--
		used[nums[i]] = 0
	}
}
