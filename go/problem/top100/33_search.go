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
	/*
	   变异二分法：
	   1. 先找出中间值，如果它比右边界小，那么右边是有序的，反之亦然；
	   2. 随后就在右边界里面查找即可（按有无调整边界值）
	   如果中间的数小于最右边的数，则右半段是有序的，若中间数大于最右边数，则左半段是有序的，
	   我们只要在有序的半段里用首尾两个数组来判断目标值是否在这一区域内，这样就可以确定保留哪半边了
	*/
	left, right := 0, len(nums)-1

	for left <= right {
		mid := (left + right) >> 1

		if nums[mid] == target {
			return mid
		} else if nums[mid] < nums[right] {
			// 中间值小于右边界，找右边界
			if nums[mid] < target && nums[right] >= target {
				left = mid + 1
			} else {
				right = mid - 1
			}
		} else {
			if nums[left] <= target && nums[mid] > target {
				right = mid - 1
			} else {
				left = mid + 1
			}
		}
	}

	return -1
}

// func Search(nums []int, target int) int {
// 	var mid int = 0
// 	/*
// 		直接二分查找：
// 		可以根据每段数据的边界值来判断需要保留哪一段：
// 		1. 如果
// 	*/
// 	left, right := 0, len(nums)-1
// 	for left <= right {
// 		mid = (left + right) >> 1
// 		if nums[mid] == target {
// 			return mid
// 		} else if nums[mid] < nums[right] {
// 			// 右半边是递增的
// 			if nums[mid] < target && target <= nums[right] {
// 				left = mid + 1
// 			} else {
// 				right = mid - 1
// 			}
// 		} else {
// 			// 坐半边是递增的
// 			if nums[left] <= target && target < nums[mid] {
// 				right = mid - 1
// 			} else {
// 				left = mid + 1
// 			}
// 		}
// 	}

// 	// 没找到，直接返回-1
// 	return -1
// }
