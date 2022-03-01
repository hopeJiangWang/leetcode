package top100

import "fmt"

func reverseList(head *ListNode) *ListNode {
    return listReverse(head, nil)
}

// 反转链表
// func listReverse(head, tail *ListNode) *ListNode {
// 	cur := head
// 	var pre *ListNode

// 	for cur != tail {
// 		pre, cur, cur.Next = cur, cur.Next, pre
// 	}

// 	return pre
// }