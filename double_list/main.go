package main

import "fmt"

type DoubleListNode struct {
	prev  *DoubleListNode
	next  *DoubleListNode
	value int
	num   int
}

func (n *DoubleListNode) add(node *DoubleListNode) {
	tmp := n
	for {
		if tmp.next == nil {
			break
		}
		tmp = tmp.next
	}
	tmp.next = node
	node.prev = tmp
}
func (n *DoubleListNode) list() {
	tmp := n
	for {
		if tmp.next == nil {
			break
		}
		fmt.Println(fmt.Sprintf("value:%d,num:%d\n", tmp.value, tmp.num))
		tmp = tmp.next
	}
}
func (n *DoubleListNode) delete(node *DoubleListNode) {
	tmp := n
	for {
		if tmp.next == nil {
			break
		}

		if tmp.next.num == node.num {

			tmp.next = tmp.next.next
			if tmp.next != nil {
				tmp.next.prev = tmp
			}

			break
		}
		tmp = tmp.next
	}
}
func main() {
	doubleList := &DoubleListNode{}

	node1 := &DoubleListNode{num: 1, value: 1}
	node2 := &DoubleListNode{num: 2, value: 2}
	node3 := &DoubleListNode{num: 3, value: 3}
	doubleList.add(node1)
	doubleList.add(node2)
	doubleList.add(node3)
	doubleList.delete(node2)
	doubleList.list()
}
