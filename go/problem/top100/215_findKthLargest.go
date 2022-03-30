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
	/*
		那就自己实现排序：快排
	*/
	quickSort(nums, 0, len(nums)-1)
	return nums[len(nums)-k]
}

func quickSort(nums []int, left, right int) {
	if left < right {
		// 先确定基准值
		pivot := partition(nums, left, right)
		// 排序左边
		quickSort(nums, left, pivot-1)
		// 排序右边
		quickSort(nums, pivot+1, right)
	}
}

// 快排的核心，找一个基准，然后将所有大于它的
func partition(nums []int, left, right int) int {
	// 先取基准值
	pivot := nums[left]

	for left < right {
		// 从右边开始找，找到第一个小于基准值的数，将其和基准值交换，这时候右边就都是大于基准值的数了
		for left < right && nums[right] >= pivot {
			right--
		}

		nums[left] = nums[right]

		// 开始从左边开始找
		for left < right && nums[left] < pivot {
			left++
		}
		nums[right] = nums[left]
	}

	nums[left] = pivot
	return left
}