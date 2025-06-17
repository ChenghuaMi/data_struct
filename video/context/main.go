package main

import (
	"context"
	"fmt"
	"sync"
	"time"
)

// 批量 goroutine 统一退出控制
// 实现主协程退出，所有的子协程 都 退出

func main() {
	ctx, cancel := context.WithCancel(context.Background())
	var wg sync.WaitGroup
	for i := 1; i <= 3; i++ {
		wg.Add(1)
		go func(workId int) {
			defer wg.Done()
			for {
				select {
				case <-ctx.Done():
					fmt.Printf("goroutine %d exit \n", workId)
					return // 退出循环
				default:
					fmt.Printf("goroutine %d is working \n", workId)
					time.Sleep(time.Millisecond * 500) // 睡眠 500毫秒
				}
			}
		}(i)
	}
	time.Sleep(2 * time.Second) // 睡眠2 秒
	cancel()                    //执行取消操作，通知 goroutine
	wg.Wait()
	fmt.Println("main exit...")
}
