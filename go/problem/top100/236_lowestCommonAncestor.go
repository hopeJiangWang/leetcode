package top100

/*
	最近公共祖先：两个节点在一个节点的左右子树上，可以是节点本身， 也可以是在单独一个子树上。
*/

func lowestCommonAncestor(root, p, q *TreeNode) *TreeNode {
	/*
		如果节点等于目标值或者为空（到达底部），那么当前节点为根节点的子树的贡献就是目标值或者空，记录下
		（1）左右子树都按上面的处理，这时候，如果左右子树的返回分别是p，q，那么此时的根节点就是结果了；
		（2）如果左右子树其中一个为空，那么返回另外一个非空的即可；（p：5，q：6，他们的根节点为7，此时都在右子树上面，那么左子树的返回就是空，我们需要的就是右子树）
		（3）两个都是空，那么就是空；
	*/	
	if root == p || root == q || root == nil {
		return root
	}

	left := lowestCommonAncestor(root.Left, p, q)
	right := lowestCommonAncestor(root.Right, p, q)

	if !left && !right {
		return nil
	} else if !left && right {
		return right
	} else if left && !right {
		return left
	} 

	return root
}