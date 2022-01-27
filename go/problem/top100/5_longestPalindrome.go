package top100

import (
	"fmt"
)

/*
输入：s = "babad"
输出："bab"
解释："aba" 同样是符合题意的答案。

输入：s = "cbbd"
输出："bb"

输入：s = "a"
输出："a"

输入：s = "ac"
输出："a"
*/

func LongestPalindrome1(s string) string {
	/*
		使用dp[i][j]=1代表从 i 到 j 为回文，那么如果s[i-1]==s[j+1]，dp[i-1][j+1] = 1
		也即
			dp[i][j]=dp[i+1][j-1](s[i]==s[j])；
			dp[i][j]=0(s[i]!=s[j])
			dp[i][i]=1
	*/
	var res string = ""
	dp := make([][]bool, len(s))
	for i := 0; i < len(s); i++ {
		dp[i] = make([]bool, len(s))
	}

	for i := 0; i < len(s); i++ {
		for j := 0; j <= i; j++ {
			// 头尾相同，且中间也是回文
			if s[i] == s[j] && (i-1 < j+1 || dp[i-1][j+1]) {
				dp[i][j] = true
			}
			// 不断迭代最长的长度即可
			if dp[i][j] && i-j+1 > len(res) {
				res = s[j:i+1]
				fmt.Printf("res: %s", res)
			}
		}
	}
	
	return res
}

func LongestPalindrome(s string) string {
	/*
		中心扩展：字符串长度可以是奇数也可以是偶数

	*/
	n := len(s)
	if n < 2 {
		return s
	}
	start, end := 0, 0
	for i := 0; i < n; {
		l, r := i, i
		//如果字符串相同则分别从前一个和后一个开始回文
		for r < n-1 && s[r] == s[r+1] {
			r++
		}
		i = r + 1
		for l > 0 && r < n-1 && s[l-1] == s[r+1] {
			l--
			r++
		}
		if end < r-l {
			start = l
			end = r - l
		}
	}
	return s[start : start+end+1]
}
