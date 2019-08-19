package main

import (
	"fmt"
)

/*
要求:


尝试：
递归,当前组合的结果基于其子集的结果,如当前组合为{1,2},则有两种组合1-2和2-1,则{1,2,3}有3-1-2,1-3-2,1-2-3,3-2-1,2-3-1,2-1-3 6中组合,在{1,2}的2种组合的基础上产生的

学习：


关键点：


*/

func permute(nums []int) [][]int { // faster 93.28% less 54.76%
	result := make([][]int, 0)
	if len(nums) < 1 {
		return result
	}
	result = append(result, nums[:1])
	for i := 1; i < len(nums); i++ {
		result = p(result, nums[i])
	}
	return result
}

func p(base [][]int, num int) [][]int {
	result := make([][]int, 0)
	for i := range base {
		for j := range base[i] {
			// 此处不可使用temp:=base[i][:j],因在temp = append(temp, num)后,会使base[i]也改变
			var temp []int
			temp = append(temp, base[i][:j]...)
			temp = append(temp, num)
			temp = append(temp, base[i][j:]...)
			result = append(result, temp)
		}
		// 加上num在base[i]末尾的情况
		var temp []int
		temp = append(base[i], num)
		result = append(result, temp)
	}
	return result
}

func main() {
	nums := []int{1, 2, 3}
	result := permute(nums)
	for _, v := range result {
		fmt.Println(v)
	}
}
