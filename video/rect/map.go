package main

import "fmt"

func charCount(s string) map[rune]int {
	rs := make(map[rune]int)
	for _, c := range s {
		rs[c]++
	}
	return rs
}
func main() {
	rs := charCount("abcda,csd,efefef")
	for key, val := range rs {
		fmt.Println(string(key), val)
	}
	r := 'a'
	rr := []rune{'a', 'A'}
	fmt.Println(r, rr, string(r), string(rr))
}
