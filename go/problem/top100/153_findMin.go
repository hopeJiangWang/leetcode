package top100

import "fmt"

/*
输入：nums = [4,5,6,7,0,1,2]
输出：0

输入：nums = [11,13,15,17]
输出：11

3,4,5,1,2
1
*/

func FindMin(nums []int) int {
	/*
		反正这个一看就应该是要用二分法来做了：
		nums[i-1] > nums[i] || i == 0
	*/
	return findBroken(nums, 0, len(nums)-1)
}

func findBroken(nums []int, left, right int) int {
	for left + 1 < right {
		mid := (left + right) >> 1
		// 只剩一个数，或者找到了nums[i-1] > nums[i]
		if (mid >= 1 && nums[mid-1] > nums[mid]) || mid == 0 {
			return nums[mid]
			// 右边不有序，最小值在右边
		} else if nums[mid+1] > nums[right] {
			left = mid
			// 左边不有序，最小值在左边
		} else if mid >= 1 && nums[mid-1] < nums[left] {
			right = mid
		} else {
			// 左右两端都是顺序的
			if nums[left] < nums[mid+1] {
				right = mid
			} else {
				left = mid
			}
		}
	}

	// 剩下两个数，取最小值
	return min(nums[left], nums[right])
}

func FindMin2(nums []int) int {
	left, right := 0, len(nums)-1

	for left < right {
		mid := (left + right) >> 1
		if nums[mid] < nums[right] {
			right = mid
		} else {
			left = mid + 1
		}
	}
	return nums[left]
}