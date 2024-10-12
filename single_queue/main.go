package main

import (
	"errors"
	"fmt"
)

type queue struct {
	front   int    //队首
	tail    int    // 队为
	array   [5]int // 队列
	maxSize int    // 队列最大长度

}

func NewQueue(maxSize int) *queue {
	return &queue{
		front:   -1,
		tail:    -1,
		maxSize: maxSize,
	}
}
func (q *queue) add(data int) error {
	if q.tail == q.maxSize-1 {
		return errors.New("队列已满")
	}
	q.tail++
	q.array[q.tail] = data
	return nil
}
func (q *queue) get() (int, error) {
	if q.front == q.tail {
		return -1, errors.New("队列为空。")
	}
	q.front++
	return q.array[q.front], nil
}
func (q *queue) list() {
	for i := q.front + 1; i <= q.tail; i++ {
		fmt.Println(fmt.Sprintf("array[%d] = %d", i, q.array[i]))
	}
}

func main() {
	q := NewQueue(5)
	var name string
	var num int
	for {
		fmt.Println("增加数字 add")
		fmt.Println("获取数字 get")
		fmt.Println("列表数字 list")
		fmt.Scan(&name)
		switch name {
		case "add":
			fmt.Println("请输入数字：")
			fmt.Scan(&num)
			if err := q.add(num); err != nil {
				fmt.Println(err)
			}
		case "list":
			q.list()
		case "get":
			data, err := q.get()
			if err != nil {
				fmt.Println(err.Error())
			} else {
				fmt.Println("get data:", data)
			}

		}
	}
}
