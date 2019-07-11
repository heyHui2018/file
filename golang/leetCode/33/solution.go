package main

import (
	"fmt"
)

/*
要求:
时间复杂度为O(logn)

尝试：
目标与数组第一个数进行比较,目标较大则开始遍历查找,目标较小则从尾部开始遍历查找,此法时间复杂度O(n).
func search(nums []int, target int) int { // faster 100% less 93.10%
	if len(nums) > 0 {
		switch {
		case nums[0] > target:
			// 因升序排列,当头比target大时,可直接从尾获取其index
			for i := 1; i < len(nums); i++ {
				if nums[len(nums)-i] == target {
					return len(nums) - i
				}
			}
		case nums[0] < target:
			for i := 0; i < len(nums); i++ {
				if nums[i] == target {
					return i
				}
			}
		default:
			return 0
		}
	}
	return -1
}

学习：
先计算旋转距离,随后还原数组进行二分查找时间复杂度O(2logn).

关键点：
时间复杂度,只能用二分查找
*/

func search(nums []int, target int) int { // faster 100% less 93.10%
	rotated := indexOfMin(nums) /* 数组旋转了的距离 */
	fmt.Println(rotated)
	size := len(nums)
	left, right := 0, size-1

	for left <= right {
		mid := (left + right) / 2
		/* nums 是 rotated，所以需要使用 rotatedMid 来获取 mid 的值 */
		rotatedMid := (rotated + mid) % size
		switch {
		case nums[rotatedMid] < target:
			left = mid + 1
		case target < nums[rotatedMid]:
			right = mid - 1
		default:
			return rotatedMid
		}
	}

	return -1
}

/* nums 是被旋转了的递增数组 */
func indexOfMin(nums []int) int {
	size := len(nums)
	left, right := 0, size-1
	for left < right {
		mid := (left + right) / 2
		if nums[right] < nums[mid] {
			left = mid + 1
		} else {
			right = mid
		}
	}
	return left
}

func main() {
	nums := []int{4, 5, 6, 7, 0, 1, 2}
	result := search(nums, 3)
	fmt.Println(result)
}
