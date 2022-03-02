package top100

// import "fmt"

func reverseList(head *ListNode) *ListNode {
    return listReverse(head, nil)
}

func reverseList2(head *ListNode) {
    var pre *ListNode
    cur := head

    for cur != nil {
        next := cur.Next    // 先获取下一个节点
        cur.Next = pre      // 反转当前节点的next指针
        pre = cur           // 更新前置节点
        cur = next          // 跳到下一个节点
    }
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