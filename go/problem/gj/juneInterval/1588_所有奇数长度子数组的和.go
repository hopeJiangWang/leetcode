package juneinterval

/*
输入：arr = [1,4,2,5,3]
输出：58
解释：所有奇数长度子数组和它们的和为：
[1] = 1
[4] = 4
[2] = 2
[5] = 5
[3] = 3
[1,4,2] = 7
[4,2,5] = 11
[2,5,3] = 10
[1,4,2,5,3] = 15
我们将所有值求和得到 1 + 4 + 2 + 5 + 3 + 7 + 11 + 10 + 15 = 58
*/

func SumOddLengthSubarrays(arr []int) int {
	len1 := len(arr)
	preSum := make([]int, len1+1)

	// 先获取所有前缀和
	for i := 0; i < len1; i++ {
		preSum[i+1] = preSum[i] + arr[i]
	}

	// 依次按奇数获取各个子数组的和
	res := 0
	for start := range arr {
        for length := 1; start+length <= len1; length += 2 {
            end := start + length - 1
            res += preSum[end+1] - preSum[start]
        }
    }
	return res
}