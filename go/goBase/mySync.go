package main

import (
	"fmt"
	"sync"
)

// 直接崩掉
func main() {
	wg := sync.WaitGroup{}
	wg.Add(1)
	for i := 1; i < 2; i++ {
		go func(i int) {
			fmt.Println(i)
			wg.Done()
			wg.Add(1)
		}(i)
	}
	wg.Wait()
}
