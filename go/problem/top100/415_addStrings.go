package top100

import (
	"fmt"
	"strconv"
)

/*
输入：num1 = "456", num2 = "77"
输出："533"

*/

func AddStrings(num1, num2 string) string {
	var res string = ""

	len1, len2, l1, l2, carry := len(num1), len(num2), 0, 0, 0
	for l1 < len1 || l2 < len2 || carry > 0 {
		tmpNum1, tmpNum2 := 0, 0
		if l1 < len1 {
			tmpNum1 = int(num1[len1-l1-1]-'0')
		}

		if l2 < len2 {
			tmpNum2 = int(num2[len2-l2-1]-'0') 
		}
		now := tmpNum1 + tmpNum2 + carry 
		carry = now / 10
		res = strconv.Itoa(now % 10) + res
		fmt.Printf("now: %d, carry: %d, res: %s \n", now, carry, res)
		
		l1++
		l2++
	}

	fmt.Println(res)
	return res
}

func reverseString(s string) string {
	a := []rune(s)
	lenS := len(s)

	for i, j := 0, lenS-1; i < j; i, j = i+1, j-1 {
		a[i], a[j] = a[j], a[i]
	}

	return string(a)
}