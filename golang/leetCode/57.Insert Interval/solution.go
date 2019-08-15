package main

import (
	"fmt"
)

/*
要求：


解题思路：
遍历intervals与newInterval比较,当intervals[i]较小且无交集时加入res,当newInterval较小且无交集且未被加入时加入res,当有交集时,更新newInterval并continue

关键点：


*/

func insert(intervals [][]int, newInterval []int) [][]int { // faster 91.02% less 100%
	res := [][]int{}
	if len(intervals) == 0 {
		return append(res, newInterval)
	}
	if len(newInterval) == 0 {
		return intervals
	}
	flag := false
	for i := 0; i < len(intervals); i++ {
		fmt.Println("000000 newInterval = ", newInterval, " intervals = ", intervals[i], " res = ", res)
		if intervals[i][1] < newInterval[0] {
			res = append(res, intervals[i])
			if i == len(intervals)-1 {
				res = append(res, newInterval)
			}
			continue
		}
		if intervals[i][0] > newInterval[1] {
			if !flag {
				res = append(res, newInterval)
			}
			res = append(res, intervals[i])
			flag = true
			continue
		}
		newInterval[0] = min(newInterval[0], intervals[i][0])
		newInterval[1] = max(newInterval[1], intervals[i][1])
		if i == len(intervals)-1 {
			res = append(res, newInterval)
		}
	}
	return res
}

func min(a, b int) int {
	if a > b {
		return b
	}
	return a
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func main() {
	intervals := [][]int{{2, 5}, {6, 7}, {8, 9}}
	newInterval := []int{0, 1}
	result := insert(intervals, newInterval)
	fmt.Println(result)
}
