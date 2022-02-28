package top100

import (
	// "fmt"
)

type TreeNode struct {
	Val int
	Left *TreeNode
	Right *TreeNode
}

func LevelOrder(root *TreeNode) [][]int {
	/*
		二叉树的层序遍历：
		从根节点开始，依次保存下一层所有节点，遍历之

	*/
	var res [][]int
	if root == nil {
		return res
	}

	var queue []*TreeNode
	queue = append(queue, root)

	for len(queue) > 0 {
		// 获取当前层级
		curLevel := queue
		tmpArr := []int{}
		queue = []*TreeNode{}
		// 遍历当前层级数据
		for len(curLevel) > 0 {
			top := curLevel[0]
			tmpArr = append(tmpArr, top.Val)
			// 获取下一层的数据
			if top.Left != nil {
				queue = append(queue, top.Left)
			}
			if top.Right != nil {
				queue = append(queue, top.Right)
			}
			curLevel = curLevel[1:]
		}
		res = append(res, tmpArr)
	}

	return res
}