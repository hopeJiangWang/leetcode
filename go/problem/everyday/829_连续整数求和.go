package everyday

/*
输入: n = 15
输出: 4
解释: 15 = 8 + 7 = 4 + 5 + 6 = 1 + 2 + 3 + 4 + 5
*/

func consecutiveNumbersSum(n int) int {
	/*
		因为连续，所以直接等差数列求和：首项a，项数k，公差1
		sum(a,a+k-1) = (a+a+k-1)*k/2 = N
		==>(2a+k-1)*k = 2n(1)
		==>2a=2n/k - k + 1 >= 2(3)(a,k均为正整数)
		==>2n/k >= k+1
		==>2n/k > k
		==>k < sqrt(2n)(2)：确认区间

		k 是2n的约数
	*/

	ans := 0

	for k := 1; k*k < 2*n; k++ {
		if 2*n % k == 0 && (2*n/k-k+1)%2 == 0 {
			ans++
		}
	}
	return ans
}