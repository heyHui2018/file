package main

import (
	"fmt"
	"sort"
)

/*
要求：


解题思路：
先根据数组的第一个值递增排序,随后逐一比较相邻的数组,有重叠则合并,合并后再与下一个数组比较

关键点：


*/

func merge(intervals [][]int) [][]int { // faster 96.13% less 100%
	if len(intervals) < 2 {
		return intervals
	}
	sort.Slice(intervals, func(i, j int) bool {
		return intervals[i][0] < intervals[j][0]
	})
	res := [][]int{}
	res = append(res, intervals[0])
	for i := 1; i < len(intervals); i++ {
		if res[len(res)-1][1] >= intervals[i][0] {
			res[len(res)-1][0] = min(res[len(res)-1][0], intervals[i][0])
			res[len(res)-1][1] = max(res[len(res)-1][1], intervals[i][1])
			continue
		}
		res = append(res, intervals[i])
	}
	return res
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func main() {
	nums := [][]int{{1, 4}, {0, 2}, {3, 5}}
	result := merge(nums)
	fmt.Println(result)
}
