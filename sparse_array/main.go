package main

import "fmt"

type ValueNode struct {
	row   int
	col   int
	value any
}

func main() {
	array()
}

func array() {
	// 10 行 10 列 二维数组
	var arr [10][10]int
	arr[1][2] = 1
	arr[3][4] = 2
	for _, v := range arr {
		for _, vv := range v {
			fmt.Printf("%d \t", vv)
		}
		fmt.Println("")
	}
	// 存储为稀疏数组
	var sparseArr []ValueNode
	sparseArr = append(sparseArr, ValueNode{
		row:   10,
		col:   10,
		value: 0,
	})
	for i, v := range arr {
		for j, vv := range v {
			if vv != 0 {
				sparseArr = append(sparseArr, ValueNode{
					row:   i,
					col:   j,
					value: vv,
				})
			}
		}
	}
	fmt.Println("sparse_array:")
	fmt.Printf("%v\n", sparseArr)

	//还原矩阵

	first := sparseArr[0]
	sourceData := make([][]int, first.row)
	for i := 0; i < first.row; i++ {
		sourceData[i] = make([]int, first.col)
		for j := 0; j < first.col; j++ {

			sourceData[i][j] = first.value.(int)

		}
	}
	for i, v := range sourceData {
		for j := range v {
			for _, vvv := range sparseArr[1:] {
				if i == vvv.col && j == vvv.row {
					sourceData[i][j] = vvv.value.(int)
				}
			}
		}
	}
	fmt.Println("还原的数据。。。。。。")
	for _, v := range sourceData[1:] {
		for _, vv := range v {
			fmt.Printf("%d \t", vv)
		}
		fmt.Println("")
	}

}
