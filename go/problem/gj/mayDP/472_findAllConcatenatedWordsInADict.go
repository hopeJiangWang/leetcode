package mayDP

/*
输入：words = ["cat","cats","catsdogcats","dog","dogcatsdog","hippopotamuses","rat","ratcatdogcat"]
输出：["catsdogcats","dogcatsdog","ratcatdogcat"]
解释："catsdogcats" 由 "cats", "dog" 和 "cats" 组成; 
     "dogcatsdog" 由 "dog", "cats" 和 "dog" 组成; 
     "ratcatdogcat" 由 "rat", "cat", "dog" 和 "cat" 组成。

*/
var hashSet map[int]int
var p int = 131
var offset int = 128
func findAllConcatenatedWordsInADict(words []string) []string {
	/*
		dp[i]为s=words[i]的前i个字符（[1,i]），能够切分出的最大item（每个连接部分）个数，最少也能切成本身dp[i]>1
		假设dp[i]可以由dp[j]转移而来（j < i），那么dp[j]!=0 && 子串s[j+1,i]在words中

		直接遍历判断的话，会达到O{n^3}，超时
		通过字符串哈希优化：
		1、先对words进行遍历，存入hashset中，将判断子串是否在words中转换为常数复杂度；
		2、调整转移逻辑为：从dp[i]出发，枚举范围[i+1,n]，找到可有dp[i]所能更新的状态dp[j]，尝试用dp[i]来更新dp[j]
			dp[j] = max(dp[j], dp[i]+1)
		
		字符串哈希会产生哈希碰撞，增加一个offset
	
	*/
	hashSet = make(map[int]int)
	for _, v := range words {
		hash := 0

		for _, c := range v {
			hash = hash * p + int(c - 'a') + offset
		}
		hashSet[hash]++
	}

	var res []string
	for _, s := range words {
		if check(s) {
			res = append(res, s)
		}
	}
	return res
} 

func check(s string) bool {
	len1 := len(s)

	dp := make([]int, len1+1)
	for i := 0; i < len1+1; i++ {
		dp[i] = -1
	}

	dp[0] = 0
	for i := 0; i <= len1; i++ {
		if dp[i] == -1 {
			continue
		}

		cur := 0
		for j := i+1; j <= len1; j++ {
			cur = cur * p + int(s[j-1] - 'a') + offset
			if _, ok := hashSet[cur]; ok {
				dp[j] = max(dp[j], dp[i]+1)
			}
		}
		if dp[len1] > 1 {
			return true
		}
	}
	return false
}

func findAllConcatenatedWordsInADict2(words []string) []string{
	/*
	1、用一个Set记录所有的单词，对单词数组words进行遍历，要把自己移除掉，因为不能自己构成自己。
	2、判断是否能由目前Set(已不包含自己)中的元素构成当前word，如果可以，将word将入结果集，
	具体判断过程如下：
	（1）如果Set集合没有元素或者字符串就是空的，直接返回false；
	（2）dp[i]表示word从[0, i)之间的字符串是否可以分解成Set集合已经包含的较短的字符串，dp[n]就表示可以完全分解掉；
	（3）假设现在要计算word[0, i - 1]是否可以形成连接词，可以分解为两部分：word[0, j - 1]和word[j, i]，
		i最终可以取到n，这也就对应着DP数组的两部分：word[0, j - 1]对应dp[j], 
		word[j, i]需要判断下Set集合中是否包含word[j, i]，如果包含则dp[i] = true；
		返回dp[n]；
	3、最后再将当前单词word加入Set进行下一轮的判断。
	*/

	// 先生成hashMap
	var res []string
	wordSet := make(map[string]bool)
	for _, v := range words {
		wordSet[v] = true
	}

	// 再依次遍历，看每个单词是否能够构成
	for _, v := range words {
		// 需要先将当前单词移除，因为自己不能构成自己
		delete(wordSet, v)
		if checkCanBreak(v, wordSet) {
			res = append(res, v)
		}
		// 后面需要再加回来
		wordSet[v] = true
	}

	return res
}

func checkCanBreak(s string, wordSet map[string]bool) bool {
	len1 := len(s)
	if len1 == 0 || len(wordSet) == 0 {
		return false
	}

	dp := make([]bool, len1+1)
	dp[0] = true

	// 如果dp[j] 能够构成，那么我们只需要判断s[j:i]是否在字典中
	// 在的话，那么dp[i] 能够构成
	for i := 1; i <= len1; i++ {
		for j := 0; j < i; j++ {
			if !dp[j] {
				continue
			}

			if wordSet[s[j:i]] {
				dp[i] = true
				break
			}
		}
	}
	// 最后，判断下[0,len1)是否能够构成即可
	return dp[len1]
}