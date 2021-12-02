package main

import (
	"fmt"
	"mycode/leetcode/go/problem/everyday"
)

func main() {

	// s := []string{"abcw","baz","foo","bar","xtfn","abcdef"}

	score := []int{10, 3, 8, 9, 7}

	fmt.Println(problem.FindRelativeRanks(score))
}
