package main

import (
	"fmt"
)

/*
要求:
时间复杂度O(logn)

尝试：
二分查找,当中间值=目标值时向前后遍历以找到目标值的最前/最后index
func searchRange(nums []int, target int) []int { // faster 89.15% less 82.50%
	result := []int{-1, -1}
	lenN := len(nums)
	if lenN > 0 {
		left, right := 0, lenN-1
		mid := lenN / 2
		fmt.Println("mid = ", mid, " nums[mid] = ", nums[mid], " result = ", result)
		for left <= right {
			switch {
			case nums[mid] > target:
				right = mid - 1
			case nums[mid] < target:
				left = mid + 1
			default:
				left, right = mid, mid
				for {
					if left == 0 || (left > 0 && nums[left] == target && nums[left-1] < target) {
						result[0] = left
					} else {
						left--
					}
					if right == lenN-1 || (right < lenN-1 && nums[right] == target && nums[right+1] > target) {
						result[1] = right
					} else {
						right++
					}
					if result[0] != -1 && result[1] != -1 {
						return result
					}
				}
			}
			mid = (left + right) / 2
			// fmt.Println("mid = ", mid, " nums[mid] = ", nums[mid], " left = ", left, " nums[left] = ", nums[left], " right = ", right, " nums[right] = ", nums[right])
		}
	}
	return result
}

学习：


关键点：

*/

func searchRange(nums []int, target int) []int { // faster 89.15% less 82.50%
	result := []int{-1, -1}
	lenN := len(nums)
	if lenN > 0 {
		left, right := 0, lenN-1
		mid := lenN / 2
		for left <= right {
			switch {
			case nums[mid] > target:
				right = mid - 1
			case nums[mid] < target:
				left = mid + 1
			default:
				left, right = mid, mid
				for {
					if left == 0 || (left > 0 && nums[left] == target && nums[left-1] < target) {
						result[0] = left
					} else {
						left--
					}
					if right == lenN-1 || (right < lenN-1 && nums[right] == target && nums[right+1] > target) {
						result[1] = right
					} else {
						right++
					}
					if result[0] != -1 && result[1] != -1 {
						return result
					}
				}
			}
			mid = (left + right) / 2
		}
	}
	return result
}

func main() {
	nums := []int{1}
	result := searchRange(nums, 1)
	fmt.Println(result)
}
