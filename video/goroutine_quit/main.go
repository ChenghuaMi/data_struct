package main

import "fmt"

//实现 如何使用channel 优雅的关闭 goroutine

func main() {
	ch := make(chan int)        //写入数据
	quit := make(chan struct{}) //关闭通道，推出协程
	go func() {
		for {
			select {
			case num, ok := <-ch:
				if ok {
					fmt.Println("get num:", num)
				}
			case <-quit:
				fmt.Println("quit...")
				return
			}
		}
	}()
	for i := 0; i < 10; i++ {
		ch <- i
	}
	close(ch)
	quit <- struct{}{}
	fmt.Println("main....")
}
