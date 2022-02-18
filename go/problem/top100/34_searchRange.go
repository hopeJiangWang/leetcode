package top100

/*
输入：nums = [5,7,7,8,8,10], target = 8
输出：[3,4]

输入：nums = [5,7,7,8,8,10], target = 6
输出：[-1,-1]

输入：nums = [], target = 0
输出：[-1,-1]
*/

func SearchRange(nums []int, target int) []int {
	/*
		二分法查询：先找到一个目标值，然后以此往两边查找到边界值
	*/
	var numsLen int = len(nums)
	if numsLen == 0 || nums[0] > target || nums[numsLen-1] < target {
		return []int{-1, -1}
	}

	left := binarySearchLeft(nums, target)
	right := binarySearchRight(nums, target)

	return []int{left, right}
}

func binarySearchLeft(nums []int, target int) int {
	// 先找到左边界
	mid, left, right := 0, 0, len(nums)-1
	for left <= right {
		mid = (left + right) >> 1
		if nums[mid] == target {
			right = mid - 1 // 继续往左边找
		} else if nums[mid] > target {
			right = mid - 1
		} else {
			left = mid + 1
		}
	}

	if nums[left] == target {
		return left
	}
	return -1
}

func binarySearchRight(nums []int, target int) int {
	// 先找到左边界
	mid, left, right := 0, 0, len(nums)-1
	for left <= right {
		mid = (left + right) >> 1
		if nums[mid] == target {
			left = mid + 1 // 继续往右边找
		} else if nums[mid] > target {
			right = mid - 1
		} else {
			left = mid + 1
		}
	}

	if nums[right] == target {
		return right
	}
	return -1
}
