package main

import "fmt"

type CircleNode struct {
	Num  int
	Next *CircleNode
}

func NewCircleNode(count int) *CircleNode {
	if count < 1 {
		fmt.Println("count not lt 1")
		return nil
	}
	first := &CircleNode{}
	next := &CircleNode{}
	for i := 0; i < count; i++ {
		node := &CircleNode{Num: i}
		if i == 0 {

			first = node
			next = node
			next.Next = first
		} else {

			next.Next = node
			next = next.Next
			next.Next = first
		}
	}
	return first
}

// start 从哪个位置开始移动
// count 移动的位数
func (l *CircleNode) move(start int, count int) {
	tail := l // 临时指针移动到单向环形队列尾部

	for {
		if tail.Next == l {
			break
		}
		tail = tail.Next
	}
	// tail 和 l 同时移动到start
	for i := 0; i < start-1; i++ {
		l = l.Next
		tail = tail.Next
	}
	for {
		for i := 0; i < count-1; i++ {
			l = l.Next
			tail = tail.Next
		}
		fmt.Println("move out num:", l.Num)
		l = l.Next
		tail.Next = l

		if l == tail {
			break
		}

	}
	fmt.Println("last is num:", l.Num)
}
func (l *CircleNode) list() {
	if l.Next == nil {
		fmt.Println("empty list")
		return
	}
	tmp := l
	for {
		fmt.Printf("node num: %d \n", tmp.Num)
		if tmp.Next == l {
			break
		}
		tmp = tmp.Next
	}
}
func main() {
	l := NewCircleNode(5)
	l.list()
	fmt.Println(">>>>>>>")
	l.move(2, 3)
}
