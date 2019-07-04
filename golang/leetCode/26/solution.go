package main

import (
	"fmt"
)

/*
思路一：
仅允许在原数组上进行排序,空间复杂度要求O(1),故先遍历原数组,用index记录满足条件的数字个数,循环中判断此次数字是否之前已经出现过,若没有则将此数字与第index个交换,最终返回index.
func removeDuplicates(nums []int) int { // faster 5.33% less 26.08%
	index := 0
	for k := range nums {
		fmt.Println("1111111111111 ", nums[:index], nums[k], index, nums[index])
		if !contain(nums[:index], nums[k]) {
			nums[index], nums[k] = nums[k], nums[index]
			index++
		}
	}
	return index
}

func contain(list []int, num int) bool {
	for _, v := range list {
		if v == num {
			return true
		}
	}
	return false
}
*/

/*
思路二：
用left、right两个变量来代替思路一中的contains函数以避免时间浪费.
*/

func removeDuplicates(nums []int) int { // faster 98.87% less 75.41%
	left, right, size := 0, 1, len(nums)
	for ; right < size; right++ {
		if nums[left] == nums[right] {
			continue
		}
		left++
		nums[left], nums[right] = nums[right], nums[left]
	}
	return left + 1
}

func main() {
	nums := []int{0, 0, 0, 1, 1, 1, 2, 2, 3, 3, 4}
	result := removeDuplicates(nums)
	fmt.Println(result)
}
