package everyday

import "fmt"

/*
输入：strs = ["cba","daf","ghi"]
输出：1
解释：网格示意如下：
  cba
  daf
  ghi
列 0 和列 2 按升序排列，但列 1 不是，所以只需要删除列 1 。
*/

func MinDeletionSize(strs []string) int {
	var res int

	len1, len2 := len(strs), len(strs[0])

	// 逐列遍历
	for j := 0; j < len2; j++ {
		// 需要不断迭代当前数据，判断一整列是否都是字典序
		nowValue := strs[0][j]
		for i := 0; i < len1; i++ {
			if nowValue > strs[i][j] {
				res++
				fmt.Println("i: ", string(strs[0][j]), "i+1: ", string(strs[i][j]), "res:",   res)
				break
			} else {
				nowValue = strs[i][j]
			}
		}
	}

	return res
}