package main

import (
	"fmt"
	"leetcode/go/problem/gj/juneinterval"
	// "leetcode/go/problem/everyday"
)

func main() {

	// s := []string{"rrjk","furt","guzm"}
	// str := "applepenapple"
	// score := []int{10, 3, 8, 9, 7}
	// fmt.Println(problem.NumWaterBottles(9, 3))
	// nums := []int{6,4,8,1,3,2}
	// nums1 := []int{4,7,6,2,3,8,6,1}

	nums2 := []int{1, 7, 3, 6, 5, 6}
	// target := 9
	// s := "()()"[5,4],[6,4],[6,7],[2,3]
	// matrix := [][]int{{5,4}, {6,4}, {6,7}, {2, 3}}
	//top100.Rotate(matrix)
	//fmt.Println("matrix: ", matrix)
	// fmt.Printf("res: %v\n", everyday.OneEditAway("aaaaa", "aaaab"))

	fmt.Printf("res: %v\n", juneinterval.PivotIndex(nums2))
}
