package top100

import (
	"fmt"
	"math"
)

/*
输入：x = 0
输出：0

输入：x = -123
输出：-321

输入：x = 123
输出：321
*/

func Reverse(x int) int {
	var res int

	/*
		先判断首位：
		1. -：记录下来，翻转完后拼接上`去
		2. 0：直接返回0即可

		翻转：
		1. 从尾部开始遍历，res = res * 10 + nowValue
	*/
	var negativeFlag int = 1
	maxInt := math.MaxInt32
	if x < 0 {
		negativeFlag = -1
		x = -1 * x
	} else if x == 0 {
		return 0
	}

	for x > 0 {
		// 超过32位整数范围就返回0
		if res*10 > maxInt {
			return 0
		}
		var nowValue = x % 10
		res = res*10 + nowValue
		x = (x - (x % 10)) / 10

		fmt.Printf("nowValue: %v, res: %v, x: %v\n", nowValue, res, x)
	}

	return res * negativeFlag
}
