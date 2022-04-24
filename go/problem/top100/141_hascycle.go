package top100

func hasCycle(head *ListNode) bool {
	/*
		使用快慢指针，快的一次跳两个数，慢的那个一次跳一个数
		如果有环的话，肯定能达到fast==slow，否则就没环
	*/

	// 需要注意几点：链表可以为空
    if head == nil {
        return false
    }
    fast, slow := head, head

	// 这里要注意fast为空的几种情况
	for fast != nil {
		if fast.Next == nil {
            return false
        } 
        fast = fast.Next.Next
		slow = slow.Next

		if fast == slow {
			return true
		}
	}

	return false
}