package top100

import (
	// "fmt"
)

/*
输入：[1,8,6,2,5,4,8,3,7]
输出：49 
解释：图中垂直线代表输入数组 [1,8,6,2,5,4,8,3,7]。在此情况下，容器能够容纳水（表示为蓝色部分）的最大值为 49。

输入：height = [1,1]
输出：1

输入：height = [4,3,2,1,4]
输出：16

输入：height = [1,2,1]
输出：2
*/

func MaxArea(height []int) int {
	/*
		因为要获取更大的面积，所以应该是矮的那一边往中间靠拢
	*/
	res := 0
	
	left, right := 0, len(height) - 1
	
		for left < right {
			hight := min(height[left], height[right])
			res = max(res, hight * (right - left))
			// fmt.Printf("left: %d, right: %d, res: %d \n", left, right, res)
			
			if height[left] < height[right] {
				left++
			} else {
				right--
			}
	
		}
	
		return res
	}
