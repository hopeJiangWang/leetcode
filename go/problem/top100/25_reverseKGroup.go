package top100

import "fmt"

func reverseKGroup1(head *ListNode, k int) *ListNode {
	/*
		可以用一个双端队列：
		1. 累计到k个数，就倒序输出，加到结果链表尾部
		2. 没有累计到k个数就到链表尾部了，就正序输出，加到结果链表尾部
	*/
	res := &ListNode{Val: -1}
	cur := res
	for head != nil {
		cnt := 0
		var deque []int
		for cnt < k && head != nil {
			fmt.Println("now: ", head.Val, "deque: ", deque)
			deque = append(deque, head.Val)
			head = head.Next
			cnt++
		}

		if cnt == k {
			for cnt > 0 {
				cur.Next = &ListNode{Val: deque[cnt-1]}
				cur = cur.Next
				fmt.Println("deque[cnt]: ", deque[cnt-1], "cnt: ", cnt)
				cnt--
			}
		} else {
			for i := 0; i < cnt; i++ {
				cur.Next = &ListNode{Val: deque[i]}
				cur = cur.Next
				fmt.Println("deque[i]: ", deque[i], "cnt: ", i)
			}
		}
	}

	return res.Next
}

func reverseKGroup(head *ListNode, k int) *ListNode {
	/*
		常数额外空间的算法来解决此问题
		不只是单纯的改变节点内部的值，而是需要实际进行节点交换

		那么就是找到k个节点了，就反转他；否则，维持原状即可
	*/
	res := &ListNode{Val: -1}
	res.Next = head

	prev := res
	next := nil
	cur := res.Next

	lenOfList := 0
	for head != nil {
		lenOfList++
		head = head.Next
	}

	for i := 0; i < lenOfList/k; i++ {
		// 满足k个，翻转它
		for j := 0; j < k; j++ {
			next = cur.Next
			cur.Next = next.Next
			next.Next = prev.Next
			prev.Next = next
		}
		// 不足k个，直接遍历即可

	}

	return res.Next
}
func reverseKGroup2(head *ListNode, k int) *ListNode {
	hair := &ListNode{Next: head}
	pre := hair

	for head != nil {
		tail := pre
		for i := 0; i < k; i++ {
			tail = tail.Next
			if tail == nil {
				return hair.Next
			}
		}
		nex := tail.Next
		head, tail = myReverse(head, tail)
		pre.Next = head
		tail.Next = nex
		pre = tail
		head = tail.Next
	}
	return hair.Next
}

func myReverse(head, tail *ListNode) (*ListNode, *ListNode) {
	prev := tail.Next
	p := head
	for prev != tail {
		nex := p.Next
		p.Next = prev
		prev = p
		p = nex
	}
	return tail, head
}
