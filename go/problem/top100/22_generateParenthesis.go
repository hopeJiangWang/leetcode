package top100

import (
	"fmt"
)

/*
输入：n = 3
输出：["((()))","(()())","(())()","()(())","()()()"]

输入：n = 1
输出：["()"]

void backtracking(参数) {
    if (终止条件) {
        存放结果;
        return;
    }

    for (选择：本层集合中元素（树中节点孩子的数量就是集合的大小）) {
        处理节点;
        backtracking(路径，选择列表); // 递归
        回溯，撤销处理结果
    }
}

*/

func GenerateParenthesis(n int) []string {
	// 使用new方法开辟内存返回内存地址
	res := new([]string)

	ParenthesisBacktracking(n, n, res, "")

	return *res
}

func ParenthesisBacktracking(left, right int, res *[]string, tmpStr string) {
	/*
		回溯跳出的条件：右括号使用完毕
	*/
	fmt.Printf("tmpStr: %v, res: %v \n", tmpStr, res)
	if right == 0 {
		*res = append(*res, tmpStr)
		return 
	}
	
	// 有左括号，直接加进去
	if left > 0 {
		ParenthesisBacktracking(left - 1, right, res, tmpStr + "(")
	}

	// 需要存在左括号，才能加右括号
	if right > left {
		ParenthesisBacktracking(left, right - 1, res, tmpStr + ")")
	}
}