package mayDP

import "sort"

/*
输入：envelopes = [[5,4],[6,4],[6,7],[2,3]]
输出：3
解释：最多信封的个数为 3, 组合为: [2,3] => [5,4] => [6,7]。

*/

func MaxEnvelopes(envelopes [][]int) int {
	/*
		dp[i]表示以第i个信封结尾的最多信封数目
		dp[i] = max(dp[i], dp[j]+1)(0≤j<i且envelopes_size[j]<envelopes_size[i])

		res = max(dp[i])

		设当前已求出的最多信封序列的长度为len（初始时为 1），从前往后遍历数组envelopes，在遍历到envelopes[i] 时：
		(1)如果envelopes_size[i]>d[len] ，则直接加入到 d 数组末尾，并更新 len=len+1；
		(2)否则，在 d 数组中二分查找，找到第一个比 envelopes_size[i] 小的数 d[k]，并更新d[k+1]=envelopes_size[i]。
	
	*/
	sort.Slice(envelopes, func(i, j int) bool {
        a, b := envelopes[i], envelopes[j]
        return a[0] < b[0] || a[0] == b[0] && a[1] > b[1]
    })


	f := []int{}
    for _, e := range envelopes {
        h := e[1]
		// 找到第一个比当前h 小的数，更新f[i]
        if i := sort.SearchInts(f, h); i < len(f) {
            f[i] = h
        } else {
			// 否则，直接加入到f 末尾
            f = append(f, h)
        }
    }
    return len(f)


	// len1, lend := len(envelopes), 1
	// if len1 == 0 {
	// 	return 0
	// }

	
	// d := make([][2]int, len1+1)
	// d[lend] = [2]int{envelopes[0][0], envelopes[0][1]}

	// for i := 0; i < len1; i++ {
	// 	if d[lend][0] < envelopes[i][0] && d[lend][1] < envelopes[i][1] {
	// 		lend++
	// 		d[lend] = [2]int{envelopes[i][0], envelopes[i][1]}
	// 	} else {
	// 		l, r, k := 1, lend, 0
	// 		for l <= r {
	// 			mid := (l + r) >> 1
	// 			if d[mid][0] < envelopes[i][0] && d[mid][1] < envelopes[i][1] {
	// 				k = mid
	// 				l = mid + 1
	// 			} else {
	// 				r = mid - 1
	// 			}
	// 		}
	// 		d[k+1] = [2]int{envelopes[i][0], envelopes[i][1]}
	// 	}
	// }

	// return lend
}

func MaxEnvelopes2(envelopes [][]int) int {
	/*
		dp[i]表示以第i个信封结尾的最多信封数目
		dp[i] = max(dp[i], dp[j]+1)(0≤j<i且envelopes_size[j]<envelopes_size[i])

		res = max(dp[i])

		设当前已求出的最多信封序列的长度为len（初始时为 1），从前往后遍历数组envelopes，在遍历到envelopes[i] 时：
		(1)如果envelopes_size[i]>d[len] ，则直接加入到 d 数组末尾，并更新 len=len+1；
		(2)否则，在 d 数组中二分查找，找到第一个比 nums[i] 小的数 d[k]，并更新d[k+1]=nums[i]。
	
		以输入序列 [0, 8, 4, 12, 2] 为例：
			第一步插入 0，d = [0]；
			第二步插入 8，d = [0, 8]；
			第三步插入 4，d = [0, 4]；
			第四步插入 12，d = [0, 4, 12]；
			第五步插入 2，d = [0, 2, 12].
	*/
	sort.Slice(envelopes, func(i, j int) bool {
		if envelopes[i][0] != envelopes[j][0] {
			return envelopes[i][0] < envelopes[j][0]
		} 
		return envelopes[i][1] < envelopes[j][1]
	})

	len1, res := len(envelopes), 1
	dp := make([]int, len1)

	for i := 0; i < len1; i++ {
		dp[i] = 1
		for j := 0; j < i; j++ {
			if envelopes[j][0] < envelopes[i][0] && envelopes[j][1] < envelopes[i][1] {
				dp[i] = max(dp[i], dp[j]+1)
			}
		}
		res = max(res, dp[i])
	}
	return res
}