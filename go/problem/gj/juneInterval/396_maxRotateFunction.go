package juneinterval

/*
输入: nums = [4,3,2,6]
输出: 26
解释:
F(0) = (0 * 4) + (1 * 3) + (2 * 2) + (3 * 6) = 0 + 3 + 4 + 18 = 25
F(1) = (0 * 6) + (1 * 4) + (2 * 3) + (3 * 2) = 0 + 4 + 6 + 6 = 16
F(2) = (0 * 2) + (1 * 6) + (2 * 4) + (3 * 3) = 0 + 6 + 8 + 9 = 23
F(3) = (0 * 3) + (1 * 2) + (2 * 6) + (3 * 4) = 0 + 2 + 12 + 12 = 26
所以 F(0), F(1), F(2), F(3) 中的最大值是 F(3) = 26 。
*/

func MaxRotateFunction(nums []int) int {
	/*
	F1-F0 = a0 + a1 + a2 - 3 * a3
	F2-F1 = a3 + a0 + a1 - 3 * a2
	F3-F2 = a3 + a2 + a0 - 3 * a1
	
	sum = a0 + a1 + a2 + a3
	==> 
		F1-F0 = sum - 4 * a3
		F2-F1 = sum - 4 * a2
		F3-F2 = sum - 4 * a1
		...
		Fn-Fn-1 = sum - n * a0
		==> Fi+1 - Fi = sum - n * a[n-1-i]
		
	*/
	sum, f0 := 0, 0
	len1 := len(nums)
	for i := 0; i < len1; i++ {
		sum += nums[i]
		f0 += i * nums[i]
	}

	tmpSum, res := f0, f0
	for i := 0; i < len1; i++ {
		tmpSum += sum - len1 * nums[len1-1-i]
		res = max(tmpSum, res)
	}
	return res
}

func max(x, y int) int {
	if x < y {
		return y
	}
	return x
}






