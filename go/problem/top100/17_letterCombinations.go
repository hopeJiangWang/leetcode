package top100

import (
	"fmt"
)

/*
输入：digits = "23"
输出：["ad","ae","af","bd","be","bf","cd","ce","cf"]

输入：digits = ""
输出：[]

输入：digits = "2"
输出：["a","b","c"]
*/

var phoneMap map[string]string = map[string]string{
    "2": "abc",
    "3": "def",
    "4": "ghi",
    "5": "jkl",
    "6": "mno",
    "7": "pqrs",
    "8": "tuv",
    "9": "wxyz",
}

func LetterCombinations(digits string) []string {
	res := new([]string)

	// 如果字符串为空，直接返回空值
	if len(digits) == 0 {
		return *res
	}

	LetterBacktracking(0, res, digits, "")
	return *res
}


func LetterBacktracking(h int, res *[]string, digits, tmpStr string) {
	/*
		回溯跳出的条件：按键使用完毕
	*/
	fmt.Printf("tmpStr: %v, res: %v \n", tmpStr, res)
	if h == len(digits) {
		*res = append(*res, tmpStr)
	} else {
		// 依次选择本层数据，选好了就进入下一层递归
		nowStr := phoneMap[string(digits[h])]
		for i := 0; i < len(nowStr); i++ {
			/*
			tmpStr += string(nowStr[i])
			// fmt.Printf("tmpStr: %v, nowStr: %v, h: %v \n", tmpStr, nowStr, h)
			LetterBacktracking(h+1, res, digits, tmpStr)
			// 返回上一层，下一个数据
			tmpStr = tmpStr[0:len(tmpStr)-1]
			*/
			LetterBacktracking(h+1, res, digits, tmpStr + string(nowStr[i]))
		}
	}
}