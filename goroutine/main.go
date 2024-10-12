package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

type Queue struct {
	tk      *time.Ticker
	array   [5]int
	maxSize int
	front   int
	tail    int
	ch      chan int
	mu      sync.Mutex
	send    chan struct{}
}

func NewQueue(num time.Duration, maxSize int) *Queue {
	return &Queue{
		tk:      time.NewTicker(time.Second * num),
		maxSize: maxSize,
		ch:      make(chan int),
		send:    make(chan struct{}),
	}
}
func (q *Queue) stop() {
	q.tk.Stop()
}
func (q *Queue) add(n int) {
	if q.isFull() {
		fmt.Println("full")
		q.send <- struct{}{}
	} else {
		q.tail = (q.tail + 1) % q.maxSize
		q.array[q.tail] = n

		fmt.Println(q.array)
	}

}
func (q *Queue) tkData() {
	for {
		select {
		case <-q.tk.C:
			q.add(rand.Intn(100))
		}
	}
}
func (q *Queue) get(i int) {
	for {
		select {
		case <-q.send:
			fmt.Println("get..")
			q.mu.Lock()

			q.front = (q.front + 1) % q.maxSize
			v := q.array[q.front]
			fmt.Printf("goroutine:%d,value:%d\n", i, v)
			q.mu.Unlock()
		}
	}

}
func (q *Queue) isFull() bool {
	return (q.tail+1)%q.maxSize == q.front
}
func (q *Queue) isEmpty() bool {
	return q.tail == q.front
}
func main() {
	q := NewQueue(2, 4)
	defer q.stop()
	go q.tkData()

	for i := 0; i < 2; i++ {
		go q.get(i)
	}
	<-q.ch
	fmt.Println("main......")
}
