package main

import (
	"fmt"
)

/*
思路一：
(未完成)
*/

/*
思路二：
从后往前找到最长降序序列,将其转换为升序序列,把序列前的元素与序列中第一个大于他的元素互换.当发现是纯降序序列即已是最大值时,将其转为最小值
*/

func nextPermutation(nums []int) { // faster 100.00% less 20.18%
	// 查找最长降序序列
	index := len(nums)
	if index < 2 {
		return
	}
	for i := len(nums) - 1; i > 0; i-- {
		if nums[i] > nums[i-1] {
			index = i
			break
		}
	}
	// 纯降序序列,即已是最大值
	if index == len(nums) {
		index = 0
	}
	// 将其转换为升序
	left := index
	right := len(nums) - 1
	for {
		if left >= right {
			break
		}
		nums[left], nums[right] = nums[right], nums[left]
		left++
		right--
	}
	// 交换序列前元素与序列中第一个大于它的元素
	if index > 0 {
		for i := index; i < len(nums); i++ {
			if nums[i] > nums[index-1] {
				nums[i], nums[index-1] = nums[index-1], nums[i]
				break
			}
		}
	}
}

func main() {
	nums := []int{3, 2, 1}
	nextPermutation(nums)
	fmt.Println(nums)
}
