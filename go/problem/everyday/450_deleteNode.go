package everyday

func deleteNode(root *TreeNode, key int) *TreeNode {
	/*
		区分以下几种情况：
		1、如果当前节点值大于key，那么就递归查找左子树；
		2、如果当前节点值小于key，那么就递归查找右子树；
		3、当前节点值等于当前节点：
		1）无左子，右子顶替其位置，删除该节点；
		2）无右子，左子顶替其位置，删除该节点；
		3）左右子节点都有，将左子树移动到右子树的最左边节点的左子树上，
		右子树顶替其位置
	*/
	if root == nil {
		return nil
	}

	if root.Val > key {
		root.Left = deleteNode(root.Left, key)
		return root
	} else if root.Val < key {
		root.Right = deleteNode(root.Right, key)
		return root
	} else {
		if root.Left == nil {
			return root.Right
		}
		if root.Right == nil {
			return root.Left
		}

		tmpNode := root.Right
		// 找到右子树的最左边节点
		for tmpNode.Left != nil {
			tmpNode = tmpNode.Left
		}
		tmpNode.Left = root.Left
		root = root.Right
	}
	return root
}