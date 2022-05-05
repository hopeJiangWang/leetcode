package mayDP

func maxUncrossedLines(nums1 []int, nums2 []int) int {
    /*
        可以看作是找两个字符串的最长公共子序列，同1143题
    */
    len1, len2 := len(nums1), len(nums2)
    dp := make([][]int, len1+1)
    for i :=range dp {
        dp[i] = make([]int, len2+1)
    }
    
    dp[0][0] = 0
    for i := 1; i <= len1; i++ {
        for j := 1; j <= len2; j++ {
            if nums1[i-1] == nums2[j-1] {
                dp[i][j] = dp[i-1][j-1] + 1
            } else {
                dp[i][j] = max(dp[i-1][j], dp[i][j-1])
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