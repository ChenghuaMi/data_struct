package main

import "fmt"

// 实现 使用goroutine 和channel 实现并发查找数组中的最大值
func findMax(ch chan int, sli []int) {
	maxOne := sli[0]
	for _, v := range sli {
		if maxOne < v {
			maxOne = v
		}
	}
	ch <- maxOne
}

func main() {
	ch := make(chan int)
	sli := []int{1, 2, 8, 3, 4, 19, 3, 65, 44}
	go findMax(ch, sli)
	fmt.Println(<-ch)
}
