package main

import (
	"fmt"
	"time"
)

/*
return, defer, 返回值的执行调度：
1. 



*/



func myFunc0() {
	p := 29

	defer func(test int) {
		fmt.Println("func0: ", test)
	}(p)

	p = 30
	defer func(test int) {
		fmt.Println("func1: ", test)
	}(p)

	p = 31
	fmt.Println("func2: ", p)
	time.Sleep(time.Duration(2) * time.Second)
}




func main() {
	myFunc0()
}