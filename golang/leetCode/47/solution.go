package main

import (
	"fmt"
)

/*
要求:


尝试：
当前组合的结果基于其子集的结果,如1/2有两种,1/2/3则为1/2的两种组合基础上加上3这个数形成的组合

学习：


关键点：


*/

func permuteUnique(nums []int) [][]int { // faster 93.28% less 54.76%
	
}

func main() {
	nums := []int{1, 2, 3}
	result := permute(nums)
	for _, v := range result {
		fmt.Println(v)
	}
}
