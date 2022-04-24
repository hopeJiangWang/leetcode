package top100 

func detectCycle(head *ListNode) *ListNode {
	/*
		起点到入环点的距离为a，入环点到相遇点的距离为b，环长度为b+c
		则：2(a+b) = a+(b+c)n+b
		==>a+b = (b+c)n
		==>a = (n-1)(b+c)+c
		即，慢指针从相遇点继续走c，同时用第三个指针从开头开始走，他们会在入环点相遇
	*/
	if head == nil {
        return nil
    }
    fast, slow := head, head

	hasCycle := false
	// 这里要注意fast为空的几种情况
	for fast != nil {
		if fast.Next == nil {
            return nil
        } 
        fast = fast.Next.Next
		slow = slow.Next

		if fast == slow {
			hasCycle = true
			break
		}
	}

	// 有环
	if hasCycle == true {
		third := head
		for third != slow {
			third = third.Next
			slow = slow.Next
		}
		return third
	}
	return nil
}