package main

import (
	"fmt"
)

/*
要求：


解题思路：
遍历找0的位置,随后往前遍历看是否有能跳跃过去的index,若有则继续往后遍历找0,若没有则返回false

关键点：


*/

func canJump(nums []int) bool { // faster 95.40% less 85.71%
	for k := range nums {
		if nums[k] == 0 && k < len(nums)-1 {
			flag := false
			for i := k - 1; i >= 0; i-- {
				fmt.Println("i = ", i, " nums[i] = ", nums[i], " k = ", k)
				if i+nums[i] > k {
					flag = true
					break
				}
			}
			if !flag {
				return false
			}
		}
	}
	return true
}

func main() {
	nums := []int{2, 0, 0}
	result := canJump(nums)
	fmt.Println(result)
}
