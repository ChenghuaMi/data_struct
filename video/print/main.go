package main

import (
	"fmt"
	"sync"
)

// 今天面试，笔试了一道编程题 10分钟内完成

//输入字符串为“hello world!! 2025.5.27”，用多线程（如5个）打印输出；
//要求：
//1. 每个线程打印一个字母
//2. 打印的字母顺序与输入字符串一致；

// 就这样一个题，
func main() {
	str := "hello world!! 2025.5.27"
	ch1 := make(chan int)
	ch2 := make(chan struct{})
	var wg sync.WaitGroup
	for i := 0; i < 5; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			//有报错，持续修改
			for j := 0; j < len(str); j++ {
				if v, ok := <-ch1; ok {
					fmt.Println(string(str[v]))
					ch2 <- struct{}{}
				}

			}
		}()
	}
	go func() {
		for i := 0; i < len(str); i++ {
			ch1 <- i
			<-ch2
		}
		close(ch1)
	}()
	wg.Wait()

}
