package top100

import "fmt"

// O(n^2)
func WordBreak(s string, wordDict []string) bool {
	// 动规：转换为s[0,i]是否可以拆分为wordDict中的单词，本题需要求s[0,s.lenghth()-1]
	// 取中间位置为j，也即：s[0,i]可以拆分，转换为s[0,j]可以拆分，且s[j,i]在wordDict中

	dictHash := make(map[string]bool)
	for _, v := range wordDict {
		dictHash[v] = true
	}

	lenOfS := len(s)
	dp := make([]bool, lenOfS+1)
	dp[0] = true

	for i := 1; i <= lenOfS; i++ {
		for j := 0; j < i; j++ {
			tmpStr := s[j:i]
			fmt.Println("tmpStr: ", tmpStr)
			if dp[j] && dictHash[tmpStr] {
				dp[i] = true
				break
			}
		}
	}

	return dp[lenOfS]
}
