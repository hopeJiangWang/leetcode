package main
 
import "fmt"
 
//import "time"
 
//此通道只能写，不能读
func producer(out chan int){
	for i := 0; i < 10; i++ {
		fmt.Println("Producer: ", i)
		out <- i
	}
	close(out)
}
 
//此通大道只能读，不能写
func consumer(in chan int){
	for num := range in {
		fmt.Println("Consumer: ", num)
	}
}
 
func main(){
	//创建双向，channel
	ch := make(chan int)
 
	//生产者，生产数字，写入channel
	//新开一个协程
	go producer(ch) //channel传参，引用传递
 
	//消费者，从channel读取内存，打印
	consumer(ch)
}