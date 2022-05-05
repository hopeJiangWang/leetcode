package mayDP

func longestCommonSubsequence(text1 string, text2 string) int {
	/*
		dp[i][j]：表示长度为[0,i-1]的字符串text1与长度为[0,j-1]的字符串text2的最长公共子序列的长度；
		1、text1[i-1] == text2[j-1]：那么找到了一个共同元素，dp[i][j] = dp[i-1][j-1] + 1
		2、text1[i-1] != text2[j-1]：那么这时候就要找之前的最大值了，dp[i][j] = max(dp[i][j-1], dp[i-1][j])
	*/
	len1, len2 := len(text1), len(text2)
	dp := make([][]int, len1+1)
	for i := range dp {
		dp[i] = make([]int, len2+1)
	}

	dp[0][0] = 0
	for i := 1; i <= len1; i++ {
		for j := 1; j <= len2; j++ {
			if text1[i-1] == text2[j-1] {
				dp[i][j] = dp[i-1][j-1] + 1
			} else {
				dp[i][j] = max(dp[i][j-1], dp[i-1][j])
			}
		}
	}
	return dp[len1][len2]
}

// func max(x, y int) int {
// 	if x < y {
// 		return y
// 	}
// 	return x
// }