package main

import (
	"fmt"
	"mycode/leetcode/go/problem/top100"
)

func main() {

	// s := []string{"apple", "pen"}
	// str := "applepenapple"
	// score := []int{10, 3, 8, 9, 7}

	// fmt.Println(problem.NumWaterBottles(9, 3))
	nums := []int{1,1,2}
	target := 3
	// s := "()()"
	// matrix := [][]int{{2, 3}, {1, 2}, {3, 4}, {1, 3}}
	//top100.Rotate(matrix)
	//fmt.Println("matrix: ", matrix)

	fmt.Printf("res: %v\n", top100.TwoSum(nums, target))
}
