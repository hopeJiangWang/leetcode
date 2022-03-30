package top100

import "sort"

/*
输入: [3,2,1,5,6,4] 和 k = 2
输出: 5

输入: [3,2,3,1,2,4,5,5,6] 和 k = 4
输出: 4

*/

func FindKthLargest(nums []int, k int) int {
	sort.Ints(nums)

	return nums[len(nums)-k]
}

func FindKthLargest2(nums []int, k int) int {
	
}