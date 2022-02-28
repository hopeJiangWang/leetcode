package top100

import "math"

func maxPathSum(root *TreeNode) int {
	/*
		最大路径和：
		1.左+根+右
		2.左+根+根的父节点
		3.右+根+根的父节点
		*/
	res := math.MinInt32

	dfs(root, &res)

	return res
}

func dfs(root *TreeNode, res *int) int {
	if root == nil {
		return 0
	}

	leftSum := max(0, dfs(root.Left, res))   // 左子树的最大路径和
	rightSum := max(0, dfs(root.Right, res)) // 右子树的最大路径和
	innerSum := leftSum + rightSum + root.Val    // 当前子树的内部路径和

	*res = max(innerSum, *res)  // 看下是否能更新总的最大和
	// fmt.Println("cur: ", root.Val, "res: ", res)
	return root.Val + max(leftSum, rightSum) // 当前子树对外提供的最大路径和
}

