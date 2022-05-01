package everyday

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
 func getAllElements(root1 *TreeNode, root2 *TreeNode) []int {
    /*
        中序遍历之后，得到两个有序数组，再合并
    */
    var res []int
    var nums1 []int
    var nums2 []int
    
    innerOrder(root1, &nums1)
    innerOrder(root2, &nums2)
    
    len1, len2, i, j := len(nums1), len(nums2), 0, 0
    
    for i < len1 && j < len2 {
        if nums1[i] < nums2[j] {
            res = append(res, nums1[i])
            i++
        } else {
            res = append(res, nums2[j])
            j++
        }
    }
    
    if i < len1 {
        res = append(res, nums1[i:]...)
    } else if j < len2 {
        res = append(res, nums2[j:]...)
    }
    return res
}

func innerOrder(root * TreeNode, res *[]int){
    if root == nil {
        return
    }
    
    innerOrder(root.Left, res)
    *res = append(*res, root.Val)
    innerOrder(root.Right, res)
}