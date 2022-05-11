package everyday

import (
	"fmt"
	"strconv"
	"strings"
	"math"
)

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

type Codec struct {
    
}

func Constructor() Codec {
    return Codec{}
}

// 先序：根-》左-》右
func preOrder(root *TreeNode, res *[]int) {
	if root == nil {
		return 
	}

	*res = append(*res, root.Val)
	preOrder(root.Left, res)
	preOrder(root.Right, res)
}

func backOrder(root *TreeNode, res *[]int) {
	if root == nil {
		return 
	}

	backOrder(root.Left, res)
	backOrder(root.Right, res)
	*res = append(*res, root.Val)
}


// Serializes a tree to a single string.
func (this *Codec) serialize(root *TreeNode) string {
	var resInt []int
    backOrder(root, &resInt)
	fmt.Println("resInt: ", resInt)

	var resString string
	len1 := len(resInt)
	for i := 0; i < len1; i++ {
		resString += strconv.Itoa(resInt[i])
		if i != len1-1 {
			resString += ","
		}
	}
	return resString
}

// func buildTree(l, r int, data []string) *TreeNode {
// 	if l > r {
// 		return nil
// 	}

// 	// 先将当前节点创建出来
// 	nowValue, _ := strconv.Atoi(data[l])
// 	nowNode := &TreeNode{Val: nowValue}

// 	// 找到第一个大于当前数的节点，那就是右子树的起点
// 	// （比较大小：小于当前数据，都应该是左子树；大于当前数据，都应该是右子树）
// 	j := l+1
// 	for j <= r {
// 		tmpValue, _ := strconv.Atoi(data[j])
// 		if tmpValue <= nowValue {
// 			j++
// 		}
// 		fmt.Println("j: ", j)
// 	}
// 	// 递归处理[l+1, j-1]构建左子树
// 	nowNode.Left = buildTree(l+1, j-1, data)
// 	// 递归处理[j, r]构建右子树
// 	nowNode.Right = buildTree(j, r, data)
// 	return nowNode
// }

// Deserializes your encoded data to tree.
func (this *Codec) deserialize(data string) *TreeNode {
	// dataSlice := strings.Split(data, ",")
    // return buildTree(0, len(data)-1, dataSlice)
	// 后序遍历得到的数组中，根结点的值位于数组末尾，左子树的节点均小于根节点的值，右子树的节点均大于根节点的值
	if len(data) == 0 {
		return nil
	}
	arr := strings.Split(data, ",")
    var construct func(int, int) *TreeNode
    construct = func(lower, upper int) *TreeNode {
        if len(arr) == 0 {
            return nil
        }
        val, _ := strconv.Atoi(arr[len(arr)-1])
        if val < lower || val > upper {
            return nil
        }
        arr = arr[:len(arr)-1]
        return &TreeNode{Val: val, Right: construct(val, upper), Left: construct(lower, val)}
    }
    return construct(math.MinInt32, math.MaxInt32)
}


/**
 * Your Codec object will be instantiated and called as such:
 * ser := Constructor()
 * deser := Constructor()
 * tree := ser.serialize(root)
 * ans := deser.deserialize(tree)
 * return ans
 */