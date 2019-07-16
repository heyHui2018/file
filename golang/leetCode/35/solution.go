package main

import (
	"fmt"
)

/*
要求:


尝试：
二分查找,在迭代中考虑mid的前一位/后一位与target的关系

学习：


关键点：

*/

func searchInsert(nums []int, target int) int { // faster 93.23% less 98.89%
	lenN := len(nums)
	if lenN == 0 {
		return 0
	}
	left, right := 0, lenN-1
	mid := (left + right) / 2
	switch {
	case nums[left] > target:
		return 0
	case nums[right] < target:
		return lenN
	case nums[mid] == target:
		return mid
	}
	for left < right {
		switch {
		case nums[mid] < target:
			if mid == lenN-1 {
				return lenN
			}
			if nums[mid+1] > target {
				return mid + 1
			}
			left = mid + 1
		case nums[mid] > target:
			if mid == 0 {
				return 0
			}
			if nums[mid-1] < target {
				return mid
			}
			right = mid - 1
		default:
			return mid
		}
		mid = (left + right) / 2
	}
	return mid
}

func main() {
	nums := []int{1, 3, 5}
	result := searchInsert(nums, 1)
	fmt.Println(result)
}
