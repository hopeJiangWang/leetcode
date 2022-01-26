package top100

import (
	"fmt"
)

func TwoSum(nums []int, target int) []int {
	tmpMap := make(map[int]int)
	res := make([]int, 0)
	for i := 0; i< len(nums); i++ {
		_, ok := tmpMap[target - nums[i]] 
		fmt.Println("target: ", target, "nums[i]: ", nums[i], "find: ", target - nums[i])
		if ok {
			return []int{i, tmpMap[target - nums[i]]}
		}
		tmpMap[nums[i]] = i
	}
	return res
}
