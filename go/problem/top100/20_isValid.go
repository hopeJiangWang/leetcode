package top100

import (
	"fmt"
	"container/list"
)

/*
有效字符串需满足：

左括号必须用相同类型的右括号闭合。
左括号必须以正确的顺序闭合。

*/

func IsValid(s string) bool {
	stack := list.New()

	for _, v := range s {
		backString := stack.Back()
		
		if stack.Len() > 0 && isMatch(backString.Value.(string), string(v)) {
			stack.Remove(backString)
			fmt.Printf("now: %v, backString: %v \n", string(v), backString.Value.(string))
		} else {
			stack.PushBack(string(v))
			fmt.Printf("now: %v \n", string(v))
		}
	}

	if stack.Len() != 0 {
		return false
	}

	return true
}

func isMatch(s1, s2 string) bool {
	if (s1 == "(" && s2 == ")") || (s1 == "[" && s2 == "]") || (s1 == "{" && s2 == "}") {
		return true
	}
	return false
}
