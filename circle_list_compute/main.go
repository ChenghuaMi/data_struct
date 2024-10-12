package main

import "fmt"

// n 个人围城一个圈，
// 从第 m( 0< m<= n)个人开始数数，每数 count 次把当前的人踢出去，
//直到剩下最后一个人

type PersonNode struct {
	Num  int         // 每个人的编号
	Next *PersonNode // 下一个人
}

func NewPersonNode(count int) *PersonNode {
	if count < 1 {
		fmt.Println("count 必须 >= 1")
		return nil
	}
	head := &PersonNode{}
	move := &PersonNode{} // 用于移动指针
	for i := 0; i < count; i++ {

		node := &PersonNode{
			Num: i,
		}
		if i == 0 {
			head = node
			move = node
			move.Next = head
		} else {
			move.Next = node
			move = move.Next
			move.Next = head

		}
	}
	return head
}

// 打印环形列表
func (l *PersonNode) list() {
	if l.Next == nil {
		fmt.Println("空列表")
		return
	}
	tmp := l
	for {
		if tmp.Next == l {
			break
		}
		fmt.Println("列表元素num=", tmp.Num)
		tmp = tmp.Next //指针下移
	}
}

// 移动位置，删除人员
// m 从第几个人开始移动
// count 移动的位数
func (l *PersonNode) move(m, count int) {
	tail := l // tail指针移动到环形队列的尾部
	for {
		if tail.Next == l {
			break
		}
		tail = tail.Next
	}
	//此时 l 指向环形队列队列头部，tail 指向环形队列尾部
	//从第m 个人开始移动
	for i := 0; i < m-1; i++ {
		l = l.Next
		tail = tail.Next
	}
	//移动count 次
	//这是一个多次循环
	for {
		for i := 0; i < count-1; i++ {
			l = l.Next
			tail = tail.Next
		} //此时移动到了最后
		//开始踢出人员
		fmt.Println("踢出的人员编号是：", l.Num)
		l = l.Next
		tail.Next = l
		//循环的退出情况是 l== tail
		if l == tail {
			break
		}

	}
	//这里还有最后一个人
	fmt.Println("最后一个人踢出的编号是：", l.Num)

}

func main() {
	head := NewPersonNode(6)
	head.move(3, 2)
}
