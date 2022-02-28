package top100

func invertTree(root *TreeNode) *TreeNode {
    /*
        直接交换左右子树即可
    */

    if root == nil {
        return nil
    }

    tmp := root.Right
    root.Right = root.Left
    root.Left = tmp
    // fmt.Println(root.Val)

    invertTree(root.Right)
    invertTree(root.Left)
	return root
}