package main

import "fmt"

// 快速排序
func quickSort(arr []int) []int {
	if len(arr) <= 1 {
		return arr
	}
	first := arr[0]
	less, greater := make([]int, 0), make([]int, 0)
	for _, v := range arr[1:] {
		if v <= first {
			less = append(less, v)
		} else {
			greater = append(greater, v)
		}
	}
	less = quickSort(less)
	greater = quickSort(greater)
	return append(append(less, first), greater...)
}
func main() {
	arr := []int{5, 2, 8, 3, 1, 6, 4}
	rs := quickSort(arr)
	fmt.Println(rs)
}
