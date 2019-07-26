package main

import (
	"fmt"
)

/*
要求:


尝试：
func jump(nums []int) int { // faster 40.46% less 73.68%
	count := 0
	for i := 0; i < len(nums); {
		if i == len(nums)-1 {
			return count
		} else if i+nums[i] >= len(nums)-1 {
			return count + 1
		}
		i += FindTheMostfarther(nums[i+1:i+nums[i]+1]) + 1
		count++
		fmt.Println("i = ", i, " count = ", count)
	}
	return count
}

func FindTheMostfarther(nums []int) int {
	index := 0
	for i := 0; i < len(nums); i++ {
		if i+nums[i] >= index+nums[index] {
			index = i
		}
	}
	fmt.Println("nums = ", nums, " index = ", index)
	return index
}


学习：


关键点：
确保每次跳都是最远的即index+nums[index]在nums[i+1:i+nums[i]+1]中是最大的

*/

func jump(nums []int) int { // faster 100% less 73.68%
	i, count, end := 0, 0, len(nums)-1
	var nextI, maxNextI, maxI int
	for i < end {
		if i+nums[i] >= end {
			return count + 1
		}
		nextI, maxNextI = i+1, i+nums[i]
		for nextI <= maxNextI {
			if nextI+nums[nextI] > maxI {
				maxI, i = nextI+nums[nextI], nextI
			}
			nextI++
		}
		count++
	}
	return count
}

func main() {
	nums := []int{10, 9, 8, 7, 6, 5, 4, 3, 2, 1, 1, 0}
	result := jump(nums)
	fmt.Println(result)
}
