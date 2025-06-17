package main

import "fmt"

// 实现 =》 并发实现两个channel，合并到同一个 channel中
// 采用通道进行并发控制

func main() {
	ch1, ch2, ch3 := make(chan int), make(chan int), make(chan int)
	//接受通道的数据,写入到 ch3
	go func() {
		for {
			select {
			case num1, ok := <-ch1:
				if ok {
					ch3 <- num1
				} else { //当ch1关闭，ch1 重新赋值 为 nil
					ch1 = nil
				}
			case num2, ok := <-ch2:
				if ok {
					ch3 <- num2
				} else {
					ch2 = nil
				}
			}
			if ch1 == nil && ch2 == nil {
				close(ch3) //关闭ch3,结束协程
				return
			}
		}
	}()
	//往ch1 写入数据
	go func() {
		ch1 <- 1
		ch1 <- 2
		close(ch1)
	}()
	//往ch2 写入数据
	go func() {
		ch2 <- 3
		ch2 <- 4
		close(ch2)
	}()
	//获取合并的数据
	//问题：ch3阻塞，解决=》关闭ch3
	for v := range ch3 {
		fmt.Println(v)
	}
	fmt.Println("main...")
}
