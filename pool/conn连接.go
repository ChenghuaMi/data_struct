package main

import (
	"fmt"
	"sync"
)

type conn struct {
	Id   int
	Addr string
}
type connPool struct {
	sync.Pool
}

func NewConnPool(addr string) *connPool {
	return &connPool{
		Pool: sync.Pool{
			New: func() any {
				return &conn{
					Id:   0,
					Addr: addr,
				}
			},
		},
	}
}
func (c *connPool) Get() *conn {
	cn, ok := c.Pool.Get().(*conn)
	if !ok {
		fmt.Println("断言失败")
		return nil
	}
	fmt.Printf("conn ptr %p\n", cn)
	return cn
}
func (c *connPool) Put(x any) {
	c.Pool.Put(x)
}
func main() {
	p := NewConnPool("127.0.0.1")
	var wg sync.WaitGroup
	for i := 0; i < 10; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			con := p.Get()
			fmt.Println("con is nil:", con == nil)
			if con != nil {
				p.Put(con)
			}

		}()
	}
	wg.Wait()
	fmt.Println("main...")
}
