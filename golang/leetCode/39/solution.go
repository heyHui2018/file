package main

import (
	"fmt"
	"sort"
)

/*
要求:
数组中的数无重复,可使用多次

尝试：
(未完成)

学习：
递归,对数组排序方便处理,先取出数组第一个数,则出现2种方案:1、余数由数组中的数组成;2、目标数由数组中其余数组成;宗旨就是要么减少余数,要么减少数组中的数,
直至余数为0(采纳)或余数小于数组中剩余最小数(舍弃)或数组中没有数(舍弃)

关键点：
solution = solution[:len(solution):len(solution)]
取solution的0-len(solution)并将cap设为len(solution)
*/

func combinationSum(candidates []int, target int) [][]int { // faster 8.67% less 42.86%
	sort.Ints(candidates)
	result := [][]int{}
	solution := []int{}
	cs(candidates, solution, target, &result)
	return result
}

func cs(candidates, solution []int, target int, result *[][]int) {
	if target == 0 {
		*result = append(*result, solution)
	}

	// target < candidates[0] 因为candidates是排序好的
	if len(candidates) == 0 || target < candidates[0] {
		return
	}

	// 这样处理一下的用意是，让切片的容量等于长度，以后append的时候，会分配新的底层数组
	// 避免多处同时对底层数组进行修改，产生错误的答案。
	// todo
	solution = solution[:len(solution):len(solution)]

	cs(candidates, append(solution, candidates[0]), target-candidates[0], result)

	cs(candidates[1:], solution, target, result)
}

func main() {
	candidates := []int{2, 3, 7}
	result := combinationSum(candidates, 18)
	fmt.Println(result)
}
