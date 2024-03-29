package top100


/*
abcabcbb: 3
bbbbb: 1
pwwkew: 3
: 0
*/

func LengthOfLongestSubstring2(s string) int {
    var res int

    /*
        1.使用一个哈希表记录当前子串的所含字母；
        2.如果发现已经有重复的了，就记下当前的最大长度，然后将左指针往右移动；
        3.重复上述步骤，知道左指针到达尾部；
    */

    lenOfString := len(s)
    if lenOfString == 0 {
        return 0
    }

    wordMap := make(map[byte]int)
    right := -1 // 还未移动
    for i := 0; i < lenOfString; i++ {
        if i != 0 {
            // 此时需要移动左指针了
            delete(wordMap, s[i-1])
        }

        // 未到结尾且无重复字符
        for right < lenOfString && wordMap[s[right+1]] == 0 {
            wordMap[s[right]] = 1
            right++
        }

        res = max(res, right-i+1)
    }
    return res
}


// func LengthOfLongestSubstring(s string) int {
// 	var res int = 0

// 	if len(s) == 0 {
// 		return res
// 	}

// 	right, left := 0, 0
// 	tmpMap := make(map[byte]int)

// 	for left < len(s) {
// 		_, ok := tmpMap[s[left]]
// 		if ok {
// 			if left - right > res {
// 				res = left - right
// 			}
// 			// tmpMap = make(map[byte]int)
// 			tmp := right
// 			for tmp > 0 {
// 				delete(tmpMap, s[tmp])
// 				tmp--
// 			}
// 			right++		// 这里应该从0位置开始逐步往后移动
// 			left = right
// 		}

// 		tmpMap[s[left]] = 1
// 		left++
// 		fmt.Printf("right: %d, left: %d, tmpMap: %v, res: %d \n", right, left, tmpMap, res)
// 	}

// 	// 如果均不相同
// 	if len(tmpMap) > res {
// 		res = len(tmpMap)
// 	}
// 	return res
// }

func LengthOfLongestSubstring(s string) int {
    var res int = 0
    
    return res
}

func max(x, y int) int {
    if x < y {
        return y
    }
    return x
}
