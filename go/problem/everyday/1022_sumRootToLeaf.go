package everyday

/*
输入：root = [1,0,1,0,1,0,1]
输出：22
解释：(100) + (101) + (110) + (111) = 4 + 5 + 6 + 7 = 22

*/

func SumRootToLeaf(root *TreeNode) int {
	return helper(root, 0)
}

func helper(root *TreeNode, sum int) int {
	if root == nil {
		return 0
	}

	sum = sum * 2 + root.Val
	if root.Left == nil && root.Right == nil {
		return sum
	}

	return helper(root.Left, sum) + helper(root.Right, sum)
}
