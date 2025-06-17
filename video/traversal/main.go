package main

import "fmt"

//实现 channel 遍历树 的结构

type Node struct {
	Val       int
	LeftNode  *Node
	RightNode *Node
}

func traversal(ch chan int, node *Node) {
	if node == nil {
		return
	}
	ch <- node.Val
	//递归调用
	traversal(ch, node.LeftNode)  //递归左结点
	traversal(ch, node.RightNode) //递归右结点
}
func main() {
	node := &Node{
		Val: 1,
		LeftNode: &Node{
			Val: 2,
			LeftNode: &Node{
				Val:       3,
				LeftNode:  &Node{Val: 4},
				RightNode: &Node{Val: 5},
			},
		},
		RightNode: &Node{
			Val:       6,
			LeftNode:  &Node{Val: 7},
			RightNode: &Node{Val: 8},
		},
	}
	ch := make(chan int)
	//go traversal(ch, node) //遍历结点，
	go func() {
		traversal(ch, node) //执行完毕，关闭通道
		close(ch)
	}()
	//channel中读取值
	for v := range ch { //死锁，因为 channel 没有关闭
		fmt.Println("val:", v)
	}
	fmt.Println("main...")
}
