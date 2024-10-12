package main

import "fmt"

type SingleNode struct {
	name string
	no   int // 按照no 的顺序插入
	next *SingleNode
}

func (n *SingleNode) AddNode(name string) {
	node := &SingleNode{
		name: name,
	}
	//找到最后一个节点
	tmp := n
	for {
		if tmp.next == nil {
			break
		}
		tmp = tmp.next
	}
	tmp.next = node
}
func (n *SingleNode) AddNodeByNo(node *SingleNode) {
	tmp := n
	for {
		if tmp.next == nil { // 空列表
			tmp.next = node
			break
		} else if tmp.next.no >= node.no {
			//新节点插入到tmp后面
			node.next = tmp.next
			tmp.next = node
			break
		}
		tmp = tmp.next
	}

}
func (n *SingleNode) ListNode() {
	if n == nil {
		return
	}
	tmp := n
	if tmp == nil {
		return
	}
	for {

		fmt.Println("node:", tmp.next.name)
		tmp = tmp.next
		if tmp.next == nil {
			break
		}
	}
}
func (n *SingleNode) DeleteNode(node *SingleNode) {
	tmp := n
	for {
		if tmp == nil || tmp.next == nil {
			break
		}
		if tmp.next.no == node.no {
			tmp.next = tmp.next.next
			break
		}
		tmp = tmp.next
	}
}
func main() {
	node := &SingleNode{}
	node.AddNodeByNo(&SingleNode{
		name: "zs",
		no:   5,
	})
	node.AddNodeByNo(&SingleNode{
		name: "lisi",
		no:   2,
	})
	node.DeleteNode(&SingleNode{no: 5})
	node.DeleteNode(&SingleNode{no: 2})
	node.ListNode()
}
