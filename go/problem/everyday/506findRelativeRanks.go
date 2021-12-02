package problem

import (
	"sort"
	"strconv"
)
/*
输入：score = [10,3,8,9,4]
输出：["Gold Medal","5","Bronze Medal","Silver Medal","4"]

*/

func FindRelativeRanks(score []int) []string {
	res := make([]string, len(score))
    desc := [3]string{"Gold Medal", "Silver Medal", "Bronze Medal"}
    
	type pair struct {
		score int
		idx int
	}

	var arr []pair
    for i := 0; i < len(score); i++ {
		var tmpArr pair

		tmpArr.score = score[i]
		tmpArr.idx = i
        arr = append(arr, tmpArr)
    }

    // 排序（分数高的在前）
	sort.Slice(arr, func(i, j int) bool {
		return arr[i].score > arr[j].score
	})

    for i := 0; i < len(score); i++ {
        if i >= 3 {
            res[arr[i].idx] = strconv.Itoa(i + 1)
        } else {
            res[arr[i].idx] = desc[i];
        }
    }

	return res
}