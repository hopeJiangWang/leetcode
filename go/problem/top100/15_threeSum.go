package top100

import (
	"sort"
)

func ThreeSum(nums []int) [][]int {
	/*
		1. 首先需要数组长度大于等于3；
		2. 可以先排序下，如果最小的数已经是大于0了，或者说最大的数已经是小于0的，那就可以直接返回了；
		3. 然后就可以从前往后遍历，先确定一个数，然后找另外两个数（和相反的）；
	*/

	var res [][]int

	if len(nums) < 3 {
		return res
	}

	sort.Ints(nums)

	if nums[0] > 0 || nums[len(nums)-1] < 0 {
		return res
	}

	for i := 0; i < len(nums); i++ {
		// 如果最小的数已经是大于0的了，那么直接跳过
		if nums[i] > 0 {
			continue
		}
		// 如果有相同数字，跳过前面相同的
		if i > 0 && nums[i] == nums[i-1] {
			continue
		}

		// 这时候就是确定了一个数，另外找两个数了
		l := i + 1
		r := len(nums) - 1
		
		for l < r {
			// 如果找到了这样的三个数，直接放入结果集
			if nums[i] + nums[l] + nums[r] == 0 {
				res = append(res, []int{nums[i], nums[l], nums[r]})

				// 如果有重复的数字，一样跳过它
				for l < r && nums[l] == nums[l+1] {
					l++
				}

				for l < r && nums[r] == nums[r-1] {
					r--
				}
				// 这里需要继续往下走 ！！！
				l++
				r--
			} else if nums[i] + nums[l] + nums[r] > 0 {
				// 如果当前和大于0，那么右指针左移
				r--
			} else {
				// 如果当前和小于0，那么左指针右移
				l++
			} 
		}
	}

	return res
}