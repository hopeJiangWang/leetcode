package top100

import "fmt"

/*
输入: nums = [-1,0,3,5,9,12], target = 9
输出: 4

输入: nums = [-1,0,3,5,9,12], target = 2
输出: -1

有序的（升序）整型数组
如果目标值存在返回下标，否则返回 -1。
*/

func Search704(nums []int, target int) int {
	return searchBinary(nums, 0, len(nums)-1, target)
}

func searchBinary(nums []int, left, right, target int) int {
	for left <= right {
		mid := (left + right) >> 1
		fmt.Printf("left: %d, mid: %d, right: %d \n", left, mid, right)
		if nums[mid] == target {
			return mid
		} else if nums[mid] > target {
			right = mid - 1
		} else if nums[mid] < target {
			left = mid + 1
		}
	}
	return -1
}