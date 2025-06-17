package main

import (
	"fmt"
	"time"
)

// 实现 两个 goroutine 并发的 交替 打印 即偶数 1-10

func main() {
	ch1 := make(chan struct{})
	ch2 := make(chan struct{})
	//ch3 := make(chan struct{})
	go func() {
		for i := 1; i <= 10; i = i + 2 {
			<-ch1
			fmt.Println(i)
			ch2 <- struct{}{}
		}
		//close(ch2)
	}()
	go func() {
		for i := 2; i <= 10; i = i + 2 {
			<-ch2
			fmt.Println(i)
			ch1 <- struct{}{}
		}
		close(ch1)
	}()
	//ch2 <- struct{}{}
	ch1 <- struct{}{}
	//<-ch3
	time.Sleep(10 * time.Second)
	fmt.Println("main...")

}
