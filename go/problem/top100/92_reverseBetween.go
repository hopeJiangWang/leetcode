package top100

/*
left, right 位置之间进行翻转：
	输入：head = [1,2,3,4,5], left = 2, right = 4
	输出：[1,4,3,2,5]
*/

func reverseBetween(head *ListNode, left int, right int) *ListNode {
	cnt := 1
	cur := head
	var tmpHead, tmpTail *ListNode

	for cur != nil {
		if cnt == left {
			tmpHead = cur
		} else if cnt == right {
			tmpTail = cur
			newHead, newTail = myReverse(tmpHead, tmpTail)
		} 
		
		cur = cur.Next
		cnt++
	}
	return head
}
