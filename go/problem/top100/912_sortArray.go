package top100

/*
输入：nums = [5,1,1,2,0,0]
输出：[0,0,1,1,2,5]

*/

func sortArray(nums []int) []int {
	quickSort2(nums, 0, len(nums)-1)
	return nums
}

func quickSort2(nums []int, left, right int) {
	if left < right {
		pivot := partition2(nums, left, right)

		quickSort2(nums, left, pivot-1)
		quickSort2(nums, pivot+1, right)
	}
}

func partition2(nums []int, left, right int) int {
	pivot := nums[left]

	for left < right {
		// 先从右边找到小于基准值的数
		for left < right && nums[right] >= pivot {
			right--
		}

		// 用这个数更新基准值位置
		nums[left] = nums[right]

		// 接下来从左边找大于基准值的数
		for left < right && nums[left] < pivot {
			left++
		}

		// 用这个数更新之前空出来的位置
		nums[right] = nums[left]
	}
	// 最后更新基准值到最“中间”的位置
	nums[left] = pivot
	return left
}