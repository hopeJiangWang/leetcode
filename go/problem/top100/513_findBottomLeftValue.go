package top100

func findBottomLeftValue(root *TreeNode) int {
    var res int

	if root == nil {
		return res
	}

	var queue []*TreeNode
	queue = append(queue, root)

	for len(queue) > 0 {
		// 获取当前层级数据
		curLevel := queue
		queue = []*TreeNode{}
		res = curLevel[0].Val	// 记录当前层级的第一个数据
		// 遍历当前层级
		for len(curLevel) > 0 {
			// 获取下一层级数据
			top := curLevel[0]
			if top.Left != nil {
				queue = append(queue, top.Left)
			}
			if top.Right != nil {
				queue = append(queue, top.Right)
			}
			curLevel = curLevel[1:]
		}
	}

	return res
}