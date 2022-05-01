package mayDP

func maxTurbulenceSize(arr []int) int {
    /*
        其实只要按规则不断统计即可：
        1、如果arr[i] < arr[i-1], 这时候是往下走，即当前的结果应该是之前的up+1，并且需要更新up的值: down=up+1, up=1;
        2、同理，如果arr[i] > arr[i-1]，这时候是往上走，则：up=down+1, down=1;
        3、如果arr[i]==arr[i-1]，这时候都得更新：up=down=1;
    */
    var res int = 1
    lenOfArr := len(arr)
    
    up, down := 1, 1
    for i := 1; i < lenOfArr; i++ {
        if arr[i] < arr[i-1] {
            down, up = up + 1, 1
        } else if arr[i] > arr[i-1] {
            up, down = down + 1, 1
        } else {
            up, down = 1, 1
        }
        res = max(res, max(up, down))
    }
    
    return res
}
    
func maxTurbulenceSize2(arr []int) int {
    /*
        dp方案：
        dp[i][0]: 代表arr[i]>arr[i-1]，即是往上走，那么dp[i][0]=dp[i-1][1]+1
        dp[i][1]: 代表arr[i]<arr[i-1]，即是往下走，那么dp[i][1]=dp[i-1][0]+1
        如果arr[i]==arr[i-1]，则dp[i][0]=dp[i][1]=1
        并不断统计记录最大值
    */
    len1 := len(arr)
    var res int = 1
    dp := make([][2]int, len1)
    dp[0] = [2]int{1, 1}
    for i := 1; i < len1; i++ {
        // 默认是需要重置的
        dp[i] = [2]int{1, 1}
        // 否则，增加对应的统计值
        if arr[i] > arr[i-1] {
            dp[i][0] = dp[i-1][1] + 1
        } else if arr[i] < arr[i-1] {
            dp[i][1] = dp[i-1][0] + 1
        } 
        res = max(res, max(dp[i][0], dp[i][1]))
    }
    
    return res
}
