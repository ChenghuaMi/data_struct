package main

import (
	"fmt"
	"sync"
)

// 实现 work pool, 多个work 处理任务
func work(wg *sync.WaitGroup, workId int, jobs chan int) {
	defer wg.Done()
	for job := range jobs {
		fmt.Printf("workId %d,job %d \n", workId, job)
	}
}
func main() {
	jobs := make(chan int, 10)
	wg := sync.WaitGroup{}
	for i := 1; i <= 3; i++ {
		wg.Add(1)
		go work(&wg, i, jobs)
	}
	for i := 1; i <= 5; i++ {
		jobs <- i
	}
	//关闭通道
	close(jobs)
	wg.Wait()
}
