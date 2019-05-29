package main

import (
	"fmt"
	"sort"
)

/*
思路一：
比3Sum多一层循环
*/

func fourSum(nums []int, target int) [][]int { // faster 93.19% less 79.93%
	sort.Ints(nums)
	res := [][]int{}

	for i := 0; i < len(nums)-3; i++ { // 0到len(nums)-3
		if i > 0 && nums[i] == nums[i-1] {
			continue
		}
		for j := i + 1; j < len(nums)-2; j++ { // i+1到len(nums)-2
			if j > i+1 && nums[j] == nums[j-1] {
				continue
			}
			l, r := j+1, len(nums)-1
			for l < r {
				s := nums[i] + nums[j] + nums[l] + nums[r]
				switch {
				case s < target:
					l++
				case s > target:
					r--
				default:
					res = append(res, []int{nums[i], nums[j], nums[l], nums[r]})
					l, r = next(nums, l, r)
				}
			}
		}

	}
	return res
}

func next(nums []int, l, r int) (int, int) {
	for l < r {
		switch {
		case nums[l] == nums[l+1]:
			l++
		case nums[r] == nums[r-1]:
			r--
		default:
			l++
			r--
			return l, r
		}
	}
	return l, r
}

func main() {
	nums := []int{-3, -2, -1, 0, 0, 1, 2, 3}
	target := 0
	result := fourSum(nums, target)
	fmt.Println(result)
}
