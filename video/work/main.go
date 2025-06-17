package main

import (
	"fmt"
	"runtime"
	"sync"
	"time"
)

// 使用 channel 进行流量控制
// 使用带缓冲通道的channel 进行限制
func work(ch chan int, wg *sync.WaitGroup, i int) {
	defer wg.Done()
	ch <- i
}
func main() {
	//打印出 协程 的数量、
	defer func() {
		time.Sleep(time.Second)
		fmt.Println("gorotine num:", runtime.NumGoroutine())
	}()
	ch := make(chan int, 3)           //接收数据
	limitCh := make(chan struct{}, 5) //限制5个并发请求
	//使用waitgroup 并发处理
	var wg sync.WaitGroup
	for i := 0; i < 10; i++ {
		wg.Add(1)
		go func(i int) {
			limitCh <- struct{}{}
			//工作方法
			work(ch, &wg, i)
			<-limitCh
		}(i)
	}
	// 出现死锁，通道没有关闭，解决方法，使用协程 获取数据
	//改造如下
	go func() {
		//还是存在协成没有关闭
		for v := range ch {
			fmt.Println("get num:", v)
		}
	}()

	wg.Wait()
	//关闭通道
	close(ch)
}
