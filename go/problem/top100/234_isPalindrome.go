package top100

/*
输入：head = [1,2,2,1]
输出：true

输入：head = [1,2]
输出：true

回文链表
*/

func isPalindrome(head *ListNode) bool {
	/*
		1、先直接遍历一遍链表，把所有的数都取出来
		2、然后再遍历一遍链表，和之前取出的数反过来对比
		这时候可以直接处理获取到的数组了
	*/

	var arr []int
	for head != nil {
		arr = append(arr, head.Val)
		head = head.Next
	}

	for first, second := 0, len(arr)-1; first < second;  {
		if arr[first] != arr[second] {
			return false
		}
        first++
        second--
	}
	return true
}