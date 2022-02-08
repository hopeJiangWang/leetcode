package top100

import (
	"fmt"
)

type TreeNode struct {
     Val int
     Left *TreeNode
     Right *TreeNode
}

func LevelOrder(root *TreeNode) [][]int {
	tmpLevel := make(map[TreeNode]TreeNode{})
}