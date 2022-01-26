package main

import (
	"fmt"
	"mycode/leetcode/go/problem/top100"
)

func main() {

	// s := []string{"abcw","baz","foo","bar","xtfn","abcdef"}

	// score := []int{10, 3, 8, 9, 7}

	// fmt.Println(problem.NumWaterBottles(9, 3))
	// nums := []int{2,7,11,15}
	// target := 9
	num1 := []int{1,2}
	num2 := []int{3,4}
	fmt.Printf("res: %v\n", top100.FindMedianSortedArrays(num1, num2))
}