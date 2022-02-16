package top100

func MergeKLists(lists []*ListNode) *ListNode {
	return merge(lists, 0, len(lists)-1)
}

// 合并两个有序链表
func mergeTwoLists(l1, l2 *ListNode) *ListNode {
	if l1 == nil {
		return l2
	} else if l2 == nil {
		return l1
	}

	res := &ListNode{Val: -1}
	cur := res

	for l1 != nil && l2 != nil {
		if l1.Val < l2.Val {
			cur.Next = l1
			l1 = l1.Next
		} else {
			cur.Next = l2
			l2 = l2.Next
		}
		cur = cur.Next
	}

	if l1 != nil {
		cur.Next = l1
	} else if l2 != nil {
		cur.Next = l2
	}

	return res.Next
}

// 分治合并
func merge(lists []*ListNode, l, r int) *ListNode {
	// 只剩下一个，直接返回即可
	if l == r {
		return lists[l]
	} else if l > r {
		// 异常情况
		return nil
	}

	mid := (l + r) >> 1
	return mergeTwoLists(merge(lists, l, mid), merge(lists, mid+1, r))
}
