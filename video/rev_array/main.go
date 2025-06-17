package main

import "fmt"

// 实现字符串倒转
// this is golang
// 变为 gnalog si siht

func main() {
	str := "this is golang"
	b := []byte(str)
	for i, j := 0, len(b)-1; i < j; i, j = i+1, j-1 {
		b[i], b[j] = b[j], b[i] // go 中 字符串是不能修改的，所以需要转换成 byte类型
	}
	//byte 在转换成字符串
	fmt.Println(string(b))
}
