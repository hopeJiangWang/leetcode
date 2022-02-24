package top100

import "sort"

/*
输入: strs = ["eat", "tea", "tan", "ate", "nat", "bat"]
输出: [["bat"],["nat","tan"],["ate","eat","tea"]]

输入: strs = [""]
输出: [[""]]

输入: strs = ["a"]
输出: [["a"]]
*/

func GroupAnagrams(strs []string) [][]string {
	var res [][]string

	/*
		1. 将每个单词出现的字母作为key，key相同的单词就放到一起；
		2. 遍历结束之后，也就全部处理完毕了
	*/
	if len(strs) == 0 {
		return res
	}

	hashMap := make(map[[26]int][]string)

	// O(nk)
	for _, v := range strs {
		cnt := [26]int{}
		for _, b := range v {
			cnt[b-'a']++
		}
		hashMap[cnt] = append(hashMap[cnt], v)
	}

	for _, v := range hashMap {
		res = append(res, v)
	}

	return res
}

func GroupAnagrams1(strs []string) [][]string {
	var res [][]string

	/*
		1. 将每个单词按字典序排序好作为key，key相同的单词就放到一起；
		2. 遍历结束之后，也就全部处理完毕了
	*/
	if len(strs) == 0 {
		return res
	}

	hashMap := make(map[string][]string)

	// O(nklogk)
	for _, v := range strs {
		// 获取每个单词的字典序key
		tmpRune := []rune(v)
		sort.Slice(tmpRune, func(i, j int) bool {
			return tmpRune[i] < tmpRune[j]
		})
		tmpKey := string(tmpRune)

		// 不存在，插入
		hashMap[tmpKey] = append(hashMap[tmpKey], v)
	}

	for _, v := range hashMap {
		res = append(res, v)
	}

	return res
}
