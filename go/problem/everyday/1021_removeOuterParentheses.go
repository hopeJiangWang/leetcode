package everyday

// import "fmt"

/*
输入：s = "(()())(())"
输出："()()()"
解释：
输入字符串为 "(()())(())"，原语化分解得到 "(()())" + "(())"，
删除每个部分中的最外层括号后得到 "()()" + "()" = "()()()"。
*/

func RemoveOuterParentheses1(s string) string {
	var stack []byte
	var res string

	if len(s) == 0 {
		return s
	}

	// 先将左括号入栈
	stack = append(stack, s[0])
	tmp := string(stack[0])
	for i := 1; i < len(s); {
		// 若栈不为空，则一直记录tmp字符即可
		for len(stack) != 0 {
			if string(stack[len(stack)-1]) == "(" && string(s[i]) == ")" {
				stack = stack[0:len(stack)-1]
			} else {
				stack = append(stack, s[i])
			}
			tmp += string(s[i])
			i++
		}
		// 栈为空，则去除最外层括号，得到原语
		res += tmp[1:len(tmp)-1]
		// 还有字符剩余，初始化栈及tmp，重复上述步骤
		if i < len(s) {
			tmp = string(s[i])
			stack = append(stack, s[i])
			i++
		}
	}

	return res
}

func RemoveOuterParentheses(s string) string {
	/*
		栈从空到下一次空的过程，则是扫描了一个原语的过程。
		一个原语中，首字符和尾字符应该舍去，其他字符需放入结果字符串中。
	*/
    // var ans, st []rune
    // for _, c := range s {
    //     if c == ')' {
    //         st = st[:len(st)-1]
    //     }
	// 	// 只要当前栈中字符数大于0，那么当前字符就需要保留
	// 	// 因为append是在后面，所以自然跳过了第一个字符
	// 	// 到尾字符，上面也截取了
    //     if len(st) > 0 {
    //         ans = append(ans, c)
    //     }
    //     if c == '(' {
    //         st = append(st, c)
    //     }
    // }
    // return string(ans)

	var res, stack []rune

	for _, c := range s {
		if c == ')' {
			stack = stack[:len(stack)-1]
		}

		if len(stack) > 0 {
			res = append(res, c)
		}

		if c == '(' {
			stack = append(stack, c)
		}
	}

	return string(res)
}
