package everyday

func findDuplicates(nums []int) []int {
	var res []int

	/*
		利用值域与数字下标空间大小的等同关系，使得每个数出现在它应该出现的位置：对于值为 
		k的数字，我们将其应该出现在的位置定为 k−1。

		1、遍历数组，nums[nums[i]-1]为正数，说明没出现过，加上负号
		2、否则出现过，加入答案
	*/

	for _, v := range nums {
		if v < 0 {
			v = -v
		}
		if nums[v-1] > 0 {
			nums[v-1] = -nums[v-1]
		} else {
			res = append(res, v)
		}
	}

	return res
}