package everyday

/*
输入：s = "IDID"
输出：[0,4,1,3,2]

输入：s = "III"
输出：[0,1,2,3]

输入：s = "DDI"
输出：[3,2,0,1]
*/

func diStringMatch(s string) []int {
	/*
		简单贪心：
		如果当前s[i]=I, 则p[i]<p[i+1]，让p[i]为此时最小的n
		如果当前s[i]=D, 则p[i]>p[i+1]，让p[i]为此时最大的n
	*/
	len1 := len(s)
	low, high := 0, len1
	var res []int
	for i := 0; i < len1; i++ {
		if s[i] == 'I' {
			res = append(res, low)
			low += 1
		} else {
			res = append(res, high)
			high -= 1
		}
	}
	// 注意最后会剩下一个数
	res = append(res, low)
	return res
}