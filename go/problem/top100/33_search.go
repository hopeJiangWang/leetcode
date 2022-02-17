package top100

/*
输入：nums = [4,5,6,7,0,1,2], target = 0
输出：4

输入：nums = [4,5,6,7,0,1,2], target = 3
输出：-1

输入：nums = [1], target = 0
输出：-1
*/

func Search(nums []int, target int) int {
	var mid int = 0
	/*
		直接二分查找：
		可以根据每段数据的边界值来判断需要保留哪一段：
		1. 如果
	*/
	left, right := 0, len(nums)-1
	for left <= right {
		mid = (left + right) >> 1
		if nums[mid] == target {
			return mid
		} else if nums[mid] < nums[right] {
			// 右半边是递增的
			if nums[mid] < target && target <= nums[right] {
				left = mid + 1
			} else {
				right = mid - 1
			}
		} else {
			// 坐半边是递增的
			if nums[left] <= target && target < nums[mid] {
				right = mid - 1
			} else {
				left = mid + 1
			}
		}
	}

	// 没找到，直接返回-1
	return -1
}
