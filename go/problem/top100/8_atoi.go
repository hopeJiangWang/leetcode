package main

import (
	"fmt"
	"math"
)

/*
1、读入字符串并丢弃无用的前导空格
2、检查下一个字符（假设还未到字符末尾）为正还是负号，读取该字符（如果有）。 确定最终结果是负数还是正数。 如果两者都不存在，则假定结果为正。
3、读入下一个字符，直到到达下一个非数字字符或到达输入的结尾。字符串的其余部分将被忽略。
4、将前面步骤读入的这些数字转换为整数（即，"123" -> 123， "0032" -> 32）。如果没有读入数字，则整数为 0 。必要时更改符号（从步骤 2 开始）。
5、如果整数数超过 32 位有符号整数范围 [−231,  231 − 1] ，需要截断这个整数，使其保持在这个范围内。具体来说，小于 −231 的整数应该被固定为 −231 ，大于 231 − 1 的整数应该被固定为 231 − 1 。
6、返回整数作为最终结果。

输入：s = "4193 with words"
输出：4193

输入：s = "   -42"
输出：-42

输入：s = "42"
输出：42
*/

func myAtoi(s string) int {
	res, flag, i := 0, 1, 0
	lenOfS := len(s)

	for i < lenOfS && s[i] == ' ' {
		i++
	}
	
	for i < lenOfS {
		if s[i] == '-' {
			flag = -1
		} else if s[i] == '+' {
			flag = 1
		}
	}

	for i < lenOfS && s[i] >= '0' && s[i] <= '9' {
		res = res * 10 + int(s[i] - '0')
		if res * flag >= math.MaxInt32 {
			return math.MaxInt32
		} else if res * flag <= math.MinInt32 {
			return math.MinInt32
		}
	}

	return res * flag
}

func main() {
	fmt.Println(myAtoi("   -42"))
}