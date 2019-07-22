package main

import (
	"fmt"
	"sort"
)

/*
要求:
数组中的数有重复,仅能使用一次

尝试：
(未完成)

学习：
递归,对数组排序方便处理,先取出数组第一个数,则出现2种方案:1、目标数由数组中的其余数组成,则需删除与第一个数相等的数;2、余数由数组中其余数组成;
直至余数为0(采纳)或余数小于数组中剩余最小数(舍弃)或数组中没有数(舍弃)

关键点：

*/

func combinationSum2(candidates []int, target int) [][]int { // faster 100% less 35.14%
	sort.Ints(candidates)
	result := [][]int{}
	solution := []int{}
	cs(candidates, solution, target, &result)
	return result
}

func cs(candidates, solution []int, target int, result *[][]int) {
	fmt.Println("1111111 candidates = ", candidates, " target = ", target)
	if target == 0 {
		*result = append(*result, solution)
	}
	if len(candidates) == 0 || target < candidates[0] {
		return
	}
	// 这样处理一下的用意是，让切片的容量等于长度，以后append的时候，会分配新的底层数组
	// 避免多处同时对底层数组进行修改，产生错误的答案。
	// 可以注释掉以下语句，运行单元测试，查看错误发生。
	solution = solution[:len(solution):len(solution)]

	// 去掉已使用了的candidates[0]
	cs(candidates[1:], append(solution, candidates[0]), target-candidates[0], result)

	// 不使用candidates[0]的话，就要把所有和candidates[0]相等的元素都去掉。
	cs(next(candidates), solution, target, result)
}

func next(candidates []int) []int {
	i := 0
	for i+1 < len(candidates) && candidates[i] == candidates[i+1] {
		i++
	}
	return candidates[i+1:]
}

func main() {
	candidates := []int{10, 1, 2, 7, 6, 1, 5}
	result := combinationSum2(candidates, 8)
	fmt.Println(result)
}
