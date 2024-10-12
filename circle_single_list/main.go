package main

import "fmt"

// 环形单列表

type CircleSingleNode struct {
	no   int
	val  int
	next *CircleSingleNode
}

func (q *CircleSingleNode) add(node *CircleSingleNode) {
	if q.next == nil { // 如果是空列表，将加入的节点赋予给列表
		q.no = node.no
		q.val = node.val
		q.next = q
		return
	}
	tmp := q // 记录q 的起始位置
	for {
		if tmp.next == q {
			// 找到环形队列的最后位置
			tmp.next = node
			node.next = q
			break
		}
		tmp = tmp.next
	}
}
func (q *CircleSingleNode) delete(no int) {
	tmp := q
	if tmp.next == nil {
		fmt.Println("empty queue")
		return
	}
	for {
		if tmp.next.no == no {
			tmp.next = tmp.next.next
			break
		}
		tmp = tmp.next
	}
}
func (q *CircleSingleNode) deleteNode(no int) *CircleSingleNode {
	tmp := q        //用于节点的值比较
	tmpPointer := q // 用于删除节点指针的移动
	if q.next == nil {
		fmt.Println("empty queue")
		return q
	}
	if q.next == q {
		// 环形列表只有一个元素
		q.next = nil
		return q
	}
	// 环形队列有多个元素
	// 将tmpPointer 移动到环形队列的尾部
	for {
		if tmpPointer.next == q {
			break
		}
		tmpPointer = tmpPointer.next
	}
	for {
		if tmp.next == q { // 最后一个元素的操作

			if tmp.no == no {
				tmpPointer.next = tmp.next
			}
			break
		}
		if tmp.no == no { //头节点到===倒数第二个节点的操作
			if tmp == q { //删除的是头节点
				q = q.next // 头节点后移一位
			}
			tmpPointer.next = tmp.next
			break
		}
		tmp = tmp.next
		tmpPointer = tmpPointer.next
	}
	return q
}
func (q *CircleSingleNode) list() {
	tmp := q
	if tmp.next == nil {
		fmt.Println("empty list")
		return
	}
	for {
		fmt.Printf("%+v \n", tmp)
		if tmp.next == q {
			break
		}
		tmp = tmp.next
	}
}

func main() {
	q := &CircleSingleNode{}
	node1 := &CircleSingleNode{
		no:  1,
		val: 1,
	}
	node2 := &CircleSingleNode{
		no:  2,
		val: 2,
	}
	q.add(node1)
	q.add(node2)
	//q.delete(1)
	q = q.deleteNode(1)
	q.list()
}
