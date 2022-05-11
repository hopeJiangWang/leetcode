package main

import (
	"fmt"
	"leetcode/go/problem/gj/mayDP"
)

func main() {

	// s := []string{"apple", "pen"}
	// str := "applepenapple"
	// score := []int{10, 3, 8, 9, 7}
	// fmt.Println(problem.NumWaterBottles(9, 3))
	// nums := []int{6,4,8,1,3,2}
	// nums1 := []int{4,7,6,2,3,8,6,1}

	// nums2 := []int{10,9,2,5,3,7,101,18}
	// target := 9
	// s := "()()"
	// matrix := [][]int{{2, 3}, {1, 2}, {3, 4}, {1, 3}}
	//top100.Rotate(matrix)
	//fmt.Println("matrix: ", matrix)
	fmt.Printf("res: %v\n", mayDP.LastRemaining(9))
	// fmt.Printf("res: %v\n", mayDP.LengthOfLIS(nums2))
}
