package top100

/*
输入：s = "(()"
输出：2
解释：最长有效括号子串是 "()"

输入：s = ")()())"
输出：4
解释：最长有效括号子串是 "()()"

输入：s = ""
输出：0
*/

func LongestValidParentheses1(s string) int {
	var res int = 0
	if len(s) == 0 || s == "" {
		return res
	}

	/*
		1.对于遇到的每个 (，我们将它的下标放入栈中
		2.对于遇到的每个 ) ，我们先弹出栈顶元素表示匹配了当前右括号：
		（1）如果栈为空，说明当前的右括号为没有被匹配的右括号，我们将其下标放入栈中来更新基准值；
		（2）如果栈不为空，当前右括号的下标减去栈顶元素即为当前阶段的最大长度
	*/
	lenOfS := len(s)
	var stack []int
	stack = append(stack, -1)
	for i := 0; i < lenOfS; i++ {
		// 对于遇到的每个 (，我们将它的下标放入栈中
		if s[i] == '(' {
			stack = append(stack, i)
		} else {
			// 对于遇到的每个 ) ，我们先弹出栈顶元素表示匹配了当前右括号
			stack = stack[:len(stack)-1]
			// 如果栈为空，说明当前的右括号为没有被匹配的右括号，我们将其下标放入栈中来更新基准值
			if len(stack) == 0 {
				stack = append(stack, i)
			} else {
				// 如果栈不为空，当前右括号的下标减去栈顶元素即为当前阶段的最大长度
				res = max(res, i-stack[len(stack)-1])
			}
		}
	}

	return res
}

func LongestValidParentheses2(s string) int {
	var res int = 0
	/*
		1.从左至右遍历，记录左右括号的个数：left，right
			如果left == right，那么这时候就记录长度；
			如果left < right，同时将计数归零；

		2.从右至左遍历，记录左右括号的个数：left，right
			如果left == right，那么这时候就记录长度；
			如果left > right，同时将计数归零；
	*/
	lenOfS := len(s)
	left, right := 0, 0
	for i := 0; i < lenOfS; i++ {
		if s[i] == '(' {
			left++
		} else {
			right++
		}

		if left == right {
			res = max(res, 2*left)
		} else if left < right {
			left, right = 0, 0
		}
	}

	left, right = 0, 0
	for i := lenOfS - 1; i >= 0; i-- {
		if s[i] == '(' {
			left++
		} else {
			right++
		}

		if left == right {
			res = max(res, 2*left)
		} else if left > right {
			left, right = 0, 0
		}
	}

	return res
}

func LongestValidParentheses(s string) int {
	/*
		dp[i]表示以下标i字符结尾的最长有效括号的长度，以'('结尾的字串的dp一定为0
		1. s[i]=')'&&s[i-1]='('，则dp[i]=dp[i-2]+2
		2. s[i]=')'&&s[i-1]=')'，则dp[i]=dp[i-1]+2+dp[i-dp[i-1]-2]
	*/
	dp := make([]int, len(s))
	var res int = 0
	for i := 1; i < len(s); i++ {
		if s[i] == ')' {
			if s[i-1] == '(' {
				if i >= 2 {
					dp[i] = dp[i-2] + 2
				} else {
					dp[i] = 2
				}
			}
		} else if i-dp[i-1] > 0 && s[i-dp[i-1]-1] == '(' {
			if i-dp[i-1] >= 2 {
				dp[i] = dp[i-1] + 2 + dp[i-dp[i-1]-2]
			} else {
				dp[i] = dp[i-1] + 2
			}
		}
		res = max(res, dp[i])
	}

	return res
}
