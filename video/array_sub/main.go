package main

import "fmt"

// 实现 查找出 数组中 重复 的 子数组
// 从[1,2,3,4,5,6,7,1,2,4,5,8] 查找出重复 子数组
//结果应该是 [1],[1,2],[2],[4],[4,5],[5]

func main() {
	nums := []int{1, 2, 3, 4, 5, 6, 7, 1, 2, 4, 5, 8}
	mp := make(map[string]bool) //判断子数组是否存在
	var res [][]int             // 结果 子数组
	for i := 0; i < len(nums); i++ {
		for j := i + 1; j <= len(nums); j++ {
			sum := nums[i:j] // 截取子数组
			key := fmt.Sprint(sum)
			if _, ok := mp[key]; !ok {
				mp[key] = true
			} else {
				res = append(res, sum) // 存储子数组
			}
		}
	}
	fmt.Println(res)
}
