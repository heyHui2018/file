package main

import (
	"fmt"
)

/*
思路一：
新建个map用于存放nums的数据,遍历nums,查询map是否存在target-nums[i]的值,存在则返回,否则继续遍历.
*/
func twoSum(nums []int, target int) []int {
	tried := make(map[int]int, len(nums))
	for k, v := range nums {
		if _, ok := tried[target-v]; ok {
			return []int{tried[target-v], k}
		}
		tried[v] = k
	}
	return nil
}

func main() {
	nums := []int{2, 7, 11, 15}
	result := twoSum(nums, 17)
	fmt.Println(result)
}
