package everyday

import (
	// "fmt"
	// "mycode/leetcode/go/problem/common"
)

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

type TreeNode struct {
	Val int
	Left *TreeNode
    Right *TreeNode
}

func MyAbsInt(tm int) int {
	y := tm >> 31
	return (tm ^ y) - y
}

// func dfs(root *TreeNode) (int, map[*TreeNode]int) {
// 	var res map[*TreeNode]int
// 	if root == nil {
// 		return 0, res
// 	}

// 	sumLeft, _ := dfs(root.Left)
// 	sumRight, _ := dfs(root.Right)

// 	res[root] = MyAbsInt(sumLeft - sumRight)

// 	return MyAbsInt(sumLeft - sumRight) + root.Val, res
// }

func sum (root *TreeNode ) int {
	if root == nil {
		return 0
	} 

	return root.Val + sum(root.Left) + sum(root.Right)
}

func FindTilt(root *TreeNode) int {
	// 空树直接返回
	if root == nil {
		return 0
	}
	
	// 结果为：根节点的坡度+其他节点的坡度
	return MyAbsInt(sum(root.Left) - sum(root.Right)) + FindTilt(root.Left) + FindTilt(root.Right);
}


func FindTilt2(root *TreeNode) (ans int) {
	// 使用dfs
	var dfs func(*TreeNode) int
	dfs = func (node *TreeNode) int {
		if node == nil {
			return 0
		}
		left := dfs(node.Left);
		right := dfs(node.Right);
		ans += MyAbsInt(left - right);
		return left + right + node.Val;
	}
	dfs(root)
	return ans
}

// func dfs(root *TreeNode) (ans int) {
// 	if root == nil {
// 		return 0
// 	}
// 	left := dfs(root.Left);
// 	right := dfs(root.Right);
// 	ans += MyAbsInt(left - right);
// 	return left + right + root.Val;
// }