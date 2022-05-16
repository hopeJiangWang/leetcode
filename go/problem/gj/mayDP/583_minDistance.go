package mayDP

/*
给定两个单词 word1 和 word2 ，返回使得 word1 和  word2 相同所需的最小步数。
每步 可以删除任意一个字符串中的一个字符。

输入: word1 = "sea", word2 = "eat"
输出: 2
解释: 第一步将 "sea" 变为 "ea" ，第二步将 "eat "变为 "ea"

输入：word1 = "leetcode", word2 = "etco"
输出：4
*/

func MinDistance(word1 string, word2 string) int {
	/*
		dp[i][j]：words1的前i-1个字符，与words2的前j-1个字符的最长相同子序列
		(1)如果words1[i]==words2[j]，dp[i][j]=dp[i-1][j-1]
		(2)如果words1[i]!=words2[j]，dp[i][j]=max(dp[i-1][j], dp[i][j-1])
	*/
	len1, len2 := len(word1), len(word2)
	dp := make([][]int, len1+1)
	for i := range dp {
		dp[i] = make([]int, len2+1)
	}

	dp[0][0] = 0
	for i := 1; i <= len1; i++ {
		for j := 1; j <= len2; j++ {
			if word1[i-1] == word2[j-1] {
				dp[i][j] = dp[i-1][j-1] + 1
			} else {
				dp[i][j] = max(dp[i][j-1], dp[i-1][j])
			}
		}
	}
	return len1 + len2 - dp[len1][len2]*2
}
