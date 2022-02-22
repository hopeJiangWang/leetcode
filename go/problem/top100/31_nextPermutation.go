package top100

import "fmt"

/*
输入：nums = [1,2,3]
输出：[1,3,2]

输入：nums = [3,2,1]
输出：[1,2,3]

输入：nums = [1,1,5]
输出：[1,5,1]
*/

func NextPermutation(nums []int) {
	/*
		从左至右遍历，找到第一个nums[f] < nums[f+1]的位置，
		然后在[f+1, len(nums))中，找到最小的大于nums[f]的值，交换
		然后将[f+1, len(nums))的所有数从小到大排序一下
	*/
	if len(nums) == 0 {
		return
	}

	index := -1
	for i := len(nums) - 1; i > 0; i-- {
		// 找到第一个 nums[f] < nums[f+1]的位置
		if nums[i-1] < nums[i] {
			index = i - 1
			fmt.Println("index: ", nums[index])
			// 在[f+1, len(nums))中，找到最小的大于nums[f]的值，交换
			// 因为这一段是从右至左递增的，所有从右开始找到第一个大于目标值的即可
			target := nums[i-1]
			for j := len(nums) - 1; j >= index; j-- {
				if nums[j] > target {
					swap(nums, index, j)
					BubbleSort(&nums, index+1, len(nums)-1)
					break
				}
			}
			break
		}
	}

	if index == -1 {
		BubbleSort(&nums, 0, len(nums)-1)
	}
}

func swap(nums []int, i, j int) {
	tmp := nums[i]
	nums[i] = nums[j]
	nums[j] = tmp
}

func BubbleSort(array *[]int, begin, end int) {
	if array == nil {
		return
	}

	for i := begin; i < end; i++ {
		flag := false
		for j := begin; j < end; j++ {
			if (*array)[j+1] < (*array)[j] {
				fmt.Println("--- nums: ", (*array)[j], ", ", (*array)[j+1])
				(*array)[j], (*array)[j+1] = (*array)[j+1], (*array)[j]
				flag = true
			}
		}
		if !flag {
			break
		}
	}
}

func NextPermutation2(nums []int) {
	/*
	   从右至左找到第一个左边数小于右边数的位置index，
	   然后从index的右边，找到大于它的最小的数字，将其交换后，排序右边所有的数
	*/
	i := len(nums) - 2
	for i >= 0 && nums[i] >= nums[i+1] {
		i--
	}

	if i >= 0 {
		fmt.Println("target: ", nums[i], "i: ", i)
		for j := len(nums) - 1; j >= 0; j-- {
			if nums[j] > nums[i] {
				nums[i], nums[j] = nums[j], nums[i]
				break
			}
		}
	}
	reverse(nums[i+1:])
}

func reverse(a []int) {
	n := len(a)
	for i := 0; i < n/2; i++ {
		a[i], a[n-1-i] = a[n-1-i], a[i]
	}
}
