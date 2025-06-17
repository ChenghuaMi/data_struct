package main

import "fmt"

// golang 泛型 学习

// 1，通用的函数

func PrintSum[T int | float64 | string](a T, b T) T {
	return a + b
}

//定义map

type MyMap[a string | int, b string | int] map[a]b

type User[T string, T2 int] struct {
	Name T
	Age  T2
}

func (u *User[T, T2]) GetName() T {
	return u.Name
}
func (u *User[T, T2]) SetAge(age T2) {
	u.Age = age
}
func (u *User[T, T2]) GetAge() T2 {
	return u.Age
}

func main() {
	fmt.Println(PrintSum(1, 2))
	fmt.Println(PrintSum(1.1, 2.1))
	fmt.Println(PrintSum("a", "b"))
	var mymp MyMap[string, int] = map[string]int{
		"age": 1,
	}

	fmt.Println(mymp)
	var mymp2 MyMap[int, string] = map[int]string{
		1: "haha",
	}
	fmt.Println(mymp2)

	u := &User[string, int]{
		Name: "zs",
	}
	fmt.Println(u.GetName())
	u.SetAge(10)
	fmt.Println(u.GetAge())
}
