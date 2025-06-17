package main

import "fmt"

// 实现一个简单的队列
// 入队 和 出对
type queue struct {
	ele []int
}

// 入队
func (q *queue) enqueue(x int) {
	q.ele = append(q.ele, x)
}

// 出队
func (q *queue) dequeue() int {
	if len(q.ele) == 0 {
		return -1
	}
	x := q.ele[0]
	q.ele = q.ele[1:]
	return x
}
func main() {
	q := &queue{}
	for i := 0; i < 10; i++ {
		q.enqueue(i)
	}
	//fmt.Println(q.ele)
	//eleLen := len(q.ele)
	//for i := 0; i < eleLen; i++ {
	//	fmt.Println(q.dequeue())
	//	eleLen = len(q.ele)
	//}

	for len(q.ele) > 0 { // 通过循环判断
		fmt.Println(q.dequeue()) //每次 获取 元素 ,q.ele 长度已经变了，必须每次循环重新读取
	}
}
