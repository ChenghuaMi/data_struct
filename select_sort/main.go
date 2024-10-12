package main

import "fmt"

func main() {
	var arr = []int{9, 4, 5, 2, 7, 10}
	fmt.Println("sort before:")
	fmt.Println(arr)
	SelectSort(arr)
}
func SelectSort(arr []int) {
	for i := 0; i < len(arr)-1; i++ {
		minIndex := i
		minNum := arr[i]
		for j := i + 1; j < len(arr); j++ {
			if minNum > arr[j] {
				minIndex = j
				minNum = arr[j]
			}
		}
		if minIndex != i {
			arr[i], arr[minIndex] = arr[minIndex], arr[i]
		}

	}
	fmt.Println("sort after...")
	fmt.Println(arr)
}
