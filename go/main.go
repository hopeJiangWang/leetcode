package main

import (
	"fmt"
	"mycode/leetcode/go/problem/top100"
)

func main() {

	// s := []string{"abcw","baz","foo","bar","xtfn","abcdef"}

	// score := []int{10, 3, 8, 9, 7}

	// fmt.Println(problem.NumWaterBottles(9, 3))
	nums := []int{7,1,5,3,6,4}
	// target := 3
	// s := "babad"
	fmt.Printf("res: %v\n", top100.MaxProfit(nums))
}