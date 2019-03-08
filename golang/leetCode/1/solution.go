package main

import (
	"fmt"
)

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
