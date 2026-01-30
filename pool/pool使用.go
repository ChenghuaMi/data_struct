package main

import (
	"fmt"
	"sync"
	"time"
)

type user struct {
	id   int
	name string
}

var pool = sync.Pool{
	New: func() any {
		return &user{}
	},
}

func main() {
	wg := sync.WaitGroup{}
	for i := 1; i < 100; i++ {
		wg.Add(1)
		go func(i int) {
			defer wg.Done()
			u := pool.Get().(*user) // 获取对象，没有就创建
			u.id = i
			u.name = fmt.Sprintf("user-id-%d", i)
			fmt.Printf("user %#v\n", u)
			time.Sleep(time.Second)
			u.id = 0
			u.name = "" // 重置
			pool.Put(u) //放回池中

		}(i)
	}
	wg.Wait()
	fmt.Println("main..")
}
