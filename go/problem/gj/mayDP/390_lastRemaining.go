package mayDP

/*
输入：n = 9
输出：6
解释：
arr = [1, 2, 3, 4, 5, 6, 7, 8, 9]
arr = [2, 4, 6, 8]
arr = [2, 6]
arr = [6]
           
a[n]：表示从左至右删除一遍之后剩余的数
b[n]：表示从右至左删除一遍之后剩余的数

a[n]：1, 2, 3, 4, 5, 6 ==> 2, 4, 6
b[n]：6, 5, 4, 3, 2, 1 ==> 5, 3, 1
==> a[n]+b[n]=n+1

b[n/2]: 3, 2, 1 ==> a[n] = 2*b[n/2]
==> a[n/2]+b[n/2]=n/2+1 ==> a[n]=2*(n/2+1-a[n/2])，其中a[1]=1


*/

func LastRemaining(n int) int {
	if n == 1 {
		return 1
	}

	return 2*(n/2 + 1 - LastRemaining(n/2))
}