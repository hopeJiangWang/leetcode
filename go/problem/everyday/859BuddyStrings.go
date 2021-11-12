package problem

import (
    "fmt"
)

func BuddyStrings(s string, goal string) bool {
    // 1. 长度不一致，直接错误
    if len(s) != len(goal) {
        return false
    } else {
        if s == goal {
            // 2. 长度一致，如果两个字符串完全相同，则判断有无两个相同的字符，有则通过
            var hash map[string]int 
            hash = make(map[string]int)
            for _, ch := range s {
                hash[string(ch)] ++
                fmt.Println(hash[string(ch)])
                if hash[string(ch)] > 1 {
                    return true
                } 
            }
            return false
        } else {
            // 3. 长度一致，字符不完全相同，则有且只有两个字符位置不一致，并且交换之后就要一致
            first := -1
            second := -1
            for i := 0; i < len(s); i++ {
                if s[i] != goal[i] {
                    if first < 0 {
                        first = i
                    } else if second < 0 {
                        second = i
                    } else {
                        return false
                    }
                }
            }
            // 此处易遗漏非负判断（仅需判断second即可）
            return second != -1 && s[first] == goal[second] && s[second] == goal[first]
            // return second != -1 && first != -1 && s[first] == goal[second] && s[second] == goal[first]
        }
    }

    // return false
}
