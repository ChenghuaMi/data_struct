package main

import (
	"errors"
	"fmt"
	"os"
)

type CircleQueue struct {
	maxSize int
	array   [5]int
	head    int
	tail    int //指向队尾巴
}

func (q *CircleQueue) push(value int) (err error) {
	if q.isFull() {
		return errors.New("queue full")
	}
	q.array[q.tail] = value
	q.tail = (q.tail + 1) % q.maxSize
	return
}
func (q *CircleQueue) pop() (val int, err error) {
	if q.isEmpty() {
		return -1, errors.New("queue empty")
	}
	val = q.array[q.head]
	q.head = (q.head + 1) % q.maxSize
	return
}
func (q *CircleQueue) list() {
	ss := q.size()
	if ss == 0 {
		fmt.Println("queue empty")
		return
	}
	tmp := q.head
	for i := 0; i < ss; i++ {
		fmt.Printf("array[%d]=%d \t", tmp, q.array[tmp])
		tmp = (tmp + 1) % q.maxSize
	}
	fmt.Println()
}
func (q *CircleQueue) isFull() bool {
	return (q.tail+1)%q.maxSize == q.head
}
func (q *CircleQueue) isEmpty() bool {
	return q.head == q.tail
}
func (q *CircleQueue) size() int {
	return (q.tail + q.maxSize - q.head) % q.maxSize
}
func main() {
	q := &CircleQueue{
		maxSize: 5,
	}
	var name string
	var data int
	for {
		fmt.Println("输入方法名称：push | pop | list | exit")
		fmt.Scanln(&name)
		switch name {
		case "push":
			fmt.Println("请输入值:")
			fmt.Scanln(&data)
			if err := q.push(data); err != nil {
				fmt.Println("err:", err)
			}
		case "pop":
			val, err := q.pop()
			if err != nil {
				fmt.Println("err:", err)
			} else {
				fmt.Println("get val:", val)
			}
		case "list":
			q.list()
		case "exit":
			os.Exit(0)

		}
	}
}
