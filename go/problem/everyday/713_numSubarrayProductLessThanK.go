package everyday

/*
输入：nums = [10,5,2,6], k = 100
输出：8
解释：8 个乘积小于 100 的子数组分别为：[10]、[5]、[2],、[6]、[10,5]、[5,2]、[2,6]、[5,2,6]。
需要注意的是 [10,5,2] 并不是乘积小于 100 的子数组。

输入：nums = [1,2,3], k = 0
输出：0

*/

func numSubarrayProductLessThanK(nums []int, k int) int {
	/*
		1、如果一个子串的乘积小于k，那么他的每个子集都小于k；
		2、这部分的子串数量刚好是r - l + 1
	*/

	if k == 0 || k == 1 {
		return 0
	}

	l, tmp, res := 0, 1, 0
	for r := 0; r < len(nums); r++ {
		tmp *= nums[r]
		for tmp >= k {
			tmp /= nums[l]
			l++
		}
		res += r - l + 1
	}
	return res
}