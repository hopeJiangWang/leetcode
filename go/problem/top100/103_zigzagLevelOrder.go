package top100

func ZigzagLevelOrder(root *TreeNode) [][]int {
    /*
        1.还是层次遍历的变形而已，奇数从左至右，偶数从右至左；
        2.使用一个队列存储每一层的数据，然后依次按奇偶便利；
    */

    var res [][]int

    if root == nil {
        return res
    }

    var queue []*TreeNode
    queue = append(queue, root)
    cnt := 1    // 标识层数的奇偶
    for len(queue) > 0 {
        nowLevel := queue
        tmp := []int{}
		queue = []*TreeNode{}
        if cnt % 2 != 0 {
            for i := 0; i < len(nowLevel); i++ {
                tmp = append(tmp, nowLevel[i].Val)
            }
        } else {
            for i := len(nowLevel)-1; i >= 0; i-- {
                tmp = append(tmp, nowLevel[i].Val)
            }
        }

        res = append(res, tmp)
        for i := 0; i < len(nowLevel); i++ {
            if nowLevel[i].Left != nil {
                queue = append(queue, nowLevel[i].Left)
            }

            if nowLevel[i].Right != nil {
                queue = append(queue, nowLevel[i].Right)
            }
        }
        cnt++
    }
    return res
}