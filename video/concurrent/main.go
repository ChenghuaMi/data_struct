package main

import (
	"fmt"
	"sync"
)

//实现 使用 channel 实现并发安全的 map

type concurrentMap struct {
	mp   map[string]int
	mu   *sync.RWMutex
	ch   chan func()
	wait chan struct{}
}

// 写入数据
func (c *concurrentMap) set(key string, val int) {
	c.mu.Lock()
	defer c.mu.Unlock()
	c.mp[key] = val
}

// 读取数据
func (c *concurrentMap) get(key string) int {
	c.mu.RLock()
	defer c.mu.RUnlock()
	return c.mp[key]
}

// 开启一个协程 获取 channel 中的 func
func (c *concurrentMap) start() {
	go func() {
		for f := range c.ch {
			f()
		}
		close(c.wait)
	}()
}
func (c *concurrentMap) stop() {
	v, ok := <-c.wait
	fmt.Println("v=", v, "ok=", ok)
}
func main() {
	c := &concurrentMap{
		mp:   make(map[string]int),
		mu:   &sync.RWMutex{},
		ch:   make(chan func()),
		wait: make(chan struct{}),
	}
	c.start()
	c.ch <- func() {
		c.set("key1", 1)
	}
	c.ch <- func() {
		v := c.get("key1")
		fmt.Println("get value=", v)
	}
	close(c.ch)
	c.stop()
}
