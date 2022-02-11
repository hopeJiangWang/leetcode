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
		从根节点开始，将每一层所有节点的所有子节点全部存起来，也即遍历当前层的时候把下一层的节点存起来
		然后依次遍历存起来的节点，并且存起来下一层的节点
	*/
	var nowList []*TreeNode
	var res [][]int

	// 如果根节点为空，直接返回即可
	if root == nil {
		return res
	}

	// 先将根节点压入队列
	nowList = append(nowList, root)
	// 只要还有节点，就需要新增一层，继续遍历
	for i := 0; len(nowList) > 0; i++ {
		// 每一层初始化一个列表，并且新建一个列表存储下一层的节点
		res = append(res, []int{})
        tmpAns := []*TreeNode{} 
		
		// 遍历本层所有节点，并且将每个节点对应的下一层节点保存
		for j := 0; j < len(nowList); j++ {
			res[i] = append(res[i], nowList[j].Val)

			if nowList[j].Left != nil {
				tmpAns = append(tmpAns, nowList[j].Left)
			}

			if nowList[j].Right != nil {
				tmpAns = append(tmpAns, nowList[j].Right)
			}
		}

		nowList = tmpAns
	}

	return res
}