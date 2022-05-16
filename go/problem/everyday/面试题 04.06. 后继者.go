package everyday

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
 func inorderSuccessor(root *TreeNode, p *TreeNode) *TreeNode {
    var res []*TreeNode
    innerOrder2(root, &res)

    flag := false
    for _, v := range res {
        if v.Val == p.Val {
            flag = true
            continue
        }

        if flag {
            return v
        }
    }
    return nil
}

func innerOrder2(root *TreeNode, res *[]*TreeNode){
    if root == nil {
        return
    }
    
    innerOrder2(root.Left, res)
    *res = append(*res, root)
    innerOrder2(root.Right, res)
}