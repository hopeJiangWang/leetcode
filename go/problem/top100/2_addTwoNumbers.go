package top100

import (
	// "fmt"
)

type ListNode struct {
	Val int
	Next *ListNode
}

/*
输入：l1 = [2,4,3], l2 = [5,6,4]
输出：[7,0,8]
解释：342 + 465 = 807.
*/

func AddTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	res := &ListNode{Val: -1}
	curr := res
	var carry int = 0

	for l1 != nil || l2 != nil {
		val1, val2 := 0, 0
		if l1 != nil {
			val1 = l1.Val
			l1 = l1.Next
		}

		if l2 != nil {
			val2 = l2.Val
			l2 = l2.Next
		}

		curr.Next = &ListNode{Val: (carry + val1 + val2) % 10}
		curr = curr.Next
		carry = (carry + val1 + val2) / 10
	}

	// 遍历之后，还有的剩余
	if carry != 0 {
		curr.Next = &ListNode{Val: carry}
	}

	return res.Next
}