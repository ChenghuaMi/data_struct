package main

import "fmt"

//统计字符串中相同字符出现的次数

func strCount(s string) map[rune]int {
	rs := make(map[rune]int) // rune 转成 unicode
	for _, c := range s {
		rs[c]++
	}
	return rs
}
func main() {
	str := "这是一个测试，测试，just is a test"
	res := strCount(str)
	for key, val := range res {
		fmt.Printf("字符 %s 出现次数 %d \n", string(key), val)
	}
	fmt.Println(res)
}
