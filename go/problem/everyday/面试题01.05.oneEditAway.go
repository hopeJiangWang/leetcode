package everyday

import "fmt"

/*
字符串有三种编辑操作:插入一个字符、删除一个字符或者替换一个字符。
给定两个字符串，编写一个函数判定它们是否只需要一次(或者零次)编辑。

输入:
first = "pale"
second = "ple"
输出: True

输入:
first = "pales"
second = "pal"
输出: False
*/

func OneEditAway(first string, second string) bool {
	/*
		依次从头至尾遍历，
		1、因为他们最多只能相差一个字符，做好计数，超过一个时，直接返回失败
		2、有一个到了末尾，那么退出遍历，校验剩余的那个是否只剩一个或者也到了末尾
	*/

	pf, ps, cnt := 0, 0, 0
	len1, len2 := len(first), len(second)

	if len1 == len2 && len1 == 0 {
		return true
	}

	for pf < len1 && ps < len2 {
		if cnt > 1 {
			return false
		}
		
		if first[pf] == second[ps] {
			fmt.Printf("pf: %s, ps: %s \n", string(first[pf]), string(second[ps]))
		} else if pf+1 < len1 && first[pf+1] == second[ps] {
			fmt.Printf("pf+1: %s, ps: %s \n", string(first[pf+1]), string(second[ps]))
			cnt++
			pf++
		} else if ps+1 < len2 && first[pf] == second[ps+1] {
			fmt.Printf("pf: %s, ps+1: %s \n", string(first[pf]), string(second[ps+1]))
			cnt++
			ps++
		} else {
			cnt++
		}
		pf++
		ps++
	}

	fmt.Printf("pf: %d, ps: %d, cnt: %d \n", pf, ps, cnt)
	if pf == len1-1 && cnt == 0 {
		return true
	} else if ps == len2-1 && cnt == 0 {
		return true
	} else if pf == len1 && ps == len2 && cnt <= 1 {
		return true
	}
	return false
}