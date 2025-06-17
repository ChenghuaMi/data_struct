package main

import "fmt"

type rect struct {
	size int
	data [][]int
}

func newRect(size int) *rect {
	data := make([][]int, size)
	for i := range data {
		data[i] = make([]int, size)
	}
	return &rect{
		size: size,
		data: data,
	}
}
func (r *rect) addData(row, col int, val int) {
	if row >= 0 && row < r.size && col >= 0 && col < r.size {
		r.data[row][col] = val
		r.data[col][row] = val
	}
}
func (r *rect) delData(row, col int) {
	if row >= 0 && row < r.size && col >= 0 && col < r.size {
		r.data[row][col] = 0
		r.data[col][row] = 0
	}
}
func (r *rect) printRect() {
	for i := 0; i < r.size; i++ {
		for j := 0; j < r.size; j++ {
			fmt.Printf("%d  ", r.data[i][j])
		}
		fmt.Println()
	}
}
func main() {
	r := newRect(10)
	r.addData(0, 4, 1)
	r.addData(3, 2, 1)
	r.printRect()
}
