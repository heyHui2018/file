package main

import (
	"fmt"
)

/*
要求:


尝试：
(未完成)

学习：
i点的存水量为min(max(height[:i+1]...), max(height[i:]...)) - height[i],取left数组存储height[:i+1]中的最大值,取right数组存储height[i:]中的最大值,然后取i同时遍历两数组,取left,right中对应较小的

关键点：
i点的存水量为min(max(height[:i+1]...), max(height[i:]...)) - height[i]

*/

func bigger(a int, b int) int {
	if a > b {
		return a
	}
	return b
}

func smaller(a int, b int) int {
	if a < b {
		return a
	}
	return b
}

func trap(height []int) int { // faster 87.69% less 17.05%
	length := len(height)
	if length <= 2 {
		return 0
	}

	left, right := make([]int, length), make([]int, length)

	left[0], right[length-1] = height[0], height[length-1]

	for i := 1; i < length; i++ {
		left[i] = bigger(left[i-1], height[i])
		right[length-1-i] = bigger(right[length-i], height[length-1-i])

		// left[i] 是 height[:i+1] 中的最大值
		// right[i] 是 height[i:] 中的最大值
	}

	water := 0
	for i := 0; i < length; i++ {
		// 存水量取决于 左右最大值 中的较小值
		water += smaller(left[i], right[i]) - height[i]
	}

	return water
}

func main() {
	height := []int{0, 1, 0, 2, 1, 0, 1, 3, 2, 1, 2, 1}
	result := trap(height)
	fmt.Println(result)
}
