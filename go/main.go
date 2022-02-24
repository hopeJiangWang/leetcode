package main

import (
	"fmt"
	"mycode/leetcode/go/problem/top100"
)

func main() {

	//s := []string{"eat", "tea", "tan", "ate", "nat", "bat"}

	// score := []int{10, 3, 8, 9, 7}

	// fmt.Println(problem.NumWaterBottles(9, 3))
	nums := []int{2, 3, 1, 1, 4}
	// target := 3
	// s := "()()"
	//matrix := [][]int{{1,2,3},{4,5,6},{7,8,9}}
	//top100.Rotate(matrix)
	//fmt.Println("matrix: ", matrix)

	fmt.Printf("res: %v\n", top100.CanJump(nums))
}
