package main

import (
	"fmt"
	"time"
)

// 两个协程 交替的打印出数字
//第一个写成 1，3，5，7，9
//第二个协程 2，4，6，8，10
// 最终打印 1，2，3，4，5，6，7，8，9，10
// 或者 2，1，4，3，6，5，8，7，10，9

func main() {
	ch1, ch2 := make(chan struct{}), make(chan struct{})
	go func() {
		for i := 1; i <= 10; i += 2 {
			<-ch1
			fmt.Println(i)
			ch2 <- struct{}{}
		}
	}()
	go func() {
		for i := 2; i <= 10; i += 2 {
			<-ch2
			fmt.Println(i)
			ch1 <- struct{}{}
		}
	}()
	//没有开启开关
	//ch1 <- struct{}{}
	ch2 <- struct{}{}
	time.Sleep(time.Second * 10)
}
