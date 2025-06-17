package main

import "fmt"

// 实现将 一个切片 拆分 成多个子切片，并发的计算数组的和
// 使用channel 进行 和 的传递，求和

func sliSum(ch chan int, sli []int) {
	sum := 0
	for _, v := range sli {
		sum += v
	}
	ch <- sum
}
func main() {
	sli := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10} // 切片求和
	ch := make(chan int)
	go sliSum(ch, sli[:5])
	go sliSum(ch, sli[5:])
	fmt.Println("sum=", <-ch+<-ch)
}
