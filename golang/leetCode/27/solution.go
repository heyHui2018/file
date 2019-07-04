package main

import (
	"fmt"
)

/*
思路一：
仅允许在原数组上进行排序,空间复杂度要求O(1).
在26题基础上更改即可，需额外考虑nums为空、nums第一个值等于val和nums中值全等于val的情况
func removeElement(nums []int, val int) int { // faster 100% less 49.68%
	if len(nums) == 0 {
		return 0
	}
	left, right, size := 0, 1, len(nums)
	if nums[left] == val {
		flag := false
		for k := range nums {
			if nums[k] != val {
				nums[left], nums[k] = nums[k], nums[left]
				flag = true
				break
			}
		}
		if !flag {
			return 0
		}
	}
	for ; right < size; right++ {
		if nums[right] == val {
			continue
		}
		left++
		nums[left], nums[right] = nums[right], nums[left]
	}
	return left + 1
}
*/

/*
思路二：
同时从首尾开始扫描,首取等于val的位置,尾取不等于val的位置,然后互换,直到首尾扫描相交则结束
*/

func removeElement(nums []int, val int) int { // faster 100% less 49.68%
	// j指向最后一个不为val的位置
	// i指向第一个值为val的位置
	i, j := 0, len(nums)-1
	for {
		for i < len(nums) && nums[i] != val {
			i++
		}
		for j >= 0 && nums[j] == val {
			j--
		}
		if i >= j {
			break
		}
		nums[i], nums[j] = nums[j], nums[i]
	}
	return i
}

func main() {
	nums := []int{3, 2, 2, 3}
	result := removeElement(nums, 3)
	fmt.Println(result, nums)
}
