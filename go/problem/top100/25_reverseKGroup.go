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

// 反转链表
func listReverse(head, tail *ListNode) *ListNode {
	cur := head
	var pre *ListNode

	/*
		1.先将pre往后移，指向cur；
		2.cur指向cur.Next，往后移；
		3.反转，cur.Next指向pre；
	*/
	for cur != tail {
		pre, cur, cur.Next = cur, cur.Next, pre
	}

	return pre
}

func reverseKGroup2(head *ListNode, k int) *ListNode {
	cur := head
	for i := 0; i < k; i++ {
		if cur == nil {
			return head
		}
		cur = cur.Next
	}
	newHead := listReverse(head, cur)
	head.Next = reverseKGroup2(cur, k)
	return newHead
}

// 直接模拟

/*
	时间：O(n) = O(n/k)*O(k)
*/
func reverseKGroup3(head *ListNode, k int) *ListNode {
	res := &ListNode{Next: head}
	pre := res

	for head != nil {
		tail := pre
		for i := 0; i < k; i++ {
			tail = tail.Next
			// 不足k个，直接返回剩下的即可
			if tail == nil {
				return pre.Next
			}
		}

		// 满足k个，则需要将head->tail之间的进行反转，并且将其拼接回去
		next := tail.Next // 记录下一个节点
		head, tail = reverseHeadAndTail(head, tail)

		// 开始进行下一个阶段的操作
		pre.Next = head
		pre = tail
		tail.Next = next
		head = next
	}
	return res.Next
}

// 反转头尾之间的链表
func reverseHeadAndTail(head, tail *ListNode) (*ListNode, *ListNode) {
	pre := tail.Next // 因为头节点需要变成尾节点
	cur := head

	for pre != tail {
		next := cur.Next
		cur.Next = pre
		pre = cur
		cur = next
	}
	return tail, head
}
