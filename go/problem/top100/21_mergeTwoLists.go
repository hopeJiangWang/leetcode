package top100

func mergeTwoLists(list1 *ListNode, list2 *ListNode) *ListNode {
	res := &ListNode{Val: -1}	// 必须非空，否则后续next会空
	cur := res 

	for list1 != nil && list2 != nil {
		if list1.Val >= list2.Val {
			cur.Next = list2
			list2 = list2.Next
		} else if list1.Val < list2.Val {
			cur.Next = list1
			list1 = list1.Next
		}
		cur = cur.Next
	}

	if list1 != nil {
		cur.Next = list1
	} else if list2 != nil {
		cur.Next = list2
	}
	return res.Next
}