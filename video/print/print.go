package main

import "fmt"

// 打印 一个 乘法表
// 有点忘了
// 翻车了，吃饭去了
//没准备。临时发挥

func main() {
	for i := 1; i <= 9; i++ {
		for j := 1; j <= i; j++ {
			fmt.Printf("%d X %d = %d\t", j, i, j*i)
		}
		fmt.Println()
	}
}
