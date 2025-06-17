package main

import "fmt"

// 实现 字符串 的倒转
// abcdef = > fedcba

func main() {
	str := "abcdef"
	byt := []byte(str)
	for i, j := 0, len(byt)-1; i < j; i, j = i+1, j-1 {
		byt[i], byt[j] = byt[j], byt[i]
	}

	fmt.Println(string(byt))
}
