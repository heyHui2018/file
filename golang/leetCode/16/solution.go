package main

import (
	"fmt"
	"sort"
	"math"
)

/*
思路一：
遍历
func threeSumClosest(nums []int, target int) int { // faster 8.44% less 14.91%
	if len(nums) < 3 {
		return 0
	}
	var gap float64 = 100000
	result := 0
	for i := 0; i < len(nums)-2; i++ {
		for j := i + 1; j < len(nums)-1; j++ {
			for k := j + 1; k < len(nums); k++ {
				newGap := math.Abs(float64(nums[i] + nums[j] + nums[k] - target))
				if newGap < gap {
					result = nums[i] + nums[j] + nums[k]
					gap = newGap
				}
			}
		}
	}
	return result
}
*/

func threeSumClosest(nums []int, target int) int { // faster 99.57% less 25.89%
	sort.Ints(nums)
	res, delta := 0, math.MaxInt64
	for k := range nums {
		if k > 0 && nums[k-1] == nums[k] {
			continue
		}
		l, r := k+1, len(nums)-1
		for l < r {
			s := nums[k] + nums[l] + nums[r]
			switch {
			case s < target:
				l++
				if delta > target-s {
					delta = target - s
					res = s
				}
			case s > target:
				r--
				if delta > s-target {
					delta = s - target
					res = s
				}
			default:
				return s
			}
		}
	}
	return res
}

func main() {
	nums := []int{-1, 2, 2, -3}
	target := 1
	result := threeSumClosest(nums, target)
	fmt.Println(result)
}
