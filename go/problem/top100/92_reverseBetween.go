package top100

/*
left, right 位置之间进行翻转：
	输入：head = [1,2,3,4,5], left = 2, right = 4
	输出：[1,4,3,2,5]
*/

func ReverseBetween(head *ListNode, left int, right int) *ListNode {
	res := &ListNode{Val: -1}
	res.Next = head

	pre := res
	for i := 1; i < left; i++ {
		pre = pre.Next
	}

	/*
		1.定位到要反转部分的头节点 2，head = 2；前驱结点 1，pre = 1；
		2.当前节点的下一个节点3调整为前驱节点的下一个节点 1->3->2->4->5,
		3.当前结点仍为2， 前驱结点依然是1，重复上一步操作。。。
		4.1->4->3->2->5.

		next := cur.Next    // 先获取下一个节点
        cur.Next = pre      // 反转当前节点的next指针
        pre = cur           // 更新前置节点
        cur = next          // 跳到下一个节点
	*/
	cur := pre.Next
	for i := left; i < right; i++ {
		nxt := cur.Next		// 先获取下一个节点
		cur.Next = nxt.Next	// 反转部分的尾部指向下下个节点
		nxt.Next = pre.Next	// 反转当前节点的next指针
		pre.Next = nxt		// 将 下一节点 放到反转部分的头部
	}

	return res.Next
}
