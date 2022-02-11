package top100

import (
	// "fmt"
)

/*
输入：head = [1,2,3,4,5], n = 2
输出：[1,2,3,5]

输入：head = [1], n = 1
输出：[]

输入：head = [1,2], n = 1
输出：[1]
*/

func RemoveNthFromEnd(head *ListNode, n int) *ListNode {
	/*
		用两个指针，间隔n长度，等头指针到达链表尾部，就可以删除第二个
		指针对应的节点了
	*/
	first, second := head, head

	// 找到需要删除的节点，这里使用头节点的next来做循环出口
	// 这样子，二号节点对应的就是需要删除的节点的前一个节点
	// 方便删除
	cnt := 0
	for first.Next != nil {
        first = first.Next
		if cnt < n {
			cnt++
		} else {
            second = second.Next
        }
	}

	// 删除找到的节点
	// 这一步操作需要保证不会为空，也即链表长度够n 
    if cnt == n {
        second.Next = second.Next.Next
        return head
    } else {
		// 删除第一个节点
        return second.Next
    }
}
