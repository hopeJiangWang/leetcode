package main

import (
	"fmt"
	"mycode/leetcode/go/problem/top100"
)

func main() {

	// s := []string{"abcw","baz","foo","bar","xtfn","abcdef"}

	// score := []int{10, 3, 8, 9, 7}

	// fmt.Println(problem.NumWaterBottles(9, 3))
	nums := []int{2, 3, 6, 7}
	// target := 3
	// s := "()()"

	fmt.Printf("res: %v\n", top100.CombinationSum(nums, 8))
}
