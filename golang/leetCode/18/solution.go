package main

import (
	"fmt"
	"sort"
)

/*
思路一：

*/

func fourSum(nums []int, target int) [][]int { // faster 100% less 5.06%
	sort.Ints(nums)
	res := [][]int{}

	for i := 0; i < len(nums)-1; i++ {
		for j := 1; j < len(nums); j++ {
			if j > 1 && nums[j] == nums[j-1] {
				continue
			}
			l, r := j+1, len(nums)-1
			for l < r {
				s := nums[i] + nums[j] + nums[l] + nums[r]
				switch {
				case s < 0:
					l++
				case s > 0:
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
	nums := []int{1, 0, -1, 0, -2, 2}
	target := 0
	result := fourSum(nums, target)
	fmt.Println(result)
}
