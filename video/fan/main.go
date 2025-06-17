package main

import "fmt"

func printArr[T any](arr []T) {
	for _, v := range arr {
		fmt.Println(v)
	}
}
func main() {
	arr := []int{1, 23, 4}
	printArr(arr)
	arr2 := []string{"aq", "b"}
	printArr(arr2)

	type myMap[key string | int, value int | string] map[key]value
	var mp myMap[string, int] = map[string]int{
		"name": 1,
		"age":  10,
	}
	fmt.Println(mp)

	type MyStru[T string | int, T2 int] struct {
		Name T
		Age  T2
	}
	fmt.Println(MyStru[string, int]{Name: "zs", Age: 1})

	mm := MySli[int]{1, 2, 3}
	mm.Sum()
	fmt.Println(compareMax[int](1, 3))
}

type MySli[T int | float64] []T

func (m MySli[T]) Sum() T {
	var result T
	for _, v := range m {
		result += v
	}
	println(result)
	return result
}

type MyComp interface {
	int | int32 | float64
}

func compareMax[T MyComp](a T, b T) T {
	if a > b {
		return a
	}
	return b
}
