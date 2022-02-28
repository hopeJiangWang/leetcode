package top100

func diameterOfBinaryTree(root *TreeNode) int {
    res := 0
	/*
		其实就是看每个节点的最大直径
		一个节点的最大直径可以这样看：
		1.这个节点就是结果的根节点：那么它的最大直径就是左子树的最大直径+右子树的最大直径；
			其中每一层的直径贡献是 1
		2.这个节点只是最终结果的一个子树，那么我们就要看它对外提供的最大直径了：
			也就是max(左子树的最大直径，右子树的最大直径)
	*/

    if root == nil {
        return 0
    }

	dfs543(root, &res)
	return res
}

func dfs543(root *TreeNode, res *int) int {
    // 叶子节点的对外提供的直径为0
	if root.Left == nil && root.Right == nil {
		return 0
	}

    leftSum, rightSum := 0, 0
    if root.Left != nil {
        leftSum = dfs543(root.Left, res) + 1   // 左子树的最大直径
    }

	if root.Right != nil {
        rightSum = dfs543(root.Right, res) + 1 // 右子树的最大直径
    }
	
	innerSum := leftSum + rightSum    // 当前子树的最大路径

	*res = max(innerSum, *res)  // 看下是否能更新总的最大和
	// fmt.Println("res: ", *res)
	return max(leftSum, rightSum) // 当前子树对外提供的最大直径
}