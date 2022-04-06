package main

import (
	"fmt"
	"strings"
)

// 判断字符串中字符是否全都不同
func isUniqueString(s string) bool {
	if len(s) > 3000 {
		return  false
	}
	for _,v := range s {
		if v > 127 {
			return false
		}
		if strings.Count(s,string(v)) > 1 {
			return false
		}
	}
	return true
}

// 判断两个给定的字符串排序后是否一致
func isRegroup(s1, s2 string) bool {
	sl1 := len(s1)
	sl2 := len(s2)

	if sl1 > 5000 || sl2 > 5000 || sl1 != sl2 {
		return false
	}

	// 长度一致，然后每个字符的个数相同
	for _, v := range s1 {
		if strings.Count(s1, string(v)) != strings.Count(s2, string(v)) {
			return false
		}
	}
	return true
}

// 交替打印数字和字母: 12AB34CD56
func goroutinPrint() {
    letter := make(chan bool)
    number := make(chan bool)
	wait := sync.WaitGroup{}

	go func() {
		i := 1
		for {
			select {
			case <-number:
				fmt.Print(i)
				i++
				fmt.Print(i)
				i++
				letter <- true
				break
			default:
				break
			}
		}
	}()
	wait.Add(1)
	go func(wait *sync.WaitGroup) {
		str := "ABCDEFGHIJKLMNOPQRSTUVWXYZ"

		i := 0
		for{
			select {
			case <-letter: 
                if i >= strings.Count(str,"")-1 {
					wait.Done()
					return
				}
				fmt.Print(str[i:i+1])
				i++
				fmt.Print(str[i:i+1])
				i++
				number <- true
				break
			default:
				break
			}

		}
	}(&wait)
	number<-true
	wait.Wait()
}

func main() {
	s := "test"
	fmt.Println("end: ", isUniqueString(s))
}