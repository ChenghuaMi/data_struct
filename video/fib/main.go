package main

import "fmt"

//实现 斐波那契数列
//通过 channel 实现数据传输

// ch 数据传输
// n 长度
func fib(ch chan int, n int) {
	a, b := 0, 1
	for i := 0; i < n; i++ {
		ch <- a
		a, b = b, a+b
	}
	//关闭通道
	close(ch)
}
func main() {
	ch := make(chan int)
	n := 15
	go fib(ch, n)
	for v := range ch { //出现死锁，通道没有关闭，解决方法：关闭通道
		fmt.Println(v)
	}
}
