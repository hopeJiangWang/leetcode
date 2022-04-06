package main

import (
	"fmt"
)

/*
begin:  [0 0 0]
after append:  [0 0 0 1]
end:  [2 0 0 1]
*/

func main() {
	s := make([]int, 3, 3)
    fmt.Println("begin: ", s)

	s = append(s, 1)
    fmt.Println("after append: ", s)
	s[0] = 2

	fmt.Println("end: ", s)
}