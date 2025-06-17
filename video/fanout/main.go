package main

import (
	"fmt"
	"time"
)

func fanOut(ch chan int, num int) {
	for i := 0; i < num; i++ {
		ch <- i
	}
	close(ch)
}

func main() {
	ch := make(chan int)

	go fanOut(ch, 5)
	for i := 0; i < 3; i++ {
		go func() {
			for num := range ch {
				fmt.Println(num)
			}
		}()
	}
	time.Sleep(time.Second * 5)
}
