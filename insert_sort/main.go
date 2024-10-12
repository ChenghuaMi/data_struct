package main

import "fmt"

// arr 3,2,6,9,7,10
//
//	3,2,2,9,7,10
//	3,3,2,9,7,10
func insertSort(arr []int) {
	for i := 1; i < len(arr); i++ {
		insertVal := arr[i]
		insertPos := i - 1
		for insertPos >= 0 && arr[insertPos] < insertVal {
			arr[insertPos+1] = arr[insertPos]
			insertPos--
		}
		arr[insertPos+1] = insertVal
		fmt.Printf("第 %d 次比较：%v\n", i, arr)
	}

}
func main() {
	arr := []int{2, 31, 6, 9, 7, 10}
	fmt.Println("before:", arr)
	insertSort(arr)
	//fmt.Println("after:", arr)
}
