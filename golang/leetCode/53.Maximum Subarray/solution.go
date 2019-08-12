package main

import (
	"fmt"
)

/*
要求：


解题思路：
遍历nums,n对应的sum为n与n-1对应的sum中的较大值,同时更新maxSum

关键点：
当前index能组成的最大sum,是以index-1能组成的最大sum为基础的

*/

func maxSubArray(nums []int) int { // faster 97.18% less 59.26%
	sum, maxSum := -1<<31, -1<<31
	for _, n := range nums {
		sum = max(sum+n, n)
		maxSum = max(maxSum, sum)
	}
	return maxSum
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func main() {
	nums := []int{-2, 1, -3, 4, -1, 2, 1, -5, 4}
	result := maxSubArray(nums)
	fmt.Println(result)
}
